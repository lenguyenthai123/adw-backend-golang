package handler

import (
	"backend-golang/core"
	res "backend-golang/core/response"
	"backend-golang/modules/task/api/mapper"
	"backend-golang/modules/task/api/model/req"
	"backend-golang/modules/task/constant"
	"backend-golang/pkgs/log"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *TaskHandlerImpl) HandleApplyAnalyzedTask(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// Bind request
	var applyAnalyzedTaskRequest req.ApplyAnalyzedTaskRequest
	if err := c.ShouldBindJSON(&applyAnalyzedTaskRequest); err != nil {
		log.JsonLogger.Error("HandlerApplyAnalyzedTask.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)
		panic(res.ErrInvalidRequest(err))
	}

	for _, task := range applyAnalyzedTaskRequest.TaskList {
		if task.TaskID == "" {
			panic(constant.ErrorMissingTaskIDWhenUpdate())
		}
	}

	taskEntityList := mapper.ConvertUpdateTaskListToTaskEntityList(applyAnalyzedTaskRequest.TaskList)

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	err := h.applyAnalyzedTaskUsecase.ExecApplyAnalyzedTask(ctx, taskEntityList)

	if err != nil {
		panic(err)
	}
	// Return response
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "Apply change task list successfully", nil))

}
