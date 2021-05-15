package acceptancetests

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/ConsenSysQuorum/quorum-key-manager/src/core/types"
	"github.com/ConsenSysQuorum/quorum-key-manager/tests"
	"gopkg.in/yaml.v2"

	keymanager "github.com/ConsenSysQuorum/quorum-key-manager/src"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/core/manifest"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/core/store-manager/hashicorp"
	akvclient "github.com/ConsenSysQuorum/quorum-key-manager/src/infra/akv/client"
	hashicorp2 "github.com/ConsenSysQuorum/quorum-key-manager/src/infra/hashicorp"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/infra/http"
	"github.com/ConsenSysQuorum/quorum-key-manager/tests/acceptance/utils"
	"k8s.io/apimachinery/pkg/util/rand"

	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/common"
	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/log"
	akv2 "github.com/ConsenSysQuorum/quorum-key-manager/src/infra/akv"
	hashicorpclient "github.com/ConsenSysQuorum/quorum-key-manager/src/infra/hashicorp/client"
	"github.com/ConsenSysQuorum/quorum-key-manager/tests/acceptance/docker"
	dconfig "github.com/ConsenSysQuorum/quorum-key-manager/tests/acceptance/docker/config"
	"github.com/hashicorp/vault/api"
)

const (
	hashicorpContainerID      = "hashicorp-vault"
	networkName               = "key-manager"
	localhostPath             = "http://localhost"
	HashicorpSecretStoreName  = "HashicorpSecrets"
	HashicorpKeyStoreName     = "HashicorpKeys"
	HashicorpSecretMountPoint = "secret"
	HashicorpKeyMountPoint    = "orchestrate"
	AKVSecretStoreName        = "AKVSecrets"
	AKVKeyStoreName           = "AKVKeys"
)

type IntegrationEnvironment struct {
	ctx               context.Context
	logger            *log.Logger
	hashicorpClient   hashicorp2.VaultClient
	akvClient         akv2.Client
	dockerClient      *docker.Client
	keyManager        *keymanager.App
	baseURL           string
	Cancel            context.CancelFunc
	tmpManifestYaml   string
	tmpHashicorpToken string
	cfg               *tests.Config
}

type TestSuiteEnv interface {
	Start(ctx context.Context) error
}

func StartEnvironment(ctx context.Context, env TestSuiteEnv) (gerr error) {
	ctx, cancel := context.WithCancel(ctx)

	sig := common.NewSignalListener(func(signal os.Signal) {
		gerr = fmt.Errorf("interrupt signal has been sent")
		cancel()
	})
	defer sig.Close()

	err := env.Start(ctx)
	if err != nil {
		if gerr == nil {
			return err
		}
	}

	return
}

func NewIntegrationEnvironment(ctx context.Context) (*IntegrationEnvironment, error) {
	logger := log.DefaultLogger().WithContext(ctx)

	hashicorpContainer, err := utils.HashicorpContainer(ctx)
	if err != nil {
		return nil, err
	}

	tmpTokenFile, err := newTmpFile(hashicorpContainer.RootToken)
	if err != nil {
		return nil, err
	}

	// Initialize environment container setup
	composition := &dconfig.Composition{
		Containers: map[string]*dconfig.Container{
			hashicorpContainerID: {
				HashicorpVault: hashicorpContainer,
			},
		},
	}

	// Docker client
	dockerClient, err := docker.NewClient(composition)
	if err != nil {
		logger.WithError(err).Error("cannot initialize new environment")
		return nil, err
	}

	testCfg, err := tests.NewConfig()
	if err != nil {
		logger.WithError(err).Error("cannot initialize new environment")
		return nil, err
	}

	envHTTPPort := rand.IntnRange(20000, 28080)
	hashicorpAddr := fmt.Sprintf("http://%s:%s", hashicorpContainer.Host, hashicorpContainer.Port)
	tmpYml, err := newTmpManifestYml(&manifest.Manifest{
		Kind: types.HashicorpSecrets,
		Name: HashicorpSecretStoreName,
		Specs: &hashicorp.SecretSpecs{
			Token:      hashicorpContainer.RootToken,
			MountPoint: HashicorpSecretMountPoint,
			Address:    hashicorpAddr,
			Namespace:  "",
		},
	}, &manifest.Manifest{
		Kind: types.HashicorpKeys,
		Name: HashicorpKeyStoreName,
		Specs: &hashicorp.KeySpecs{
			MountPoint: HashicorpKeyMountPoint,
			Address:    hashicorpAddr,
			TokenPath:  tmpTokenFile,
			Namespace:  "",
		},
	}, &manifest.Manifest{
		Kind:  types.AKVSecrets,
		Name:  AKVSecretStoreName,
		Specs: testCfg.AkvSecretSpecs(),
	}, &manifest.Manifest{
		Kind:  types.AKVKeys,
		Name:  AKVKeyStoreName,
		Specs: testCfg.AkvKeySpecs(),
	})

	if err != nil {
		logger.WithError(err).Error("cannot create keymanager manifest")
		return nil, err
	}

	logger.WithField("path", tmpYml).Info("new temporal manifest created")

	httpConfig := http.NewDefaultConfig()
	httpConfig.Port = uint32(envHTTPPort)
	keyManager, err := newKeyManager(&keymanager.Config{
		HTTP:         httpConfig,
		ManifestPath: tmpYml,
	}, logger)
	if err != nil {
		logger.WithError(err).Error("cannot initialize Key Manager server")
		return nil, err
	}

	// Hashicorp client for direct integration tests
	hashicorpCfg := hashicorpclient.NewConfig(hashicorpAddr, "")
	hashicorpClient, err := hashicorpclient.NewClient(hashicorpCfg)
	if err != nil {
		logger.WithError(err).Error("cannot initialize hashicorp vault client")
		return nil, err
	}
	hashicorpClient.Client().SetToken(hashicorpContainer.RootToken)

	akvClient, err := akvclient.NewClient(akvclient.NewConfig(
		testCfg.AkvClient.VaultName,
		testCfg.AkvClient.TenantID,
		testCfg.AkvClient.ClientID,
		testCfg.AkvClient.ClientSecret,
	))
	if err != nil {
		logger.WithError(err).Error("cannot initialize akv client")
		return nil, err
	}

	ctx, cancel := context.WithCancel(ctx)
	return &IntegrationEnvironment{
		ctx:               ctx,
		logger:            logger,
		hashicorpClient:   hashicorpClient,
		akvClient:         akvClient,
		dockerClient:      dockerClient,
		keyManager:        keyManager,
		baseURL:           fmt.Sprintf("%s:%d", localhostPath, envHTTPPort),
		Cancel:            cancel,
		tmpManifestYaml:   tmpYml,
		tmpHashicorpToken: tmpTokenFile,
		cfg:               testCfg,
	}, nil
}

