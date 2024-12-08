package route

import (
	"backend-golang/pkgs/transport/http/method"
	"github.com/gin-gonic/gin"
)

type GroupRoute struct {
	Prefix      string
	Middlewares []func(*gin.Context)
	Routes      []Route
}

type Route struct {
	Path        string
	Method      method.Method
	Handler     func(*gin.Context)
	Middlewares []func(*gin.Context)
}

func (r Route) CombineHandler() []gin.HandlerFunc {
	var h []gin.HandlerFunc
	for _, m := range r.Middlewares {
		h = append(h, m)
	}
	h = append(h, r.Handler)
	return h
}

// AddRoute adds a route to the gin.Engine based on the Route's Method and Path.
//
// It takes a pointer to a gin.Engine as the parameter.
// It does not return anything.
func (r Route) AddRoute(g *gin.Engine) {
	switch r.Method {
	case method.GET:
		g.GET(r.Path, r.CombineHandler()...)
	case method.POST:
		g.POST(r.Path, r.CombineHandler()...)
	case method.PUT:
		g.PUT(r.Path, r.CombineHandler()...)
	case method.PATCH:
		g.PATCH(r.Path, r.CombineHandler()...)
	case method.DELETE:
		g.DELETE(r.Path, r.CombineHandler()...)
	case method.HEAD:
		g.HEAD(r.Path, r.CombineHandler()...)
	case method.OPTIONS:
		g.OPTIONS(r.Path, r.CombineHandler()...)
	}
}

// AddGroupRoute adds a group route to the gin.Engine instance.
//
// Parameters:
// - g: a pointer to a gin.Engine instance.
//
// Return type: none
func (r GroupRoute) AddGroupRoute(g *gin.Engine) {
	gr := g.Group(r.Prefix)
	for _, m := range r.Middlewares {
		gr.Use(m)
	}
	for _, r := range r.Routes {
		switch r.Method {
		case method.GET:
			gr.GET(r.Path, r.CombineHandler()...)
		case method.POST:
			gr.POST(r.Path, r.CombineHandler()...)
		case method.PUT:
			gr.PUT(r.Path, r.CombineHandler()...)
		case method.PATCH:
			gr.PATCH(r.Path, r.CombineHandler()...)
		case method.DELETE:
			gr.DELETE(r.Path, r.CombineHandler()...)
		case method.HEAD:
			gr.HEAD(r.Path, r.CombineHandler()...)
		case method.OPTIONS:
			gr.OPTIONS(r.Path, r.CombineHandler()...)
		}
	}
}

func Middlewares(ms ...func(*gin.Context)) []func(*gin.Context) {
	return ms
}
