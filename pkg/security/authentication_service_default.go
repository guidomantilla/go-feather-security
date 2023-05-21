package security

import (
	"context"

	"go.uber.org/zap"
)

type DefaultAuthenticationService struct {
	tokenManager           TokenManager
	authenticationDelegate AuthenticationService
}

func NewDefaultAuthenticationService(tokenManager TokenManager, authenticationDelegate AuthenticationService) *DefaultAuthenticationService {

	if tokenManager == nil {
		zap.L().Fatal("starting up - error setting up authenticationService: tokenManager is nil")
	}

	if authenticationDelegate == nil {
		zap.L().Fatal("starting up - error setting up authenticationService: authenticationDelegate is nil")
	}

	return &DefaultAuthenticationService{
		tokenManager:           tokenManager,
		authenticationDelegate: authenticationDelegate,
	}
}

func (service *DefaultAuthenticationService) Authenticate(ctx context.Context, principal *Principal) error {

	var err error
	if err = service.authenticationDelegate.Authenticate(ctx, principal); err != nil {
		return err
	}

	principal.Password = nil
	if principal.Token, err = service.tokenManager.Generate(principal); err != nil {
		return err
	}

	return nil
}
