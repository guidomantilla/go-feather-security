package security

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guidomantilla/go-feather-web/pkg/rest"
	"go.uber.org/zap"

	"github.com/guidomantilla/go-feather-security/pkg/validation"
)

type DefaultAuthenticationEndpoint struct {
	authenticationService AuthenticationService
}

func NewDefaultAuthenticationEndpoint(authenticationService AuthenticationService) *DefaultAuthenticationEndpoint {

	if authenticationService == nil {
		zap.L().Fatal("starting up - error setting up authenticationEndpoint: authenticationService is nil")
	}

	return &DefaultAuthenticationEndpoint{
		authenticationService: authenticationService,
	}
}

func (endpoint *DefaultAuthenticationEndpoint) Authenticate(ctx *gin.Context) {

	var err error
	var principal *Principal
	if err = ctx.ShouldBindJSON(&principal); err != nil {
		ex := rest.BadRequestException("error unmarshalling request json to object")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if errs := validate(principal); errs != nil {
		ex := rest.BadRequestException("error validating the principal", errs...)
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	var token *string
	if token, err = endpoint.authenticationService.Authenticate(ctx.Request.Context(), principal); err != nil {
		ex := rest.UnauthorizedException(err.Error())
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func validate(principal *Principal) []error {

	var errors []error

	if err := validation.ValidateFieldIsRequired("this", "username", principal.Username); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldIsRequired("this", "password", principal.Password); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "account_non_expired", principal.AccountNonExpired); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "account_non_locked", principal.AccountNonLocked); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "password_non_expired", principal.PasswordNonExpired); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "enabled", principal.Enabled); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateFieldMustBeUndefined("this", "signup_done", principal.SignUpDone); err != nil {
		errors = append(errors, err)
	}

	if err := validation.ValidateStructMustBeUndefined("this", "authorities", principal.Authorities); err != nil {
		errors = append(errors, err)
		return errors
	}

	return errors
}
