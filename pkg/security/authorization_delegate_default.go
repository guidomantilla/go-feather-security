package security

import (
	"context"
	"reflect"

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

	if !reflect.DeepEqual(user.Authorities, principal.Authorities) {
		return ErrFailedAuthorization
	}

	return nil
}
