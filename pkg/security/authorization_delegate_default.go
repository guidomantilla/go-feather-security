package security

import (
	"context"

	"go.uber.org/zap"
)

type DefaultAuthorizationDelegate struct {
	principalManager PrincipalManager
}

func NewDefaultAuthorizationDelegate(principalManager PrincipalManager) *DefaultAuthorizationDelegate {

	if principalManager == nil {
		zap.L().Fatal("starting up - error setting up authorizationDelegate: principalManager is nil")
	}

	return &DefaultAuthorizationDelegate{
		principalManager: principalManager,
	}
}

func (delegate *DefaultAuthorizationDelegate) Authorize(ctx context.Context, principal *Principal) error {

	var err error
	var user *Principal
	if user, err = delegate.principalManager.Find(ctx, *principal.Username); err != nil {
		return ErrFailedAuthorization
	}

	if *(user.Role) != *(principal.Role) {
		return ErrFailedAuthorization
	}

	var value any
	if value = ctx.Value(ResourceCtxKey{}); value == nil {
		return ErrFailedAuthorization
	}

	if err = delegate.principalManager.VerifyResource(ctx, *user.Username, value.(string)); err != nil {
		return ErrFailedAuthorization
	}

	return nil
}
