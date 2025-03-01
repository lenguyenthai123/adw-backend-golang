package handler

import (
	"backend-golang/core"
	res "backend-golang/core/response"
	"context"
	"errors"
	"net/http"
	"strings"

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
// @Router       /task/:task_id [delete]

func (h *TaskHandlerImpl) HandleDeleteTask(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	taskID := c.Param("task_id")

	if taskID == "" {
		panic(res.ErrInvalidRequest(errors.New("task_id is required")))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	err := h.deleteTaskUsecase.ExecDeleteTask(ctx, taskID)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", nil))
}

func (h *TaskHandlerImpl) HandleDeleteTaskList(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// Get task IDs from query parameter
	taskIDsQuery := c.Query("task_ids")
	if taskIDsQuery == "" {
		panic(res.ErrInvalidRequest(errors.New("task_ids is required")))
	}

	taskIDs := strings.Split(taskIDsQuery, ",")

	// Validate if at least one task ID is provided
	if len(taskIDs) == 0 {
		panic(res.ErrInvalidRequest(errors.New("no task IDs provided")))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	err := h.deleteTaskListUsecase.ExecDeleteTaskList(ctx, taskIDs)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", nil))
}
