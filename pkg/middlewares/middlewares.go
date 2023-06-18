package middlewares

import (
	"net/http"

	"github.com/DaZZler12/MyRestServer/pkg/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := utils.TokenValid(context)
		if err != nil {
			context.String(http.StatusUnauthorized, "Unauthorized")
			context.Abort()
			return
		}
		context.Next()
	}
}
