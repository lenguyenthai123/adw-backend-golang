package handler

import (
	"backend-golang/modules/task/api/mapper"
	"backend-golang/modules/task/api/model/req"
	response "backend-golang/modules/task/api/model/res"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleCreateTask	godoc
// @Summary      Create new task
// @Description  Create new task
// @Tags         Task
// @Produce      json
// @Param        CreateTaskRequest  body	req.CreateTaskRequest  true  "CreateTaskRequest JSON"
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /task [post]

func (h *TaskHandlerImpl) HandleCreateTask(c *gin.Context) {
	// Bind request
	var createTaskRequest req.CreateTaskRequest
	// if err := c.ShouldBindJSON(&createTaskRequest); err != nil {
	// 	return res.ErrInvalidRequest(err)
	// }

	// // Validate request
	// if err := h.requestValidator.Struct(&createTaskRequest); err != nil {
	// 	return res.ErrFieldValidationFailed(err)
	// }

	// Convert request to entity
	taskEntity := mapper.ConvertTaskRequestToTaskEntity(createTaskRequest)

	// Create task
	h.createTaskUsecase.ExecCreateTask(context.Background(), taskEntity)

	// Return response
	c.JSON(http.StatusCreated, response.SuccessResponse{
		Message: "Task created successfully",
	})

}
