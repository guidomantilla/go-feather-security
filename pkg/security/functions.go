package security

import (
	"github.com/gin-gonic/gin"
)

const (
	PrincipalCtxKey = "principal"
)

func AddPrincipalToContext(ctx *gin.Context, principal *Principal) {
	ctx.Set(PrincipalCtxKey, principal)
}

func GetPrincipalFROMContext(ctx *gin.Context) (*Principal, bool) {

	var exists bool
	var principal any
	if principal, exists = ctx.Get(PrincipalCtxKey); !exists {
		return nil, false
	}
	return principal.(*Principal), true
}
