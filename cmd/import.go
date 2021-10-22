package cmd

import (
	"fmt"

	"github.com/consensys/quorum-key-manager/cmd/flags"
	"github.com/consensys/quorum-key-manager/src/auth/types"
	"github.com/consensys/quorum-key-manager/src/infra/log"
	"github.com/consensys/quorum-key-manager/src/infra/log/zap"
	manifest "github.com/consensys/quorum-key-manager/src/infra/manifests/entities"
	manifestreader "github.com/consensys/quorum-key-manager/src/infra/manifests/filesystem"
	"github.com/consensys/quorum-key-manager/src/infra/postgres/client"
	storeservice "github.com/consensys/quorum-key-manager/src/stores"
	"github.com/consensys/quorum-key-manager/src/stores/connectors/stores"
	"github.com/consensys/quorum-key-manager/src/stores/database/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newImportCmd() *cobra.Command {
	var logger *zap.Logger
	var storesConnector storeservice.Stores
	var mnf *manifest.Manifest

	importCmd := &cobra.Command{
		Use:   "import",
		Short: "Import management tool",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error

			if logger, err = getLogger(); err != nil {
				return err
			}
			if storesConnector, err = getStores(logger); err != nil {
				return err
			}
			if mnf, err = getManifest(); err != nil {
				return err
			}

			return nil
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			syncZapLogger(logger)
		},
	}

	flags.PGFlags(importCmd.Flags())
	flags.ImportFlags(importCmd.Flags())
	flags.ManifestFlags(importCmd.Flags())

	importSecretsCmd := &cobra.Command{
		Use:   "secrets",
		Short: "import secrets from a vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := storesConnector.CreateSecret(ctx, mnf.Name, manifest.VaultType(mnf.Kind), mnf.Specs, mnf.AllowedTenants); err != nil {
				return err
			}

			return storesConnector.ImportSecrets(cmd.Context(), mnf.Name, types.WildcardUser)
		},
	}
	importCmd.AddCommand(importSecretsCmd)

	importKeysCmd := &cobra.Command{
		Use:   "keys",
		Short: "import keys from a vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := storesConnector.CreateKey(ctx, mnf.Name, manifest.VaultType(mnf.Kind), mnf.Specs, mnf.AllowedTenants); err != nil {
				return err
			}

			return storesConnector.ImportKeys(cmd.Context(), mnf.Name, types.WildcardUser)
		},
	}
	importCmd.AddCommand(importKeysCmd)

	importEthereumCmd := &cobra.Command{
		Use:   "ethereum",
		Short: "import ethereum accounts from a vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := storesConnector.CreateEthereum(ctx, mnf.Name, manifest.VaultType(mnf.Kind), mnf.Specs, mnf.AllowedTenants); err != nil {
				return err
			}

			return storesConnector.ImportEthereum(cmd.Context(), mnf.Name, types.WildcardUser)
		},
	}
	importCmd.AddCommand(importEthereumCmd)

	return importCmd
}

func getLogger() (*zap.Logger, error) {
	return zap.NewLogger(flags.NewLoggerConfig(viper.GetViper()))
}

func getStores(logger log.Logger) (storeservice.Stores, error) {
	// Create Postgres DB
	postgresClient, err := client.New(flags.NewPostgresConfig(viper.GetViper()))
	if err != nil {
		return nil, err
	}

	return stores.NewConnector(nil, postgres.New(logger, postgresClient), logger), nil
}

func getManifest() (*manifest.Manifest, error) {
	vipr := viper.GetViper()
	// Get manifests
	manifestReader, err := manifestreader.New(flags.NewManifestConfig(vipr))
	if err != nil {
		return nil, err
	}

	manifests, err := manifestReader.Load()
	if err != nil {
		return nil, err
	}

	storeName := flags.GetStoreName(vipr)

	for _, mnf := range manifests {
		// TODO: Filter on Load() function from reader when Kind Store implemented
		if mnf.Kind == manifest.Role || mnf.Kind == manifest.Node {
			continue
		}

		if mnf.Name == storeName {
			return mnf, nil
		}
	}

	return nil, fmt.Errorf("inexistent store %s in the manifests definitions", storeName)
}