func (env *IntegrationEnvironment) Start(ctx context.Context) error {
	err := env.dockerClient.CreateNetwork(ctx, networkName)
	if err != nil {
		env.logger.WithError(err).Error("could not create network")
		return err
	}

	// Start Hashicorp Vault
	err = env.dockerClient.Up(ctx, hashicorpContainerID, networkName)
	if err != nil {
		env.logger.WithError(err).Error("could not up vault container")
		return err
	}

	err = env.dockerClient.WaitTillIsReady(ctx, hashicorpContainerID, 10*time.Second)
	if err != nil {
		env.logger.WithError(err).Error("could not start vault")
		return err
	}

	err = env.hashicorpClient.HealthCheck()
	if err != nil {
		env.logger.WithError(err).Error("failed to connect to hashicorp plugin")
		return err
	}

	err = env.hashicorpClient.Client().Sys().Mount("orchestrate", &api.MountInput{
		Type:        "plugin",
		Description: "Orchestrate Wallets",
		Config: api.MountConfigInput{
			ForceNoCache:              true,
			PassthroughRequestHeaders: []string{"X-Vault-Namespace"},
		},
		PluginName: utils.HashicorpPluginFilename,
	})
	if err != nil {
		env.logger.WithError(err).Error("failed to mount (enable) orchestrate vault plugin")
		return err
	}

	return nil
}

func (env *IntegrationEnvironment) Teardown(ctx context.Context) {
	env.logger.Info("tearing test suite down")

	err := env.keyManager.Stop(ctx)
	if err != nil {
		env.logger.WithError(err).Error("failed to stop key manager")
	}

	err = env.dockerClient.Down(ctx, hashicorpContainerID)
	if err != nil {
		env.logger.WithError(err).Error("could not down vault")
	}

	err = env.dockerClient.RemoveNetwork(ctx, networkName)
	if err != nil {
		env.logger.WithError(err).Error("could not remove network")
	}

	err = os.Remove(env.tmpManifestYaml)
	if err != nil {
		env.logger.WithError(err).Error("cannot remove temporary manifest yml file")
	}

	err = os.Remove(env.tmpHashicorpToken)
	if err != nil {
		env.logger.WithError(err).Error("cannot remove temporary hashicorp token file")
	}
}

func newTmpFile(data interface{}) (string, error) {
	file, err := ioutil.TempFile(os.TempDir(), "acceptanceTest_")
	if err != nil {
		return "", err
	}
	defer file.Close()

	bData, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}

	_, err = file.Write(bData)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func newTmpManifestYml(manifests ...*manifest.Manifest) (string, error) {
	file, err := ioutil.TempFile(os.TempDir(), "acceptanceTest_")
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := yaml.Marshal(manifests)
	if err != nil {
		return "", err
	}

	_, err = file.Write(data)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func newKeyManager(cfg *keymanager.Config, logger *log.Logger) (*keymanager.App, error) {
	return keymanager.New(cfg, logger)
}
