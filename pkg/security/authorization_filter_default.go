package security

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/guidomantilla/go-feather-web/pkg/rest"
	"go.uber.org/zap"
)

type DefaultAuthorizationFilter struct {
	authorizationService AuthorizationService
}

func NewDefaultAuthorizationFilter(authorizationService AuthorizationService) *DefaultAuthorizationFilter {

	if authorizationService == nil {
		zap.L().Fatal("starting up - error setting up authorizationFilter: authorizationService is nil")
	}

	return &DefaultAuthorizationFilter{
		authorizationService: authorizationService,
	}
}

func (filter *DefaultAuthorizationFilter) Authorize(ctx *gin.Context) {

	header := ctx.Request.Header.Get("Authorization")
	if !strings.HasPrefix(header, "Bearer ") {
		ex := rest.UnauthorizedException("invalid authorization header")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	splits := strings.Split(header, " ")
	if len(splits) != 2 {
		ex := rest.UnauthorizedException("invalid authorization header")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	var err error
	var principal *Principal
	ctxWithResource := context.WithValue(ctx.Request.Context(), ResourceCtxKey{}, ctx.Request.URL.Path)
	if principal, err = filter.authorizationService.Authorize(ctxWithResource, splits[1]); err != nil {
		ex := rest.UnauthorizedException("invalid authorization header", err)
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	ctx.Set("principal", principal)
	ctx.Next()
}
