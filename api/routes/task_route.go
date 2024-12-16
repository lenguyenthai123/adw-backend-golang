package routes

import (
	"backend-golang/api/middlewares"
	taskHandler "backend-golang/modules/task/api/handler"
	"backend-golang/modules/task/domain/usecase"
	"backend-golang/modules/task/repository"
	database "backend-golang/pkgs/dbs/postgres"
	"backend-golang/pkgs/transport/http/method"
	"backend-golang/pkgs/transport/http/route"
	"github.com/go-playground/validator/v10"
	"github.com/openai/openai-go"
)

func NewTaskHandler(db *database.Database, openaiClient *openai.Client, requestValidator *validator.Validate) taskHandler.TaskHandler {
	// Task Repository
	taskRepoReader := repository.NewTaskReaderRepository(*db)
	taskRepoWriter := repository.NewTaskWriterRepository(*db)

	return taskHandler.NewTaskHandler(requestValidator,
		usecase.NewCreateTaskUsecase(taskRepoWriter),
		usecase.NewGetTaskUsecase(taskRepoReader),
		usecase.NewUpdateTaskUsecase(taskRepoReader, taskRepoWriter),
		usecase.NewDeleteTaskUsecase(taskRepoWriter),
		usecase.NewGetTaskListUsecase(taskRepoReader),
		usecase.NewAnalyzeTaskUsecase(taskRepoReader, openaiClient),
		usecase.NewApplyAnalyzedTaskUsecase(taskRepoWriter))
}

func (r *RouteHandler) taskRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/task",
		Routes: []route.Route{
			{
				Path:    "/",
				Method:  method.POST,
				Handler: r.TaskHandler.HandleCreateTask,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/:task_id",
				Method:  method.GET,
				Handler: r.TaskHandler.HandleGetTask,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/:task_id",
				Method:  method.PATCH,
				Handler: r.TaskHandler.HandleUpdateTask,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/:task_id",
				Method:  method.DELETE,
				Handler: r.TaskHandler.HandleDeleteTask,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/task_list",
				Method:  method.GET,
				Handler: r.TaskHandler.HandleGetTaskList,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/analyze",
				Method:  method.POST,
				Handler: r.TaskHandler.HandleAnalyzeTask,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/analyze/apply",
				Method:  method.POST,
				Handler: r.TaskHandler.HandleApplyAnalyzedTask,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
		},
	}
}
