package security

import (
	"context"

	"go.uber.org/zap"
)

type DefaultAuthenticationDelegate struct {
	passwordEncoder  PasswordEncoder
	principalManager PrincipalManager
}

func NewDefaultAuthenticationDelegate(passwordEncoder PasswordEncoder, principalManager PrincipalManager) *DefaultAuthenticationDelegate {

	if passwordEncoder == nil {
		zap.L().Fatal("starting up - error setting up authenticationDelegate: passwordEncoder is nil")
	}

	if principalManager == nil {
		zap.L().Fatal("starting up - error setting up authenticationDelegate: principalManager is nil")
	}

	return &DefaultAuthenticationDelegate{
		passwordEncoder:  passwordEncoder,
		principalManager: principalManager,
	}
}

func (delegate *DefaultAuthenticationDelegate) Authenticate(ctx context.Context, principal *Principal) error {

	var err error
	var user *Principal
	if user, err = delegate.principalManager.Find(ctx, *principal.Username); err != nil {
		return ErrFailedAuthentication
	}

	if user.Password == nil || *(user.Password) == "" {
		return ErrFailedAuthentication
	}

	var matches *bool
	if matches, err = delegate.passwordEncoder.Matches(*(user.Password), *principal.Password); err != nil || !*(matches) {
		return ErrFailedAuthentication
	}

	var needsUpgrade *bool
	if needsUpgrade, err = delegate.passwordEncoder.UpgradeEncoding(*(user.Password)); err != nil || *(needsUpgrade) {
		return ErrFailedAuthentication
	}

	principal.Password = nil
	principal.Authorities = user.Authorities

	return nil
}
