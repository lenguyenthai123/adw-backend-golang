package handler

import (
	"backend-golang/core"
	res "backend-golang/core/response"
	"backend-golang/modules/task/api/mapper"
	"backend-golang/modules/task/api/model/req"
	response "backend-golang/modules/task/api/model/res"
	"backend-golang/pkgs/log"
	"context"
	"log/slog"
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
	if err := c.ShouldBindJSON(&createTaskRequest); err != nil {
		log.JsonLogger.Error("HandleCreateTask.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	// if err := h.requestValidator.Struct(&createTaskRequest); err != nil {
	// 	return res.ErrFieldValidationFailed(err)
	// }

	// Convert request to entity
	taskEntity := mapper.ConvertCreateTaskRequestToTaskEntity(createTaskRequest)

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	// Create task
	err := h.createTaskUsecase.ExecCreateTask(ctx, taskEntity)

	if err != nil {
		panic(err)
	}
	// Return response
	c.JSON(http.StatusCreated, response.SuccessResponse{
		Message: "Task created successfully",
	})

}
