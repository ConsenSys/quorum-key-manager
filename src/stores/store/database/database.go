package database

import (
	"context"
	models2 "github.com/consensys/quorum-key-manager/src/stores/store/models"
)

//go:generate mockgen -source=database.go -destination=mock/database.go -package=mock

type Database interface {
	ETH1Accounts() ETH1Accounts
	Keys() Keys
	RunInTransaction(ctx context.Context, persistFunc func(dbtx Database) error) error
}

type ETH1Accounts interface {
	Get(ctx context.Context, addr string) (*models2.ETH1Account, error)
	GetDeleted(ctx context.Context, addr string) (*models2.ETH1Account, error)
	GetAll(ctx context.Context) ([]*models2.ETH1Account, error)
	GetAllDeleted(ctx context.Context) ([]*models2.ETH1Account, error)
	Add(ctx context.Context, account *models2.ETH1Account) error
	Update(ctx context.Context, account *models2.ETH1Account) error
	Delete(ctx context.Context, addr string) error
	Restore(ctx context.Context, account *models2.ETH1Account) error
	Purge(ctx context.Context, addr string) error
}

type Keys interface {
	Get(ctx context.Context, id string) (*models2.Key, error)
	GetDeleted(ctx context.Context, id string) (*models2.Key, error)
	GetAll(ctx context.Context) ([]*models2.Key, error)
	GetAllDeleted(ctx context.Context) ([]*models2.Key, error)
	Add(ctx context.Context, key *models2.Key) error
	Update(ctx context.Context, key *models2.Key) error
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, key *models2.Key) error
	Purge(ctx context.Context, id string) error
}
