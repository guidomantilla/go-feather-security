package auth

import (
	"encoding/json"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	feather_commons_util "github.com/guidomantilla/go-feather-commons/pkg/util"
)

type DefaultClaims struct {
	jwt.RegisteredClaims
	Principal
}

type JwtTokenManager struct {
	issuer        string
	timeout       time.Duration
	secretKey     any
	signingMethod jwt.SigningMethod
}

func NewJwtTokenManager(issuer string, timeout time.Duration, secretKey any, signingMethod jwt.SigningMethod) *JwtTokenManager {
	return &JwtTokenManager{
		issuer:        issuer,
		timeout:       timeout,
		secretKey:     secretKey,
		signingMethod: signingMethod,
	}
}

func (manager *JwtTokenManager) Generate(principal *Principal) (*string, error) {

	claims := &DefaultClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    manager.issuer,
			Subject:   *principal.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(manager.timeout)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Principal: *principal,
	}

	token := jwt.NewWithClaims(manager.signingMethod, claims)

	var err error
	var tokenString string
	if tokenString, err = token.SignedString(manager.secretKey); err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (manager *JwtTokenManager) Validate(tokenString string) (*Principal, error) {

	getKeyFunc := func(token *jwt.Token) (any, error) {
		return manager.secretKey, nil
	}

	parserOptions := []jwt.ParserOption{
		jwt.WithIssuer(manager.issuer),
		jwt.WithValidMethods([]string{manager.signingMethod.Alg()}),
	}

	var err error
	var token *jwt.Token
	if token, err = jwt.Parse(tokenString, getKeyFunc, parserOptions...); err != nil {
		return nil, ErrTokenInvalid
	}

	if !token.Valid {
		return nil, ErrTokenInvalid
	}

	var ok bool
	var mapClaims jwt.MapClaims
	if mapClaims, ok = token.Claims.(jwt.MapClaims); !ok {
		return nil, ErrTokenInvalid
	}

	var value any
	if value, ok = mapClaims["username"]; !ok {
		return nil, ErrTokenInvalid
	}

	var username string
	if username, ok = value.(string); !ok {
		return nil, ErrTokenInvalid
	}

	if value, ok = mapClaims["authorities"]; !ok {
		return nil, ErrTokenInvalid
	}

	var authoritiesBytes []byte
	if authoritiesBytes, err = json.Marshal(value); err != nil {
		return nil, ErrTokenInvalid
	}

	var grantedAuthorities []GrantedAuthority
	if err = json.Unmarshal(authoritiesBytes, &grantedAuthorities); err != nil {
		return nil, ErrTokenInvalid
	}

	principal := &Principal{
		Username:    feather_commons_util.ValueToPtr(username),
		Authorities: feather_commons_util.ValueToPtr(grantedAuthorities),
	}

	return principal, nil
}
