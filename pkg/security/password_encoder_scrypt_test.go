package security

import (
	"reflect"
	"testing"
)

func TestNewScryptPasswordEncoder(t *testing.T) {
	type args struct {
		options []ScryptPasswordEncoderOption
	}
	tests := []struct {
		name string
		args args
		want *ScryptPasswordEncoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScryptPasswordEncoder(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScryptPasswordEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithScryptN(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want ScryptPasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithScryptN(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithScryptN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithScryptR(t *testing.T) {
	type args struct {
		r int
	}
	tests := []struct {
		name string
		args args
		want ScryptPasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithScryptR(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithScryptR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithScryptP(t *testing.T) {
	type args struct {
		p int
	}
	tests := []struct {
		name string
		args args
		want ScryptPasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithScryptP(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithScryptP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithScryptSaltLength(t *testing.T) {
	type args struct {
		saltLength int
	}
	tests := []struct {
		name string
		args args
		want ScryptPasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithScryptSaltLength(tt.args.saltLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithScryptSaltLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithScryptKeyLength(t *testing.T) {
	type args struct {
		keyLength int
	}
	tests := []struct {
		name string
		args args
		want ScryptPasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithScryptKeyLength(tt.args.keyLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithScryptKeyLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScryptPasswordEncoder_Encode(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name    string
		encoder *ScryptPasswordEncoder
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
				t.Errorf("ScryptPasswordEncoder.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScryptPasswordEncoder.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScryptPasswordEncoder_Matches(t *testing.T) {
	type args struct {
		encodedPassword string
		rawPassword     string
	}
	tests := []struct {
		name    string
		encoder *ScryptPasswordEncoder
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
				t.Errorf("ScryptPasswordEncoder.Matches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScryptPasswordEncoder.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScryptPasswordEncoder_UpgradeEncoding(t *testing.T) {
	type args struct {
		encodedPassword string
	}
	tests := []struct {
		name    string
		encoder *ScryptPasswordEncoder
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
				t.Errorf("ScryptPasswordEncoder.UpgradeEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScryptPasswordEncoder.UpgradeEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
