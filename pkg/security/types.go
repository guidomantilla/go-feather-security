package security

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type RoleCtxKey struct{}

type ResourceCtxKey struct{}

type GrantedAuthority struct {
	Role      *string  `json:"role,omitempty"`
	Resources []string `json:"resource,omitempty"`
}

type Principal struct {
	Username           *string            `json:"username,omitempty" binding:"required"`
	Password           *string            `json:"password,omitempty" binding:"required"`
	AccountNonExpired  *bool              `json:"account_non_expired,omitempty"`
	AccountNonLocked   *bool              `json:"account_non_locked,omitempty"`
	PasswordNonExpired *bool              `json:"password_non_expired,omitempty"`
	Enabled            *bool              `json:"enabled,omitempty"`
	SignUpDone         *bool              `json:"signup_done,omitempty"`
	Authorities        []GrantedAuthority `json:"authorities,omitempty"`
}

type PrincipalManager interface {
	Create(ctx context.Context, principal *Principal) error
	Update(ctx context.Context, principal *Principal) error
	Delete(ctx context.Context, username string) error
	Find(ctx context.Context, username string) (*Principal, error)
	Exists(ctx context.Context, username string) error
	ChangePassword(ctx context.Context, username string, password string) error
	VerifyRole(ctx context.Context, username string, role string) error
	VerifyResource(ctx context.Context, username string, resource string) error
}

//

type AuthenticationEndpoint interface {
	Authenticate(ctx *gin.Context)
}

type AuthenticationService interface {
	Authenticate(ctx context.Context, principal *Principal) (*string, error)
}

type AuthenticationDelegate interface {
	Authenticate(ctx context.Context, principal *Principal) error
}

//

type AuthorizationFilter interface {
	Authorize(ctx *gin.Context)
}

type AuthorizationService interface {
	Authorize(ctx context.Context, tokenString string) (*Principal, error)
}

type AuthorizationDelegate interface {
	Authorize(ctx context.Context, principal *Principal) error
}

//

type TokenManager interface {
	Generate(principal *Principal) (*string, error)
	Validate(tokenString string) (*Principal, error)
}

//

const (
	Argon2PrefixKey = "{argon2}"
	BcryptPrefixKey = "{bcrypt}"
	Pbkdf2PrefixKey = "{pbkdf2}"
	ScryptPrefixKey = "{scrypt}"
)

var (
	_ PasswordEncoder   = (*Argon2PasswordEncoder)(nil)
	_ PasswordEncoder   = (*BcryptPasswordEncoder)(nil)
	_ PasswordEncoder   = (*Pbkdf2PasswordEncoder)(nil)
	_ PasswordEncoder   = (*ScryptPasswordEncoder)(nil)
	_ PasswordEncoder   = (*DelegatingPasswordEncoder)(nil)
	_ PasswordEncoder   = (*DefaultPasswordManager)(nil)
	_ PasswordGenerator = (*DefaultPasswordGenerator)(nil)
	_ PasswordGenerator = (*DefaultPasswordManager)(nil)
	_ PasswordManager   = (*DefaultPasswordManager)(nil)
)

type PasswordEncoder interface {
	Encode(rawPassword string) (*string, error)
	Matches(encodedPassword string, rawPassword string) (*bool, error)
	UpgradeEncoding(encodedPassword string) (*bool, error)
}

type PasswordGenerator interface {
	Generate() string
	Validate(rawPassword string) error
}

type PasswordManager interface {
	PasswordEncoder
	PasswordGenerator
}
