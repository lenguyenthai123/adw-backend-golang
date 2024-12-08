package routes

import (
	userHandler "backend-golang/modules/user/api/handler"
	"backend-golang/pkgs/transport/http/route"
)

type RouteHandler struct {
	UserHandler userHandler.UserHandler
}

func (r *RouteHandler) InitGroupRoutes() []route.GroupRoute {
	var routeGroup []route.GroupRoute
	routeGroup = append(routeGroup, r.userRoute())

	return routeGroup
}
