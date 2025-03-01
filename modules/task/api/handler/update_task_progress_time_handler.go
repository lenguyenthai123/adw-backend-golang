package handler

import (
	"backend-golang/core"
	res "backend-golang/core/response"
	"backend-golang/modules/task/api/mapper"
	"backend-golang/modules/task/api/model/req"
	"backend-golang/pkgs/log"
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleCreateTask	godoc
// @Summary      Update task progress time
// @Description  Update task progress time
// @Tags         Task
// @Produce      json
// @Param        UpdateTaskProgressRequest  body	req.UpdateTaskProgressRequest  true  "UpdateTaskProgressRequest JSON"
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /task/progress_time/:task_id [patch]

func (h *TaskHandlerImpl) HandleUpdateTaskProgressTime(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	taskID := c.Param("task_id")

	if taskID == "" {
		panic(res.ErrInvalidRequest(errors.New("task_id is required")))
	}

	// Bind request
	var updateTaskProgressListRequest req.UpdateTaskProgressListRequest
	if err := c.ShouldBind(&updateTaskProgressListRequest); err != nil {
		log.JsonLogger.Error("HandleUpdateTaskProgress.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := updateTaskProgressListRequest.Validate(); err != nil {
		panic(res.ErrInvalidRequest(err))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	userId := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserIDInt()

	// Update task
	if err := h.updateTaskProgressTimeUsecase.ExecUpdateTaskProgressList(
		ctx,
		mapper.ConvertUpdateTaskProgressListRequestToTaskEntityList(updateTaskProgressListRequest, taskID, userId),
	); err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", nil))

}
