package routes

import (
	analyticsHandler "backend-golang/modules/analytics/api/handler"
	taskHandler "backend-golang/modules/task/api/handler"
	userHandler "backend-golang/modules/user/api/handler"
	"backend-golang/pkgs/transport/http/route"
)

type RouteHandler struct {
	UserHandler      userHandler.UserHandler
	TaskHandler      taskHandler.TaskHandler
	AnalyticsHandler analyticsHandler.AnalyticsHandler
}

func (r *RouteHandler) InitGroupRoutes() []route.GroupRoute {
	var routeGroup []route.GroupRoute
	routeGroup = append(routeGroup, r.userRoute())
	routeGroup = append(routeGroup, r.taskRoute())
	routeGroup = append(routeGroup, r.analyticsRoute())

	return routeGroup
}
