package security

import (
	"context"
	"log/slog"
	"os"
)

type DefaultAuthenticationService struct {
	passwordEncoder  PasswordEncoder
	principalManager PrincipalManager
	tokenManager     TokenManager
}

func NewDefaultAuthenticationService(passwordEncoder PasswordEncoder, principalManager PrincipalManager, tokenManager TokenManager) *DefaultAuthenticationService {

	if passwordEncoder == nil {
		slog.Error("starting up - error setting up authenticationService: passwordEncoder is nil")
		os.Exit(1)
	}

	if principalManager == nil {
		slog.Error("starting up - error setting up authenticationService: principalManager is nil")
		os.Exit(1)
	}

	if tokenManager == nil {
		slog.Error("starting up - error setting up authenticationService: tokenManager is nil")
		os.Exit(1)
	}

	return &DefaultAuthenticationService{
		passwordEncoder:  passwordEncoder,
		principalManager: principalManager,
		tokenManager:     tokenManager,
	}
}

func (service *DefaultAuthenticationService) Authenticate(ctx context.Context, principal *Principal) error {

	var err error
	var user *Principal
	if user, err = service.principalManager.Find(ctx, *principal.Username); err != nil {
		return ErrAuthenticationFailed(err)
	}

	var needsUpgrade *bool
	if needsUpgrade, err = service.passwordEncoder.UpgradeEncoding(*(user.Password)); err != nil || *(needsUpgrade) {
		return ErrAuthenticationFailed(ErrAccountExpiredPassword)
	}

	var matches *bool
	if matches, err = service.passwordEncoder.Matches(*(user.Password), *principal.Password); err != nil || !*(matches) {
		return ErrAuthenticationFailed(ErrAccountInvalidPassword)
	}

	principal.Password, principal.Passphrase = nil, nil
	principal.Role, principal.Resources = user.Role, user.Resources
	if principal.Token, err = service.tokenManager.Generate(principal); err != nil {
		return ErrAuthenticationFailed(err)
	}

	return nil
}
