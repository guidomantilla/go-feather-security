package security

import (
	"context"
	"reflect"
	"testing"
)

func TestNewDefaultAuthenticationService(t *testing.T) {
	type args struct {
		tokenManager           TokenManager
		authenticationDelegate AuthenticationDelegate
	}
	tests := []struct {
		name string
		args args
		want *DefaultAuthenticationService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultAuthenticationService(tt.args.tokenManager, tt.args.authenticationDelegate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultAuthenticationService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultAuthenticationService_Authenticate(t *testing.T) {
	type args struct {
		ctx       context.Context
		principal *Principal
	}
	tests := []struct {
		name    string
		service *DefaultAuthenticationService
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.service.Authenticate(tt.args.ctx, tt.args.principal)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultAuthenticationService.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultAuthenticationService.Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}
