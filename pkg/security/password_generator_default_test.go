package security

import (
	"reflect"
	"testing"
)

func TestNewDefaultPasswordGenerator(t *testing.T) {
	type args struct {
		options []DefaultPasswordGeneratorOption
	}
	tests := []struct {
		name string
		args args
		want *DefaultPasswordGenerator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultPasswordGenerator(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultPasswordGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPasswordLength(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want DefaultPasswordGeneratorOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPasswordLength(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPasswordLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMinSpecialChar(t *testing.T) {
	type args struct {
		minSpecialChar int
	}
	tests := []struct {
		name string
		args args
		want DefaultPasswordGeneratorOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMinSpecialChar(tt.args.minSpecialChar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMinSpecialChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMinNum(t *testing.T) {
	type args struct {
		minNum int
	}
	tests := []struct {
		name string
		args args
		want DefaultPasswordGeneratorOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMinNum(tt.args.minNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMinNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMinUpperCase(t *testing.T) {
	type args struct {
		minUpperCase int
	}
	tests := []struct {
		name string
		args args
		want DefaultPasswordGeneratorOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMinUpperCase(tt.args.minUpperCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMinUpperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultPasswordGenerator_Generate(t *testing.T) {
	tests := []struct {
		name      string
		generator *DefaultPasswordGenerator
		want      string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.generator.Generate(); got != tt.want {
				t.Errorf("DefaultPasswordGenerator.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultPasswordGenerator_Validate(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name      string
		generator *DefaultPasswordGenerator
		args      args
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.generator.Validate(tt.args.rawPassword); (err != nil) != tt.wantErr {
				t.Errorf("DefaultPasswordGenerator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
