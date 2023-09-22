package security

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	feather_web_rest "github.com/guidomantilla/go-feather-web/pkg/rest"
)

type DefaultAuthenticationEndpoint struct {
	authenticationService AuthenticationService
}

func NewDefaultAuthenticationEndpoint(authenticationService AuthenticationService) *DefaultAuthenticationEndpoint {

	if authenticationService == nil {
		slog.Error("starting up - error setting up authenticationEndpoint: authenticationService is nil")
		os.Exit(1)
	}

	return &DefaultAuthenticationEndpoint{
		authenticationService: authenticationService,
	}
}

func (endpoint *DefaultAuthenticationEndpoint) Authenticate(ctx *gin.Context) {

	var err error
	var principal *Principal
	if err = ctx.ShouldBindJSON(&principal); err != nil {
		ex := feather_web_rest.BadRequestException("error unmarshalling request json to object")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if errs := endpoint.authenticationService.Validate(principal); errs != nil {
		ex := feather_web_rest.BadRequestException("error validating the principal", errs...)
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if err = endpoint.authenticationService.Authenticate(ctx.Request.Context(), principal); err != nil {
		ex := feather_web_rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	ctx.JSON(http.StatusOK, principal)
}
