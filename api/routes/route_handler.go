package routes

import (
	taskHandler "backend-golang/modules/task/api/handler"
	userHandler "backend-golang/modules/user/api/handler"
	"backend-golang/pkgs/transport/http/route"
)

type RouteHandler struct {
	UserHandler userHandler.UserHandler
	TaskHandler taskHandler.TaskHandler
}

func (r *RouteHandler) InitGroupRoutes() []route.GroupRoute {
	var routeGroup []route.GroupRoute
	routeGroup = append(routeGroup, r.userRoute())
	routeGroup = append(routeGroup, r.taskRoute())

	return routeGroup
}
