package auth

import "errors"

var (
	ErrFailedAuthentication = errors.New("incorrect principal username or password")
	ErrFailedAuthorization  = errors.New("invalid authorities")
	ErrTokenInvalid         = errors.New("invalid token")
)
