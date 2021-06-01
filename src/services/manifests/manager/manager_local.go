package manager

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/errors"
	manifest "github.com/ConsenSysQuorum/quorum-key-manager/src/services/manifests/types"
	"gopkg.in/yaml.v2"
)

const ManagerID = "LocalManager"

type Config struct {
	Path string
}

type LocalManager struct {
	path  string
	isDir bool

	msgs []Message

	loaded chan struct{}
	err    error
}

func NewLocalManager(cfg *Config) (*LocalManager, error) {
	fs, err := os.Stat(cfg.Path)
	if err == nil {
		return &LocalManager{
			path:   cfg.Path,
			loaded: make(chan struct{}),
			isDir:  fs.IsDir(),
		}, nil
	}

	if os.IsNotExist(err) {
		return nil, errors.InvalidParameterError("folder or file does not exists. %s", cfg.Path)
	}

	return nil, err
}

type subscription struct {
	kinds    map[manifest.Kind]struct{}
	messages chan<- []Message
	errors   chan error
	stop     chan struct{}
	done     chan struct{}
}

func (sub *subscription) Unsubscribe() error {
	close(sub.stop)
	<-sub.done
	close(sub.errors)
	return nil
}

func (sub *subscription) Error() <-chan error { return sub.errors }

func (sub *subscription) inbox(msgs []Message) {
	var submsgs []Message
	for _, msg := range msgs {
		if sub.kinds == nil {
			submsgs = append(submsgs, msg)
			continue
		}

		if _, ok := sub.kinds[msg.Manifest.Kind]; ok {
			submsgs = append(submsgs, msg)
		}
	}

	sub.messages <- submsgs
}

func (ll *LocalManager) Subscribe(kinds []manifest.Kind, messages chan<- []Message) (Subscription, error) {
	sub := &subscription{
		messages: messages,
		errors:   make(chan error, 1),
		stop:     make(chan struct{}),
		done:     make(chan struct{}),
	}

	if kinds != nil {
		sub.kinds = make(map[manifest.Kind]struct{})
		for _, kind := range kinds {
			sub.kinds[kind] = struct{}{}
		}
	}

	go ll.processSub(sub)

	return sub, nil
}

func (ll *LocalManager) processSub(sub *subscription) {
	defer close(sub.done)

	select {
	case <-ll.loaded:
		if ll.err != nil {
			sub.errors <- ll.err
		} else {
			sub.inbox(ll.msgs)
		}
	case <-sub.stop:
	}
}

func (ll *LocalManager) load() error {
	if ll.isDir {
		return filepath.Walk(ll.path, func(fp string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			if filepath.Ext(fp) == ".yml" || filepath.Ext(fp) == ".yaml" {
				ll.msgs = append(ll.msgs, ll.buildMessages(fp)...)
			}

			return nil
		})
	}

	ll.msgs = append(ll.msgs, ll.buildMessages(ll.path)...)

	return nil
}

func (ll *LocalManager) Start(context.Context) error {
	defer close(ll.loaded)
	ll.err = ll.load()
	return ll.err
}

func (ll *LocalManager) buildMessages(fp string) []Message {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		return []Message{{
			Loader: ManagerID,
			Action: CreateAction,
			Err:    err,
		}}
	}

	mnf := &manifest.Manifest{}
	if err = yaml.Unmarshal(data, mnf); err == nil {
		return []Message{{
			Loader:   ManagerID,
			Action:   CreateAction,
			Manifest: mnf,
		}}
	}

	mnfs := []*manifest.Manifest{}
	if err = yaml.Unmarshal(data, &mnfs); err == nil {
		msgs := []Message{}
		for _, mnf := range mnfs {
			msgs = append(msgs, Message{
				Loader:   ManagerID,
				Action:   CreateAction,
				Manifest: mnf,
			})
		}
		return msgs
	}

	return []Message{{
		Loader: ManagerID,
		Action: CreateAction,
		Err:    err,
	}}
}

func (ll *LocalManager) Stop(context.Context) error { return nil }
func (ll *LocalManager) Error() error               { return ll.err }
func (ll *LocalManager) Close() error               { return nil }
