package handler

import (
	"backend-golang/modules/task/domain/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TaskHandler interface {
	HandleCreateTask(c *gin.Context)
	HandleGetTask(c *gin.Context)
	HandleUpdateTask(c *gin.Context)
	HandleDeleteTask(c *gin.Context)
}

type TaskHandlerImpl struct {
	requestValidator  *validator.Validate
	createTaskUsecase usecase.CreateTaskUsecase
	getTaskUsecase    usecase.GetTaskUsecase
	updateTaskUsecase usecase.UpdateTaskUsecase
	deleteTaskUsecase usecase.DeleteTaskUsecase
}

func NewTaskHandler(
	createTaskUsecase usecase.CreateTaskUsecase,
	getTaskUsecase usecase.GetTaskUsecase,
	updateTaskUsecase usecase.UpdateTaskUsecase,
	deleteTaskUsecase usecase.DeleteTaskUsecase,
) TaskHandler {
	return &TaskHandlerImpl{
		createTaskUsecase: createTaskUsecase,
		getTaskUsecase:    getTaskUsecase,
		updateTaskUsecase: updateTaskUsecase,
		deleteTaskUsecase: deleteTaskUsecase,
	}
}
