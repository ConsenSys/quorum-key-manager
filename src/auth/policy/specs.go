package policy

import (
	"github.com/consensys/quorum-key-manager/src/auth/types"
	manifest "github.com/consensys/quorum-key-manager/src/manifests/types"
)

//the: make it const to avoid external abuse
var GroupKind manifest.Kind = "Group"

type GroupSpecs struct {
	Policies []string `json:"policies"`
}

//the: make it const to avoid external abuse
var Kind manifest.Kind = "Policy"

type Specs struct {
	Statements []*types.Statement `json:"statements"`
}
