package handler

import (
	"backend-golang/core"
	res "backend-golang/core/response"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *AnalyticsHandlerImpl) HandleGetAIFeedback(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	openaiResponse, err := handler.getAIFeedbackUsecase.ExecuteGetAIFeedback(ctx)

	if err != nil {
		panic(err)
	}
	// Return response
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", openaiResponse))
}
