package hashicorp

import (
	hashicorp "github.com/hashicorp/vault/api"
)

//go:generate mockgen -source=hashicorp.go -destination=mocks/hashicorp.go -package=mocks

type VaultClient interface {
	Read(path string, data map[string][]string) (*hashicorp.Secret, error)
	Write(path string, data map[string]interface{}) (*hashicorp.Secret, error)
	WritePost(path string, data map[string][]string) error
	Delete(path string, data map[string][]string) error
	List(path string) (*hashicorp.Secret, error)
	SetToken(token string)
	UnwrapToken(token string) (*hashicorp.Secret, error)
	Mount(path string, mountInfo *hashicorp.MountInput) error
	HealthCheck() error
}
