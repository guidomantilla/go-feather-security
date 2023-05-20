package security

import (
	"reflect"
	"testing"
)

func TestNewPbkdf2PasswordEncoder(t *testing.T) {
	type args struct {
		options []Pbkdf2PasswordEncoderOption
	}
	tests := []struct {
		name string
		args args
		want *Pbkdf2PasswordEncoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPbkdf2PasswordEncoder(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPbkdf2PasswordEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPbkdf2Iterations(t *testing.T) {
	type args struct {
		iterations int
	}
	tests := []struct {
		name string
		args args
		want Pbkdf2PasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPbkdf2Iterations(tt.args.iterations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPbkdf2Iterations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPbkdf2SaltLength(t *testing.T) {
	type args struct {
		saltLength int
	}
	tests := []struct {
		name string
		args args
		want Pbkdf2PasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPbkdf2SaltLength(tt.args.saltLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPbkdf2SaltLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPbkdf2KeyLength(t *testing.T) {
	type args struct {
		keyLength int
	}
	tests := []struct {
		name string
		args args
		want Pbkdf2PasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPbkdf2KeyLength(tt.args.keyLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPbkdf2KeyLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithHashFunc(t *testing.T) {
	type args struct {
		hashFunc HashFunc
	}
	tests := []struct {
		name string
		args args
		want Pbkdf2PasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithHashFunc(tt.args.hashFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHashFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPbkdf2PasswordEncoder_Encode(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name    string
		encoder *Pbkdf2PasswordEncoder
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.encoder.Encode(tt.args.rawPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pbkdf2PasswordEncoder.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pbkdf2PasswordEncoder.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPbkdf2PasswordEncoder_Matches(t *testing.T) {
	type args struct {
		encodedPassword string
		rawPassword     string
	}
	tests := []struct {
		name    string
		encoder *Pbkdf2PasswordEncoder
		args    args
		want    *bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.encoder.Matches(tt.args.encodedPassword, tt.args.rawPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pbkdf2PasswordEncoder.Matches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pbkdf2PasswordEncoder.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPbkdf2PasswordEncoder_UpgradeEncoding(t *testing.T) {
	type args struct {
		encodedPassword string
	}
	tests := []struct {
		name    string
		encoder *Pbkdf2PasswordEncoder
		args    args
		want    *bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.encoder.UpgradeEncoding(tt.args.encodedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pbkdf2PasswordEncoder.UpgradeEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pbkdf2PasswordEncoder.UpgradeEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
