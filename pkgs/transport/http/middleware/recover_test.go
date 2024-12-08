package middleware

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestRecover(t *testing.T) {
	// Visit http://localhost:8080/recover-test to test
	app := gin.Default()
	app.Use(Recover())

	app.GET("/recover-test", func(c *gin.Context) {
		panic(ErrInternal(nil))
	})

	err := app.Run(":8080")
	if err != nil {
		t.Errorf("Error running server: %v", err)
	}
}
