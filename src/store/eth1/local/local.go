package local

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/ethereum"
	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/log"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/database"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/eth1"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/keys"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"

	"github.com/ConsenSysQuorum/quorum-key-manager/pkg/errors"
	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/entities"
)

const eip712DomainLabel = "EIP712Domain"

var eth1KeyAlgo = &entities.Algorithm{
	Type:          entities.Ecdsa,
	EllipticCurve: entities.Secp256k1,
}

type Store struct {
	keyStore     keys.Store
	eth1Accounts database.ETH1Accounts
	logger       *log.Logger
}

var _ eth1.Store = &Store{}

func New(keyStore keys.Store, eth1Accounts database.ETH1Accounts, logger *log.Logger) *Store {
	return &Store{
		keyStore:     keyStore,
		logger:       logger,
		eth1Accounts: eth1Accounts,
	}
}

func (s *Store) Info(context.Context) (*entities.StoreInfo, error) {
	return nil, errors.ErrNotImplemented
}

func (s *Store) Create(ctx context.Context, id string, attr *entities.Attributes) (*entities.ETH1Account, error) {
	key, err := s.keyStore.Create(ctx, id, eth1KeyAlgo, attr)
	if err != nil {
		return nil, err
	}

	acc, err := parseKey(key)
	if err != nil {
		return nil, err
	}

	err = s.eth1Accounts.Add(ctx, acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (s *Store) Import(ctx context.Context, id string, privKey []byte, attr *entities.Attributes) (*entities.ETH1Account, error) {
	key, err := s.keyStore.Import(ctx, id, privKey, eth1KeyAlgo, attr)
	if err != nil {
		return nil, err
	}

	acc, err := parseKey(key)
	if err != nil {
		return nil, err
	}

	err = s.eth1Accounts.Add(ctx, acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (s *Store) Get(ctx context.Context, addr string) (*entities.ETH1Account, error) {
	return s.eth1Accounts.Get(ctx, addr)
}

func (s *Store) GetAll(ctx context.Context) ([]*entities.ETH1Account, error) {
	return s.eth1Accounts.GetAll(ctx)
}

func (s *Store) List(ctx context.Context) ([]string, error) {
	addresses := []string{}
	accounts, err := s.eth1Accounts.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		addresses = append(addresses, account.Address)
	}

	return addresses, nil
}

func (s *Store) Update(ctx context.Context, addr string, attr *entities.Attributes) (*entities.ETH1Account, error) {
	account, err := s.eth1Accounts.Get(ctx, addr)
	if err != nil {
		return nil, err
	}

	key, err := s.keyStore.Update(ctx, account.ID, attr)
	if err != nil {
		return nil, err
	}

	acc, err := parseKey(key)
	if err != nil {
		return nil, err
	}

	err = s.eth1Accounts.Add(ctx, acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (s *Store) Delete(ctx context.Context, addr string) error {
	account, err := s.eth1Accounts.Get(ctx, addr)
	if err != nil {
		return err
	}

	err = s.keyStore.Delete(ctx, account.ID)
	if err != nil {
		return err
	}

	err = s.eth1Accounts.Remove(ctx, addr)
	if err != nil {
		return err
	}

	return s.eth1Accounts.AddDeleted(ctx, account)
}

func (s *Store) GetDeleted(ctx context.Context, addr string) (*entities.ETH1Account, error) {
	return s.eth1Accounts.GetDeleted(ctx, addr)
}

func (s *Store) ListDeleted(ctx context.Context) ([]string, error) {
	addresses := []string{}
	accounts, err := s.eth1Accounts.GetAllDeleted(ctx)
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		addresses = append(addresses, account.Address)
	}

	return addresses, nil
}

func (s *Store) Undelete(ctx context.Context, addr string) error {
	account, err := s.eth1Accounts.GetDeleted(ctx, addr)
	if err != nil {
		return err
	}

	err = s.keyStore.Undelete(ctx, account.ID)
	if err != nil {
		return err
	}

	err = s.eth1Accounts.RemoveDeleted(ctx, addr)
	if err != nil {
		return err
	}

	return s.eth1Accounts.Add(ctx, account)
}

func (s *Store) Destroy(ctx context.Context, addr string) error {
	account, err := s.eth1Accounts.GetDeleted(ctx, addr)
	if err != nil {
		return err
	}

	err = s.keyStore.Destroy(ctx, account.ID)
	if err != nil {
		return err
	}

	return s.eth1Accounts.RemoveDeleted(ctx, addr)
}

func (s *Store) Sign(ctx context.Context, addr string, data []byte) ([]byte, error) {
	account, err := s.Get(ctx, addr)
	if err != nil {
		return nil, err
	}

	signature, err := s.keyStore.Sign(ctx, account.ID, data)
	if err != nil {
		return nil, err
	}

	return s.appendRecID(data, signature, account.PublicKey)
}

func (s *Store) SignTypedData(ctx context.Context, addr string, typedData *core.TypedData) ([]byte, error) {
	encodedData, err := getEIP712EncodedData(typedData)
	if err != nil {
		return nil, err
	}

	account, err := s.Get(ctx, addr)
	if err != nil {
		return nil, err
	}

	data := crypto.Keccak256([]byte(encodedData))
	signature, err := s.Sign(ctx, addr, data)
	if err != nil {
		return nil, err
	}

	return s.appendRecID(data, signature, account.PublicKey)
}

func (s *Store) SignTransaction(ctx context.Context, addr string, chainID *big.Int, tx *types.Transaction) ([]byte, error) {
	logger := s.logger.WithField("address", addr)

	signer := types.NewEIP155Signer(chainID)
	txData := signer.Hash(tx).Bytes()
	signature, err := s.Sign(ctx, addr, txData)
	if err != nil {
		return nil, err
	}

	ethSignature, err := appendSignatureV(tx, signature, signer)
	if err != nil {
		errMessage := "failed to generate transaction signature"
		logger.WithError(err).Error(errMessage)
		return nil, errors.EncodingError(errMessage)
	}

	return ethSignature, nil
}

func (s *Store) SignEEA(ctx context.Context, addr string, chainID *big.Int, tx *ethereum.EEATxData, args *ethereum.PrivateArgs) ([]byte, error) {
	return nil, errors.ErrNotImplemented
}

func (s *Store) SignPrivate(ctx context.Context, addr string, tx *types.Transaction) ([]byte, error) {
	return nil, errors.ErrNotImplemented
}

func (s *Store) ECRevocer(_ context.Context, data, sig []byte) (string, error) {
	pubKey, err := crypto.SigToPub(data, sig)
	if err != nil {
		s.logger.WithError(err).Error("failed to recover public key")
		return "", errors.InvalidParameterError("failed to recover public key, please verify your signature and payload")
	}

	return crypto.PubkeyToAddress(*pubKey).Hex(), nil
}

func (s *Store) Verify(ctx context.Context, addr string, data, sig []byte) error {
	recoveredAddress, err := s.ECRevocer(ctx, data, sig)
	if err != nil {
		return err
	}

	if addr != recoveredAddress {
		return errors.InvalidParameterError("failed to verify signature: recovered address does not match the expected one or payload is malformed")
	}

	return nil
}

func (s *Store) VerifyTypedData(ctx context.Context, addr string, sig []byte, typedData *core.TypedData) error {
	encodedData, err := getEIP712EncodedData(typedData)
	if err != nil {
		return err
	}

	return s.Verify(ctx, addr, sig, []byte(encodedData))
}

func (s *Store) Encrypt(ctx context.Context, addr string, data []byte) ([]byte, error) {
	account, err := s.eth1Accounts.Get(ctx, addr)
	if err != nil {
		return nil, err
	}

	return s.keyStore.Encrypt(ctx, account.ID, data)
}

func (s *Store) Decrypt(ctx context.Context, addr string, data []byte) ([]byte, error) {
	account, err := s.eth1Accounts.Get(ctx, addr)
	if err != nil {
		return nil, err
	}

	return s.keyStore.Decrypt(ctx, account.ID, data)
}

// To understand this function, please read: http://coders-errand.com/ecrecover-signature-verification-ethereum/
func (s *Store) appendRecID(data, sig, pubKey []byte) ([]byte, error) {
	for _, recID := range []byte{0, 1, 2, 3} {
		appendedSignature := append(sig, recID)
		recoveredPubKey, err := crypto.SigToPub(data, appendedSignature)
		if err != nil {
			errMessage := "failed to recover public key candidate with appended recID"
			s.logger.WithField("recID", recID).WithError(err).Error(errMessage)
			return nil, errors.InvalidParameterError(errMessage)
		}

		if bytes.Equal(crypto.FromECDSAPub(recoveredPubKey), pubKey) {
			return appendedSignature, nil
		}
	}

	errMessage := "failed to compute recovery ID"
	s.logger.Error(errMessage)
	return nil, errors.DependencyFailureError(errMessage)
}

func getEIP712EncodedData(typedData *core.TypedData) (string, error) {
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return "", errors.InvalidParameterError("invalid typed data message")
	}

	domainSeparatorHash, err := typedData.HashStruct(eip712DomainLabel, typedData.Domain.Map())
	if err != nil {
		return "", errors.InvalidParameterError("invalid domain separator")
	}

	return fmt.Sprintf("\x19\x01%s%s", domainSeparatorHash, typedDataHash), nil
}

func appendSignatureV(tx *types.Transaction, sig []byte, signer types.Signer) ([]byte, error) {
	r, s, v, err := signer.SignatureValues(tx, sig)
	if err != nil {
		return nil, err
	}

	return append(append(r.Bytes(), s.Bytes()...), v.Bytes()...), nil
}
