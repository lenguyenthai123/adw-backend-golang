package middlewares

import (
	"net/http"
	"strings"

	"backend-golang/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// VerifyJWTMiddleware là middleware xác thực JWT
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy Access Token từ header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		// Tách token từ header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Xác thực token
		token, err := jwt.Parse(tokenString, config.JWKS.Keyfunc)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Token hợp lệ, tiếp tục request
		c.Next()
	}
}
