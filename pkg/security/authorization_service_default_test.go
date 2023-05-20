package security

import (
	"context"
	"reflect"
	"testing"
)

func TestNewDefaultAuthorizationService(t *testing.T) {
	type args struct {
		tokenManager          TokenManager
		authorizationDelegate AuthorizationDelegate
	}
	tests := []struct {
		name string
		args args
		want *DefaultAuthorizationService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultAuthorizationService(tt.args.tokenManager, tt.args.authorizationDelegate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultAuthorizationService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultAuthorizationService_Authorize(t *testing.T) {
	type args struct {
		ctx         context.Context
		tokenString string
	}
	tests := []struct {
		name    string
		service *DefaultAuthorizationService
		args    args
		want    *Principal
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.service.Authorize(tt.args.ctx, tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultAuthorizationService.Authorize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultAuthorizationService.Authorize() = %v, want %v", got, tt.want)
			}
		})
	}
}
