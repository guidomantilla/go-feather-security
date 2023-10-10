package security

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"
	feather_web_rest "github.com/guidomantilla/go-feather-web/pkg/rest"
)

type DefaultAuthorizationFilter struct {
	authorizationService AuthorizationService
}

func NewDefaultAuthorizationFilter(authorizationService AuthorizationService) *DefaultAuthorizationFilter {

	if authorizationService == nil {
		feather_commons_log.Fatal("starting up - error setting up authorizationFilter: authorizationService is nil")
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
	token := splits[1]

	application, exists := GetApplicationFromContext(ctx)
	if !exists {
		ex := feather_web_rest.NotFoundException("application name not found in context")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}
	resource := []string{application, ctx.Request.Method, ctx.FullPath()}

	var err error
	var principal *Principal
	ctxWithResource := context.WithValue(ctx.Request.Context(), ResourceCtxKey{}, strings.Join(resource, " "))
	if principal, err = filter.authorizationService.Authorize(ctxWithResource, token); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	AddPrincipalToContext(ctx, principal)
	ctx.Next()
}
