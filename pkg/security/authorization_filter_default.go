package security

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	feather_web_rest "github.com/guidomantilla/go-feather-web/pkg/rest"
)

type DefaultAuthorizationFilter struct {
	authorizationService AuthorizationService
}

func NewDefaultAuthorizationFilter(authorizationService AuthorizationService) *DefaultAuthorizationFilter {

	if authorizationService == nil {
		slog.Error("starting up - error setting up authorizationFilter: authorizationService is nil")
		os.Exit(1)
	}

	return &DefaultAuthorizationFilter{
		authorizationService: authorizationService,
	}
}

func (filter *DefaultAuthorizationFilter) Authorize(ctx *gin.Context) {

	header := ctx.Request.Header.Get("Authorization")
	if !strings.HasPrefix(header, "Bearer ") {
		ex := feather_web_rest.UnauthorizedException("invalid authorization header")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	splits := strings.Split(header, " ")
	if len(splits) != 2 {
		ex := feather_web_rest.UnauthorizedException("invalid authorization header")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	var err error
	var principal *Principal
	ctxWithResource := context.WithValue(ctx.Request.Context(), ResourceCtxKey{}, strings.Join([]string{ctx.Request.Method, ctx.Request.RequestURI}, " "))
	if principal, err = filter.authorizationService.Authorize(ctxWithResource, splits[1]); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	ctx.Set("principal", principal)
	ctx.Next()
}
