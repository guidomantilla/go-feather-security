package auth

import (
	"context"
	"reflect"
	"testing"

	"github.com/guidomantilla/go-feather-security/pkg/password"
)

func TestNewInMemoryPrincipalManager(t *testing.T) {
	type args struct {
		passwordManager password.PasswordManager
	}
	tests := []struct {
		name string
		args args
		want *InMemoryPrincipalManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemoryPrincipalManager(tt.args.passwordManager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemoryPrincipalManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryPrincipalManager_Create(t *testing.T) {
	type args struct {
		ctx       context.Context
		principal *Principal
	}
	tests := []struct {
		name    string
		manager *InMemoryPrincipalManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.manager.Create(tt.args.ctx, tt.args.principal); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryPrincipalManager.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryPrincipalManager_Update(t *testing.T) {
	type args struct {
		ctx       context.Context
		principal *Principal
	}
	tests := []struct {
		name    string
		manager *InMemoryPrincipalManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.manager.Update(tt.args.ctx, tt.args.principal); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryPrincipalManager.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryPrincipalManager_Delete(t *testing.T) {
	type args struct {
		in0      context.Context
		username string
	}
	tests := []struct {
		name    string
		manager *InMemoryPrincipalManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.manager.Delete(tt.args.in0, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryPrincipalManager.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryPrincipalManager_Find(t *testing.T) {
	type args struct {
		in0      context.Context
		username string
	}
	tests := []struct {
		name    string
		manager *InMemoryPrincipalManager
		args    args
		want    *Principal
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.manager.Find(tt.args.in0, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryPrincipalManager.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemoryPrincipalManager.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryPrincipalManager_Exists(t *testing.T) {
	type args struct {
		in0      context.Context
		username string
	}
	tests := []struct {
		name    string
		manager *InMemoryPrincipalManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.manager.Exists(tt.args.in0, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryPrincipalManager.Exists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryPrincipalManager_ChangePassword(t *testing.T) {
	type args struct {
		ctx       context.Context
		principal *Principal
	}
	tests := []struct {
		name    string
		manager *InMemoryPrincipalManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.manager.ChangePassword(tt.args.ctx, tt.args.principal); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryPrincipalManager.ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryPrincipalManager_Authenticate(t *testing.T) {
	type args struct {
		ctx       context.Context
		principal *Principal
	}
	tests := []struct {
		name    string
		manager *InMemoryPrincipalManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.manager.Authenticate(tt.args.ctx, tt.args.principal); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryPrincipalManager.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryPrincipalManager_Authorize(t *testing.T) {
	type args struct {
		ctx       context.Context
		principal *Principal
	}
	tests := []struct {
		name    string
		manager *InMemoryPrincipalManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.manager.Authorize(tt.args.ctx, tt.args.principal); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryPrincipalManager.Authorize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
