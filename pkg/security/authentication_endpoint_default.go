package security

import (
	"net/http"

	"github.com/gin-gonic/gin"
	feather_web_rest "github.com/guidomantilla/go-feather-web/pkg/rest"
	"go.uber.org/zap"

	feather_security_validation "github.com/guidomantilla/go-feather-security/pkg/validation"
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
		ex := feather_web_rest.BadRequestException("error unmarshalling request json to object")
		ctx.AbortWithStatusJSON(ex.Code, ex)
		return
	}

	if errs := validate(principal); errs != nil {
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

func validate(principal *Principal) []error {

	var errors []error

	if err := feather_security_validation.ValidateFieldIsRequired("this", "username", principal.Username); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "role", principal.Role); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldIsRequired("this", "password", principal.Password); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "account_non_expired", principal.AccountNonExpired); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "account_non_locked", principal.AccountNonLocked); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "password_non_expired", principal.PasswordNonExpired); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "enabled", principal.Enabled); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "signup_done", principal.SignUpDone); err != nil {
		errors = append(errors, err)
	}

	if err := feather_security_validation.ValidateStructMustBeUndefined("this", "authorities", principal.Authorities); err != nil {
		errors = append(errors, err)
		return errors
	}

	if err := feather_security_validation.ValidateFieldMustBeUndefined("this", "token", principal.Token); err != nil {
		errors = append(errors, err)
	}

	return errors
}
