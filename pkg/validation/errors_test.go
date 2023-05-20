package validation

import (
	"reflect"
	"testing"
)

func TestNewErrFieldMustBeUndefined(t *testing.T) {
	type args struct {
		structName string
		field      string
	}
	tests := []struct {
		name string
		args args
		want *ErrFieldMustBeUndefined
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewErrFieldMustBeUndefined(tt.args.structName, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrFieldMustBeUndefined() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrFieldMustBeUndefined_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *ErrFieldMustBeUndefined
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("ErrFieldMustBeUndefined.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewErrFieldIsRequired(t *testing.T) {
	type args struct {
		structName string
		field      string
	}
	tests := []struct {
		name string
		args args
		want *ErrFieldIsRequired
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewErrFieldIsRequired(tt.args.structName, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrFieldIsRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrFieldIsRequired_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *ErrFieldIsRequired
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("ErrFieldIsRequired.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
