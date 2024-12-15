package handler

import (
	"backend-golang/core"
	res "backend-golang/core/response"
	"backend-golang/modules/task/api/mapper"
	"backend-golang/modules/task/api/model/req"
	"backend-golang/pkgs/log"
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *TaskHandlerImpl) HandleAnalyzeTask(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// Bind request
	var analyzeTaskRequest req.AnalyzeTaskRequest
	if err := c.ShouldBindJSON(&analyzeTaskRequest); err != nil {
		log.JsonLogger.Error("HandlerAnalyzeTask.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)
		panic(res.ErrInvalidRequest(err))
	}

	// Create task
	startTime, err1 := mapper.ConvertToUTC7(analyzeTaskRequest.StartTime)
	endTime, err2 := mapper.ConvertToUTC7(analyzeTaskRequest.EndTime)

	if err1 != nil || err2 != nil {
		log.JsonLogger.Error("HandlerAnalyzeTask.parse_time",
			slog.String("error", err1.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)
		panic(res.ErrInvalidRequest(err1))
	}
	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	openaiResponse, err := h.analyzeTaskUsecase.ExecAnalyzeTask(ctx, startTime, endTime)

	if err != nil {
		panic(err)
	}
	// Return response
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", openaiResponse))

}
