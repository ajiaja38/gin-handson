package middleware

import (
	"net/http"
	"os"
	"res-gin/src/model"
	"res-gin/src/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

var access_token_secret []byte = []byte(os.Getenv("SECRET_ACCESS_TOKEN"))

func JWTAuthGuard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  false,
				"message": "Unauthorized Access",
			})

			ctx.Abort()
			return
		}

		token := strings.Split(authHeader, "Bearer ")[1]

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  false,
				"message": "Unauthorized Access",
			})

			ctx.Abort()
			return
		}

		claims := &model.UserClaims{}
		err := utils.ValidateToken(token, access_token_secret, claims)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"status":  false,
				"message": err.Error(),
			})
		}

		ctx.Set("user", claims)
		ctx.Next()
	}
}
