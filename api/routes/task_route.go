package routes

import (
	taskHandler "backend-golang/modules/task/api/handler"
	"backend-golang/modules/task/domain/usecase"
	"backend-golang/modules/task/repository"
	database "backend-golang/pkgs/dbs/postgres"
	"backend-golang/pkgs/transport/http/method"
	"backend-golang/pkgs/transport/http/route"

	"github.com/go-playground/validator/v10"
)

func NewTaskHandler(db *database.Database, requestValidator *validator.Validate) taskHandler.TaskHandler {
	// Task Repository
	taskRepoReader := repository.NewTaskReaderRepository(*db)
	taskRepoWriter := repository.NewTaskWriterRepository(*db)

	return taskHandler.NewTaskHandler(requestValidator,
		usecase.NewCreateTaskUsecase(taskRepoWriter),
		usecase.NewGetTaskUsecase(taskRepoReader),
		usecase.NewUpdateTaskUsecase(taskRepoReader, taskRepoWriter),
		usecase.NewDeleteTaskUsecase(taskRepoWriter))
}

func (r *RouteHandler) taskRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/task",
		Routes: []route.Route{
			{
				Path:    "/",
				Method:  method.POST,
				Handler: r.TaskHandler.HandleCreateTask,
			},
			{
				Path:    "/:task_id",
				Method:  method.GET,
				Handler: r.TaskHandler.HandleGetTask,
			},
			{
				Path:    "/:task_id",
				Method:  method.PATCH,
				Handler: r.TaskHandler.HandleUpdateTask,
			},
			{
				Path:    "/:task_id",
				Method:  method.DELETE,
				Handler: r.TaskHandler.HandleDeleteTask,
			},
		},
	}
}
