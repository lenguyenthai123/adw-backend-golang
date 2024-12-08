package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID ref >> https://minghsu.io/posts/http-context/
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(
			context.WithValue(c.Request.Context(),
				"X-Request-ID", uuid.New().String()))

		c.Next()
	}
}
