package middlewares

import (
	"backend-golang/core"
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

		// Chuyển `token.Claims` sang kiểu `jwt.MapClaims`
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Lấy giá trị từ claim `userId`
		userId, exists := claims["userId"].(string)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "userId not found in token claims"})
			c.Abort()
			return
		}

		//Set claim
		c.Set("claims", claims)

		// Set requester information in context
		c.Set(core.CurrentRequesterKeyString, core.RestRequester{
			ID:   userId,
			Role: "",
		})
		c.Next()

	}
}

// UserVipMiddleware checks if the user has the user-vip role
func UserVipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token claims"})
			c.Abort()
			return
		}

		jwtClaims, ok := claims.(jwt.MapClaims)
		if !ok || !isUserVip(jwtClaims) {
			c.JSON(http.StatusForbidden, gin.H{"error": "User is not vip"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// isUserVip kiểm tra xem token có phải là vip không
func isUserVip(claims jwt.MapClaims) bool {
	realmAccess, exists := claims["realm_access"].(map[string]interface{})
	if !exists {
		return false
	}

	roles, exists := realmAccess["roles"].([]interface{})
	if !exists {
		return false
	}

	for _, role := range roles {
		if role == "user-vip" {
			return true
		}
	}

	return false
}
