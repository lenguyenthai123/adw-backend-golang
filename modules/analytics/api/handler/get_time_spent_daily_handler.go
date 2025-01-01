package handler

import (
	"backend-golang/core"
	res "backend-golang/core/response"
	"backend-golang/modules/analytics/api/model/req"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *AnalyticsHandlerImpl) HandleGetTimeSpentDaily(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var getUserProgressRequest req.GetAnalyticsRequest
	if err := c.ShouldBindJSON(&getUserProgressRequest); err != nil {
		panic(res.ErrInvalidRequest(err))
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	getUserProgressResponse, err := handler.getTimeSpentDailyUsecase.Execute(ctx, getUserProgressRequest.StartTime, getUserProgressRequest.EndTime)
	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", getUserProgressResponse))
}
