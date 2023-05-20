package validation

import "testing"

func TestValidateStructMustBeUndefined(t *testing.T) {
	type args struct {
		structName string
		field      string
		value      any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateStructMustBeUndefined(tt.args.structName, tt.args.field, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateStructMustBeUndefined() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateStructIsRequired(t *testing.T) {
	type args struct {
		structName string
		field      string
		value      any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateStructIsRequired(tt.args.structName, tt.args.field, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateStructIsRequired() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateFieldMustBeUndefined(t *testing.T) {
	type args struct {
		structName string
		field      string
		value      *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateFieldMustBeUndefined(tt.args.structName, tt.args.field, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateFieldMustBeUndefined() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateFieldIsRequired(t *testing.T) {
	type args struct {
		structName string
		field      string
		value      *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateFieldIsRequired(tt.args.structName, tt.args.field, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateFieldIsRequired() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
