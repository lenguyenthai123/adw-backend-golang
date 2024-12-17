package handler

import (
	"backend-golang/core"
	res "backend-golang/core/response"
	"backend-golang/modules/task/api/mapper"
	"backend-golang/modules/task/api/model/req"
	"backend-golang/pkgs/log"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *TaskHandlerImpl) HandleDeleteTaskList(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var listIDRequest req.ListIDRequest
	if err := c.ShouldBindJSON(&listIDRequest); err != nil {
		log.JsonLogger.Error("HandleDeleteTaskList.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)
		panic(res.ErrInvalidRequest(err))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	taskIDEntityList := mapper.ConvertListIDRequestToListIDEntity(listIDRequest.TaskIDList)

	err := h.deleteTaskListUsecase.ExecDeleteTaskList(ctx, taskIDEntityList)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "Delete tasks successfully", nil))
}
