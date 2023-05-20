package auth

import (
	"reflect"
	"testing"
)

func TestPrincipal_Validate(t *testing.T) {
	tests := []struct {
		name      string
		principal *Principal
		want      []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.principal.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Principal.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
