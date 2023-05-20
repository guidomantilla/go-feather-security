package security

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewDefaultAuthorizationFilter(t *testing.T) {
	type args struct {
		authorizationService AuthorizationService
	}
	tests := []struct {
		name string
		args args
		want *DefaultAuthorizationFilter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultAuthorizationFilter(tt.args.authorizationService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultAuthorizationFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultAuthorizationFilter_Authorize(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		filter *DefaultAuthorizationFilter
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.filter.Authorize(tt.args.ctx)
		})
	}
}
