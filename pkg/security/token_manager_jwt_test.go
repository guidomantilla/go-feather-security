package security

import (
	"reflect"
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func TestNewJwtTokenManager(t *testing.T) {
	type args struct {
		issuer        string
		timeout       time.Duration
		secretKey     any
		signingMethod jwt.SigningMethod
	}
	tests := []struct {
		name string
		args args
		want *JwtTokenManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJwtTokenManager(tt.args.issuer, tt.args.timeout, tt.args.secretKey, tt.args.signingMethod); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJwtTokenManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJwtTokenManager_Generate(t *testing.T) {
	type args struct {
		principal *Principal
	}
	tests := []struct {
		name    string
		manager *JwtTokenManager
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.manager.Generate(tt.args.principal)
			if (err != nil) != tt.wantErr {
				t.Errorf("JwtTokenManager.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JwtTokenManager.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJwtTokenManager_Validate(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		manager *JwtTokenManager
		args    args
		want    *Principal
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.manager.Validate(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("JwtTokenManager.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JwtTokenManager.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
