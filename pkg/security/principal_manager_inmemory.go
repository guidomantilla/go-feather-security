package security

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

type InMemoryPrincipalManager struct {
	principalRepo   map[string]*Principal
	resourceRepo    map[string]map[string]string
	passwordManager PasswordManager
}

func NewInMemoryPrincipalManager(passwordManager PasswordManager) *InMemoryPrincipalManager {

	if passwordManager == nil {
		zap.L().Fatal("starting up - error setting up principalManager: passwordManager is nil")
	}

	return &InMemoryPrincipalManager{
		passwordManager: passwordManager,
		principalRepo:   make(map[string]*Principal),
		resourceRepo:    make(map[string]map[string]string),
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

	manager.principalRepo[*principal.Username] = principal
	manager.resourceRepo[*principal.Username] = make(map[string]string)

	for _, authority := range principal.Authorities {
		for _, resource := range authority.Resources {
			manager.resourceRepo[*principal.Username][resource] = resource
		}
	}

	return nil
}

func (manager *InMemoryPrincipalManager) Update(ctx context.Context, principal *Principal) error {
	return manager.Create(ctx, principal)
}

func (manager *InMemoryPrincipalManager) Delete(_ context.Context, username string) error {
	delete(manager.principalRepo, username)
	delete(manager.resourceRepo, username)
	return nil
}

func (manager *InMemoryPrincipalManager) Find(_ context.Context, username string) (*Principal, error) {

	var ok bool
	var user *Principal
	if user, ok = manager.principalRepo[username]; !ok {
		return nil, errors.New("username not found")
	}
	return user, nil
}

func (manager *InMemoryPrincipalManager) Exists(_ context.Context, username string) error {

	var ok bool
	if _, ok = manager.principalRepo[username]; !ok {
		return errors.New("username not found")
	}
	return nil
}

func (manager *InMemoryPrincipalManager) ChangePassword(ctx context.Context, username string, password string) error {

	var err error
	var user *Principal
	if user, err = manager.Find(ctx, username); err != nil {
		return err
	}

	if err = manager.passwordManager.Validate(password); err != nil {
		return err
	}

	if user.Password, err = manager.passwordManager.Encode(password); err != nil {
		return err
	}

	return nil
}

func (manager *InMemoryPrincipalManager) VerifyResource(ctx context.Context, username string, resource string) error {

	var err error
	if err = manager.Exists(ctx, username); err != nil {
		return err
	}

	if _, ok := manager.resourceRepo[username][resource]; !ok {
		return errors.New("resource not found")
	}

	return nil
}
