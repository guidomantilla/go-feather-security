package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type AuthenticationEndpoint interface {
	Authenticate(ctx *gin.Context)
}

type AuthenticationService interface {
	Authenticate(ctx context.Context, principal *Principal) (*string, error)
}

type AuthenticationDelegate interface {
	Authenticate(ctx context.Context, principal *Principal) error
}

type AuthorizationFilter interface {
	Authorize(ctx *gin.Context)
}

type AuthorizationService interface {
	Authorize(ctx context.Context, tokenString string) (*Principal, error)
}

type AuthorizationDelegate interface {
	Authorize(ctx context.Context, principal *Principal) error
}

type TokenManager interface {
	Generate(principal *Principal) (*string, error)
	Validate(tokenString string) (*Principal, error)
}

type PrincipalManager interface {
	Create(ctx context.Context, principal *Principal) error
	Update(ctx context.Context, principal *Principal) error
	Delete(ctx context.Context, username string) error
	Find(ctx context.Context, username string) (*Principal, error)
	Exists(ctx context.Context, username string) error
	ChangePassword(ctx context.Context, principal *Principal) error
}
