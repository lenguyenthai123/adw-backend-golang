package handler

import (
	res "backend-golang/core/response"
	"backend-golang/modules/task/api/model/req"
	response "backend-golang/modules/task/api/model/res"
	"backend-golang/pkgs/log"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

func (h *TaskHandlerImpl) HandleAnalyzeTask(c *gin.Context) {
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
	startTime, err1 := time.Parse(time.RFC3339, analyzeTaskRequest.StartTime)
	endTime, err2 := time.Parse(time.RFC3339, analyzeTaskRequest.EndTime)

	if err1 != nil || err2 != nil {
		log.JsonLogger.Error("HandlerAnalyzeTask.parse_time",
			slog.String("error", err1.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)
		panic(res.ErrInvalidRequest(err1))
	}

	err := h.analyzeTaskUsecase.ExecAnalyzeTask(c.Request.Context(), startTime, endTime)

	if err != nil {
		panic(err)
	}
	// Return response
	c.JSON(http.StatusCreated, response.SuccessResponse{
		Message: "Task created successfully",
	})

}
