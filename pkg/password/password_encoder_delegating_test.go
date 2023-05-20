package password

import (
	"reflect"
	"testing"
)

func TestNewDelegatingPasswordEncoder(t *testing.T) {
	type args struct {
		encoder PasswordEncoder
		options []DelegatingPasswordEncoderOption
	}
	tests := []struct {
		name string
		args args
		want *DelegatingPasswordEncoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDelegatingPasswordEncoder(tt.args.encoder, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDelegatingPasswordEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithSupportedDecoders(t *testing.T) {
	type args struct {
		decoders map[string]PasswordEncoder
	}
	tests := []struct {
		name string
		args args
		want DelegatingPasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithSupportedDecoders(tt.args.decoders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithSupportedDecoders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelegatingPasswordEncoder_Encode(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name    string
		encoder *DelegatingPasswordEncoder
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
				t.Errorf("DelegatingPasswordEncoder.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelegatingPasswordEncoder.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelegatingPasswordEncoder_Matches(t *testing.T) {
	type args struct {
		encodedPassword string
		rawPassword     string
	}
	tests := []struct {
		name    string
		encoder *DelegatingPasswordEncoder
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
				t.Errorf("DelegatingPasswordEncoder.Matches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelegatingPasswordEncoder.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelegatingPasswordEncoder_UpgradeEncoding(t *testing.T) {
	type args struct {
		encodedPassword string
	}
	tests := []struct {
		name    string
		encoder *DelegatingPasswordEncoder
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
				t.Errorf("DelegatingPasswordEncoder.UpgradeEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelegatingPasswordEncoder.UpgradeEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
