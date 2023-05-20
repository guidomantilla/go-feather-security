package security

import (
	"context"

	"go.uber.org/zap"
)

type DefaultAuthenticationService struct {
	tokenManager           TokenManager
	authenticationDelegate AuthenticationDelegate
}

func NewDefaultAuthenticationService(tokenManager TokenManager, authenticationDelegate AuthenticationDelegate) *DefaultAuthenticationService {

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

func (service *DefaultAuthenticationService) Authenticate(ctx context.Context, principal *Principal) (*string, error) {

	var err error
	if err = service.authenticationDelegate.Authenticate(ctx, principal); err != nil {
		return nil, err
	}

	var token *string
	if token, err = service.tokenManager.Generate(principal); err != nil {
		return nil, err
	}

	return token, nil
}