package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinOption func(*gin.Engine)

// NewGin creates a new Gin engine with the given options.
//
// It accepts a variable number of GinOption parameters, which can be used to configure the engine.
// The function returns a pointer to the created gin.Engine.
func NewGin(options ...GinOption) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	for _, option := range options {
		option(r)
	}

	return r
}

// AddMiddlewares returns a GinOption that adds the given middlewares to a Gin Engine.
//
// ms is a variadic parameter that accepts multiple functions that take a *gin.Context as a parameter.
// The functions passed to ms will be used as middlewares in the Gin Engine.
// The function returns a GinOption, which is a function that takes a *gin.Engine as a parameter.
// When the GinOption is called, it adds each of the middlewares in ms to the Gin Engine using the Use() method.
func AddMiddlewares(ms ...func(c *gin.Context)) GinOption {
	return func(g *gin.Engine) {
		for _, m := range ms {
			g.Use(m)
		}
	}
}

// AddGroupRoutes adds a group of routes to a Gin Engine.
//
// gr: A slice of GroupRoute structs representing the routes to be added.
// Returns a GinOption function that adds the specified routes to a Gin Engine.
func AddGroupRoutes(gr []GroupRoute) GinOption {
	return func(g *gin.Engine) {
		for _, r := range gr {
			r.AddGroupRoute(g)
		}
	}
}

// AddRoutes returns a GinOption function that adds multiple routes to a gin.Engine.
//
// It takes a slice of Route objects as a parameter.
// The function iterates over each Route object in the slice and calls the AddRoute method on it,
// passing the gin.Engine as an argument.
// The function is used to configure the gin.Engine with the provided routes.
// The GinOption function is then returned.
func AddRoutes(rs []Route) GinOption {
	return func(g *gin.Engine) {
		for _, r := range rs {
			r.AddRoute(g)
		}
	}
}

// StrictSlash sets the RemoveExtraSlash field of the gin.Engine struct.
//
// The function takes a boolean parameter, `strict`, which determines whether
// extra slashes should be removed from the URL path. If `strict` is set to
// `true`, extra slashes will be removed. If `strict` is set to `false`, extra
// slashes will not be removed.
//
// The function returns a GinOption, which is a function that takes a pointer to
// a gin.Engine struct and modifies its RemoveExtraSlash field based on the
// value of `strict`.
func StrictSlash(strict bool) GinOption {
	return func(g *gin.Engine) {
		g.RemoveExtraSlash = strict
	}
}

// SetMaximumMultipartSize sets the maximum size of a multipart request in bytes.
//
// Parameters:
// - size: the maximum size in bytes.
//
// Returns:
// - GinOption: a function that sets the maximum multipart size.
func SetMaximumMultipartSize(size int64) GinOption {
	return func(g *gin.Engine) {
		g.MaxMultipartMemory = size
	}
}

// AddGinOptions generates a GinOption that applies a series of GinOptions to a *gin.Engine.
//
// options: A variadic parameter of GinOptions to be applied.
// returns: A GinOption that applies the given options to a *gin.Engine.
func AddGinOptions(options ...GinOption) GinOption {
	return func(e *gin.Engine) {
		for _, o := range options {
			o(e)
		}
	}
}

// AddHealthCheckRoute generates a GinOption function that adds a health check route to the Gin engine.
//
// This function takes no parameters.
// It returns a GinOption function.
func AddHealthCheckRoute() GinOption {
	return func(g *gin.Engine) {
		g.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})
	}
}

// AddNoRouteHandler returns a GinOption function that sets the handler for when no route is found.
//
// The returned GinOption function takes a *gin.Engine as a parameter and sets the handler for the NoRoute
// method of the gin.Engine. The handler is a function that takes a *gin.Context as a parameter and
// returns no values. It sets the response status code to http.StatusNotFound and returns a JSON response
// with status_code, message, and key fields.
//
// It returns the GinOption function that sets the NoRoute handler.
func AddNoRouteHandler() GinOption {
	return func(g *gin.Engine) {
		g.NoRoute(func(g *gin.Context) {
			g.JSON(http.StatusNotFound, gin.H{
				"status_code": http.StatusNotFound,
				"message":     "api not found",
				"key":         "ERR_API_NOT_FOUND",
			})
		})
	}
}
