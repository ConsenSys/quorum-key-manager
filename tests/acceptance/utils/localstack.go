package utils

import (
	"context"

	dockerlocalstack "github.com/consensysquorum/quorum-key-manager/tests/acceptance/docker/config/localstack"
)

func LocalstackContainer(ctx context.Context) (*dockerlocalstack.Config, error) {

	localstackHost := "localhost"
	localstackPort := "4566"
	localstackServices := []string{"s3", "kms", "secretsmanager"}

	vaultContainer := dockerlocalstack.
		NewDefault().
		SetHostPort(localstackPort).
		SetHost(localstackHost).
		SetServices(localstackServices)

	return vaultContainer, nil
}
