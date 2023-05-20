package security

import (
	"context"

	"go.uber.org/zap"
)

type DefaultAuthorizationService struct {
	tokenManager          TokenManager
	authorizationDelegate AuthorizationDelegate
}

func NewDefaultAuthorizationService(tokenManager TokenManager, authorizationDelegate AuthorizationDelegate) *DefaultAuthorizationService {

	if tokenManager == nil {
		zap.L().Fatal("starting up - error setting up authorization service: authorization delegate is nil")
	}

	if authorizationDelegate == nil {
		zap.L().Fatal("starting up - error setting up authorization service: authorizationDelegate is nil")
	}

	return &DefaultAuthorizationService{
		tokenManager:          tokenManager,
		authorizationDelegate: authorizationDelegate,
	}
}

func (service *DefaultAuthorizationService) Authorize(ctx context.Context, tokenString string) (*Principal, error) {

	var err error
	var principal *Principal
	if principal, err = service.tokenManager.Validate(tokenString); err != nil {
		return nil, err
	}

	if err = service.authorizationDelegate.Authorize(ctx, principal); err != nil {
		return nil, err
	}

	return principal, nil
}
