package authorizator

import (
	"fmt"

	"github.com/consensys/quorum-key-manager/pkg/errors"
	"github.com/consensys/quorum-key-manager/src/infra/log"

	"github.com/consensys/quorum-key-manager/src/auth/types"
)

type Authorizator struct {
	logger      log.Logger
	permissions map[types.Permission]bool // We use a map to avoid iterating an array, the boolean is irrelevant and always true
	tenant      string
}

func New(permissions []types.Permission, tenant string, logger log.Logger) *Authorizator {
	pMap := map[types.Permission]bool{}
	for _, p := range permissions {
		pMap[p] = true
	}

	return &Authorizator{
		permissions: pMap,
		tenant:      tenant,
		logger:      logger,
	}
}

func (auth *Authorizator) CheckPermission(ops ...*types.Operation) error {
	for _, op := range ops {
		permission := buildPermission(op.Action, op.Resource)
		if _, ok := auth.permissions[permission]; !ok {
			errMessage := "user is not authorized to perform this operation"
			auth.logger.With("permission", permission).Error(errMessage)
			return errors.ForbiddenError(errMessage)
		}
	}

	return nil
}

func (auth *Authorizator) CheckAccess(allowedTenants []string) error {
	if len(allowedTenants) == 0 {
		return nil
	}

	if auth.tenant == "" {
		errMessage := "missing tenant in credentials"
		auth.logger.Error(errMessage)
		return errors.UnauthorizedError(errMessage)
	}

	for _, t := range allowedTenants {
		if t == auth.tenant {
			return nil
		}
	}

	errMessage := "resource not found"
	auth.logger.With("tenant", auth.tenant, "allowed_tenants", allowedTenants).Error(errMessage)
	return errors.NotFoundError(errMessage)
}

func buildPermission(action types.OpAction, resource types.OpResource) types.Permission {
	return types.Permission(fmt.Sprintf("%s:%s", action, resource))
}
