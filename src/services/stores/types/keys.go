package types

import (
	manifest "github.com/ConsenSysQuorum/quorum-key-manager/src/services/manifests/types"
)

const (
	HashicorpKeys manifest.Kind = "HashicorpKeys"
	AKVKeys       manifest.Kind = "AKVKeys"
	AWSKeys       manifest.Kind = "AWSKeys"
	KMSKeys       manifest.Kind = "KMSKeys"
)
