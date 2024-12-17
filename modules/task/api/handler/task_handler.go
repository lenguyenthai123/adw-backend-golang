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
	HandleDeleteTaskList(c *gin.Context)
	HandleGetTaskList(c *gin.Context)
	HandleAnalyzeTask(c *gin.Context)
	HandleApplyAnalyzedTask(c *gin.Context)
}

type TaskHandlerImpl struct {
	requestValidator         *validator.Validate
	createTaskUsecase        usecase.CreateTaskUsecase
	getTaskUsecase           usecase.GetTaskUsecase
	updateTaskUsecase        usecase.UpdateTaskUsecase
	deleteTaskUsecase        usecase.DeleteTaskUsecase
	deleteTaskListUsecase    usecase.DeleteTaskListUsecase
	getTaskListUsecase       usecase.GetTaskListUsecase
	analyzeTaskUsecase       usecase.AnalyzeTaskUsecase
	applyAnalyzedTaskUsecase usecase.ApplyAnalyzedTaskUsecase
}

func NewTaskHandler(
	requestValidator *validator.Validate,
	createTaskUsecase usecase.CreateTaskUsecase,
	getTaskUsecase usecase.GetTaskUsecase,
	updateTaskUsecase usecase.UpdateTaskUsecase,
	deleteTaskUsecase usecase.DeleteTaskUsecase,
	deleteTaskListUsecase usecase.DeleteTaskListUsecase,
	getTaskListUsecase usecase.GetTaskListUsecase,
	analyzeTaskUsecase usecase.AnalyzeTaskUsecase,
	applyAnalyzedTaskUsecase usecase.ApplyAnalyzedTaskUsecase,

) TaskHandler {
	return &TaskHandlerImpl{
		requestValidator:         requestValidator,
		createTaskUsecase:        createTaskUsecase,
		getTaskUsecase:           getTaskUsecase,
		updateTaskUsecase:        updateTaskUsecase,
		deleteTaskUsecase:        deleteTaskUsecase,
		deleteTaskListUsecase:    deleteTaskListUsecase,
		getTaskListUsecase:       getTaskListUsecase,
		analyzeTaskUsecase:       analyzeTaskUsecase,
		applyAnalyzedTaskUsecase: applyAnalyzedTaskUsecase,
	}
}
