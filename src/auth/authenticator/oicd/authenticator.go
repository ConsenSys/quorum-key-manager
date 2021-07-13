package oicd

import (
	"net/http"
	"strings"

	"github.com/consensys/quorum-key-manager/pkg/errors"
	"github.com/consensys/quorum-key-manager/src/auth/authenticator"
	"github.com/consensys/quorum-key-manager/src/auth/types"
)

const AuthMode = "JWT"
const BearerSchema = "Bearer"

type Authenticator struct {
	jwtChecker *JWTChecker
}

var _ authenticator.Authenticator = Authenticator{}

func NewAuthenticator(cfg *Config) (*Authenticator, error) {
	if len(cfg.Certificates) == 0 {
		return nil, nil
	}

	auth := &Authenticator{
		jwtChecker: NewJWTChecker(cfg.Certificates, cfg.Claims, false),
	}

	return auth, nil
}

func (a Authenticator) Authenticate(req *http.Request) (*types.UserInfo, error) {
	// Extract Access Token from context
	token, ok := extractToken(BearerSchema, req.Header.Get("Authorization"))
	if !ok {
		return nil, nil
	}

	claims, err := a.jwtChecker.Check(req.Context(), token)
	if err != nil {
		return nil, errors.UnauthorizedError(err.Error())
	}

	return &types.UserInfo{
		Username: claims.Username,
		Groups:   claims.Groups,
		AuthMode: AuthMode,
	}, nil
}

func extractToken(prefix, auth string) (string, bool) {
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return "", false
	}

	return auth[len(prefix)+1:], true
}
