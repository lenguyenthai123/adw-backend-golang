package handler

import (
	"backend-golang/core"
	res "backend-golang/core/response"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *AnalyticsHandlerImpl) HandleGetUserProgress(c *gin.Context) {
	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	getUserProgressResponse, err := handler.getUserProgressUsecase.ExecuteGetProgress(ctx)
	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", getUserProgressResponse))
}
