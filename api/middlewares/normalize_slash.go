package middlewares

import "github.com/gin-gonic/gin"

func NormalizeSlashMiddleware(c *gin.Context) {
	// Xóa dấu `/` ở cuối URL nhưng giữ lại cho root "/"
	if c.Request.URL.Path != "/" && len(c.Request.URL.Path) > 1 && c.Request.URL.Path[len(c.Request.URL.Path)-1] == '/' {
		c.Request.URL.Path = c.Request.URL.Path[:len(c.Request.URL.Path)-1]
	}
	c.Next()
}
