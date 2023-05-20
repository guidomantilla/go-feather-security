package password

import (
	"reflect"
	"testing"
)

func TestNewArgon2PasswordEncoder(t *testing.T) {
	type args struct {
		options []Argon2PasswordEncoderOption
	}
	tests := []struct {
		name string
		args args
		want *Argon2PasswordEncoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArgon2PasswordEncoder(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArgon2PasswordEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithArgon2Iterations(t *testing.T) {
	type args struct {
		iterations int
	}
	tests := []struct {
		name string
		args args
		want Argon2PasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithArgon2Iterations(tt.args.iterations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithArgon2Iterations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithArgon2Memory(t *testing.T) {
	type args struct {
		memory int
	}
	tests := []struct {
		name string
		args args
		want Argon2PasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithArgon2Memory(tt.args.memory); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithArgon2Memory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithArgon2Threads(t *testing.T) {
	type args struct {
		threads int
	}
	tests := []struct {
		name string
		args args
		want Argon2PasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithArgon2Threads(tt.args.threads); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithArgon2Threads() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithArgon2SaltLength(t *testing.T) {
	type args struct {
		saltLength int
	}
	tests := []struct {
		name string
		args args
		want Argon2PasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithArgon2SaltLength(tt.args.saltLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithArgon2SaltLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithArgon2KeyLength(t *testing.T) {
	type args struct {
		keyLength int
	}
	tests := []struct {
		name string
		args args
		want Argon2PasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithArgon2KeyLength(tt.args.keyLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithArgon2KeyLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgon2PasswordEncoder_Encode(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name    string
		encoder *Argon2PasswordEncoder
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
				t.Errorf("Argon2PasswordEncoder.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Argon2PasswordEncoder.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgon2PasswordEncoder_Matches(t *testing.T) {
	type args struct {
		encodedPassword string
		rawPassword     string
	}
	tests := []struct {
		name    string
		encoder *Argon2PasswordEncoder
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
				t.Errorf("Argon2PasswordEncoder.Matches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Argon2PasswordEncoder.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgon2PasswordEncoder_UpgradeEncoding(t *testing.T) {
	type args struct {
		encodedPassword string
	}
	tests := []struct {
		name    string
		encoder *Argon2PasswordEncoder
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
				t.Errorf("Argon2PasswordEncoder.UpgradeEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Argon2PasswordEncoder.UpgradeEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
