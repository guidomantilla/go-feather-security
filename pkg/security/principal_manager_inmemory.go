package security

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

type InMemoryPrincipalManager struct {
	repository      map[string]*Principal
	passwordManager PasswordManager
}

func NewInMemoryPrincipalManager(passwordManager PasswordManager) *InMemoryPrincipalManager {

	if passwordManager == nil {
		zap.L().Fatal("starting up - error setting up principalManager: passwordManager is nil")
	}

	return &InMemoryPrincipalManager{
		passwordManager: passwordManager,
		repository:      make(map[string]*Principal),
	}
}

func (manager *InMemoryPrincipalManager) Create(ctx context.Context, principal *Principal) error {

	var err error
	if err = manager.Exists(ctx, *principal.Username); err == nil {
		return errors.New("username already exists")
	}

	if err = manager.passwordManager.Validate(*principal.Password); err != nil {
		return err
	}

	if principal.Password, err = manager.passwordManager.Encode(*principal.Password); err != nil {
		return err
	}

	manager.repository[*principal.Username] = principal

	return nil
}

func (manager *InMemoryPrincipalManager) Update(ctx context.Context, principal *Principal) error {
	return manager.Create(ctx, principal)
}

func (manager *InMemoryPrincipalManager) Delete(_ context.Context, username string) error {
	delete(manager.repository, username)
	return nil
}

func (manager *InMemoryPrincipalManager) Find(_ context.Context, username string) (*Principal, error) {

	var ok bool
	var user *Principal
	if user, ok = manager.repository[username]; !ok {
		return nil, errors.New("username not found")
	}
	return user, nil
}

func (manager *InMemoryPrincipalManager) Exists(_ context.Context, username string) error {

	var ok bool
	if _, ok = manager.repository[username]; !ok {
		return errors.New("username not found")
	}
	return nil
}

func (manager *InMemoryPrincipalManager) ChangePassword(ctx context.Context, principal *Principal) error {

	var err error
	if err = manager.Exists(ctx, *principal.Username); err != nil {
		return err
	}

	if err = manager.passwordManager.Validate(*principal.Password); err != nil {
		return err
	}

	if principal.Password, err = manager.passwordManager.Encode(*principal.Password); err != nil {
		return err
	}

	manager.repository[*principal.Username] = principal

	return nil
}
