package tools

//go:generate mockgen -package security -destination ../pkg/security/mocks.go github.com/guidomantilla/go-feather-security/pkg/security AuthenticationEndpoint,AuthenticationService,AuthenticationDelegate,AuthorizationFilter,AuthorizationService,AuthorizationDelegate,PrincipalManager,TokenManager,PasswordEncoder,PasswordGenerator,PasswordManager
