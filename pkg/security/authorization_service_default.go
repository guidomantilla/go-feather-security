package security

import (
	"context"

	"go.uber.org/zap"
)

type DefaultAuthorizationService struct {
	tokenManager     TokenManager
	principalManager PrincipalManager
}

func NewDefaultAuthorizationService(tokenManager TokenManager, principalManager PrincipalManager) *DefaultAuthorizationService {

	if tokenManager == nil {
		zap.L().Fatal("starting up - error setting up authorization service: authorization delegate is nil")
	}

	if principalManager == nil {
		zap.L().Fatal("starting up - error setting up authorization service:  principalManager is nil")
	}

	return &DefaultAuthorizationService{
		tokenManager:     tokenManager,
		principalManager: principalManager,
	}
}

func (service *DefaultAuthorizationService) Authorize(ctx context.Context, tokenString string) (*Principal, error) {

	var err error
	var principal *Principal
	if principal, err = service.tokenManager.Validate(tokenString); err != nil {
		return nil, err
	}

	var user *Principal
	if user, err = service.principalManager.Find(ctx, *principal.Username); err != nil {
		return nil, ErrAuthorizationFailed(err)
	}

	if *(user.Role) != *(principal.Role) {
		return nil, ErrAuthorizationFailed(ErrAccountInvalidRole)
	}

	var value any
	if value = ctx.Value(ResourceCtxKey{}); value == nil {
		return nil, ErrAuthorizationFailed(ErrAccountEmptyAuthorities)
	}

	principal.Password, principal.Passphrase, principal.Token = nil, nil, nil
	if err = service.principalManager.VerifyResource(ctx, *user.Username, value.(string)); err != nil {
		return nil, ErrAuthorizationFailed(err)
	}

	return principal, nil
}
