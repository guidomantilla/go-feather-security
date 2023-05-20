package auth

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewDefaultAuthenticationEndpoint(t *testing.T) {
	type args struct {
		authenticationService AuthenticationService
	}
	tests := []struct {
		name string
		args args
		want *DefaultAuthenticationEndpoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultAuthenticationEndpoint(tt.args.authenticationService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultAuthenticationEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultAuthenticationEndpoint_Authenticate(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name     string
		endpoint *DefaultAuthenticationEndpoint
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.endpoint.Authenticate(tt.args.ctx)
		})
	}
}
