package hashicorp

import (
	"context"
	"fmt"
	"net/http"

	"time"

	"github.com/docker/docker/api/types/mount"

	"github.com/consensysquorum/quorum-key-manager/tests/acceptance/docker/config/hashicorp"
	dockercontainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	log "github.com/sirupsen/logrus"
)

type Vault struct{}

func (vault *Vault) GenerateContainerConfig(_ context.Context, configuration interface{}) (*dockercontainer.Config, *dockercontainer.HostConfig, *network.NetworkingConfig, error) {
	cfg, ok := configuration.(*hashicorp.Config)
	if !ok {
		return nil, nil, nil, fmt.Errorf("invalid configuration type (expected %T but got %T)", cfg, configuration)
	}

	containerCfg := &dockercontainer.Config{
		Image: cfg.Image,
		Env: []string{
			fmt.Sprintf("VAULT_DEV_ROOT_TOKEN_ID=%v", cfg.RootToken),
		},
		ExposedPorts: nat.PortSet{
			"8200/tcp": struct{}{},
		},
		Tty: true,
		Cmd: []string{"server", "-dev", "-dev-plugin-dir=/vault/plugins", "-log-old-level=trace"},
	}

	hostConfig := &dockercontainer.HostConfig{
		CapAdd: []string{"IPC_LOCK"},
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: cfg.PluginSourceDirectory,
				Target: "/vault/plugins",
			},
		},
	}
	if cfg.Port != "" {
		hostConfig.PortBindings = nat.PortMap{
			"8200/tcp": []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: cfg.Port}},
		}
	}

	return containerCfg, hostConfig, nil, nil
}

func (vault *Vault) WaitForService(ctx context.Context, configuration interface{}, timeout time.Duration) error {
	cfg, ok := configuration.(*hashicorp.Config)
	if !ok {
		return fmt.Errorf("invalid configuration type (expected %T but got %T)", cfg, configuration)
	}

	rctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	retryT := time.NewTicker(2 * time.Second)
	defer retryT.Stop()

	httpClient := http.Client{}

	var cerr error
waitForServiceLoop:
	for {
		select {
		case <-rctx.Done():
			cerr = rctx.Err()
			break waitForServiceLoop
		case <-retryT.C:
			resp, err := httpClient.Get(fmt.Sprintf("http://%v:%v/v1/sys/health", cfg.Host, cfg.Port))

			switch {
			case err != nil:
				log.WithContext(rctx).WithError(err).Warnf("waiting for Hashicorp Vault service to start")
			case resp.StatusCode != http.StatusOK:
				log.WithContext(rctx).WithField("status_code", resp.StatusCode).Warnf("waiting for Hashicorp Vault service to be ready")
			default:
				log.WithContext(rctx).Infof("Hashicorp Vault container service is ready")
				break waitForServiceLoop
			}
		}
	}

	return cerr
}
