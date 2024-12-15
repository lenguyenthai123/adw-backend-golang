package handler

import (
	"backend-golang/core"
	res "backend-golang/core/response"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleCreateTask	godoc
// @Summary      Delete task
// @Description  Delete task
// @Tags         Task
// @Produce      json
// @Param        task_id  path	string	true  "Task ID"
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /task/:task_id [get]

func (h *TaskHandlerImpl) HandleGetTask(c *gin.Context) {
	taskID := c.Param("task_id")

	if taskID == "" {
		panic(res.ErrInvalidRequest(errors.New("task_id is required")))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	task, err := h.getTaskUsecase.ExecGetTask(ctx, taskID)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", task))
}
