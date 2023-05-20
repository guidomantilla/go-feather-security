package security

import "errors"

var (
	ErrPasswordEncoderNotFound   = errors.New("password encoder not found")
	ErrRawPasswordIsEmpty        = errors.New("rawPassword cannot be empty")
	ErrSaltIsNil                 = errors.New("salt cannot be nil")
	ErrSaltIsEmpty               = errors.New("salt cannot be empty")
	ErrHashFuncIsNil             = errors.New("hashFunc cannot be nil")
	ErrEncodedPasswordIsEmpty    = errors.New("encodedPassword cannot be empty")
	ErrEncodedPasswordNotAllowed = errors.New("encodedPassword format not allowed")
	ErrBcryptCostNotAllowed      = errors.New("bcryptCost not allowed")
	ErrFailedAuthentication      = errors.New("incorrect principal username or password")
	ErrFailedAuthorization       = errors.New("invalid authorities")
	ErrPasswordNotStrong         = errors.New("principal password is not strong enough")
)
