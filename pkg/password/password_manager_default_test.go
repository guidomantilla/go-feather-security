package password

import (
	"reflect"
	"testing"
)

func TestNewDefaultPasswordManager(t *testing.T) {
	type args struct {
		passwordEncoder   PasswordEncoder
		passwordGenerator PasswordGenerator
	}
	tests := []struct {
		name string
		args args
		want *DefaultPasswordManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultPasswordManager(tt.args.passwordEncoder, tt.args.passwordGenerator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultPasswordManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultPasswordManager_Encode(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name    string
		manager *DefaultPasswordManager
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.manager.Encode(tt.args.rawPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultPasswordManager.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultPasswordManager.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultPasswordManager_Matches(t *testing.T) {
	type args struct {
		encodedPassword string
		rawPassword     string
	}
	tests := []struct {
		name    string
		manager *DefaultPasswordManager
		args    args
		want    *bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.manager.Matches(tt.args.encodedPassword, tt.args.rawPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultPasswordManager.Matches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultPasswordManager.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultPasswordManager_UpgradeEncoding(t *testing.T) {
	type args struct {
		encodedPassword string
	}
	tests := []struct {
		name    string
		manager *DefaultPasswordManager
		args    args
		want    *bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.manager.UpgradeEncoding(tt.args.encodedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultPasswordManager.UpgradeEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultPasswordManager.UpgradeEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultPasswordManager_Generate(t *testing.T) {
	tests := []struct {
		name    string
		manager *DefaultPasswordManager
		want    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.manager.Generate(); got != tt.want {
				t.Errorf("DefaultPasswordManager.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultPasswordManager_Validate(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name    string
		manager *DefaultPasswordManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.manager.Validate(tt.args.rawPassword); (err != nil) != tt.wantErr {
				t.Errorf("DefaultPasswordManager.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
