package stores

import (
	manifest "github.com/consensys/quorum-key-manager/src/manifests/types"
)

const (
	Ethereum manifest.Kind = "Ethereum"
)

const (
	HashicorpKeys manifest.Kind = "HashicorpKeys"
	AKVKeys       manifest.Kind = "AKVKeys"
	AWSKeys       manifest.Kind = "AWSKeys"
	LocalKeys     manifest.Kind = "LocalKeys"
)

const (
	HashicorpSecrets manifest.Kind = "HashicorpSecrets"
	AKVSecrets       manifest.Kind = "AKVSecrets"
	AWSSecrets       manifest.Kind = "AWSSecrets"
)
