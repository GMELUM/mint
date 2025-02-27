package middleware

import (
	"mint/config"
	"mint/utils/msg"

	"github.com/gin-gonic/gin"
)

func Secret(ctx *gin.Context) {
	// Try to get the token from the "Authorization" header first
	token := ctx.GetHeader("Authorization")

	// If the header is not set or empty, use the query parameter
	if len(token) == 0 {
		token = ctx.Query("secret") // Use the actual query parameter name if different
	}

	if token != config.Secret {
		msg.Forbidden(ctx)
		ctx.Abort()
		return
	}

	ctx.Next()
}
