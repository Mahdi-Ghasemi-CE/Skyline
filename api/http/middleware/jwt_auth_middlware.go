package middleware

import (
	"Skyline/internal/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		trimedTokenString := strings.TrimPrefix(tokenString, "Bearer ")
		err := utils.ValidateToken(trimedTokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
