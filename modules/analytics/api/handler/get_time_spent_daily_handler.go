package handler

import (
	"backend-golang/core"
	"backend-golang/modules/analytics/api/model/req"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *AnalyticsHandlerImpl) HandleGetTimeSpentDaily(c *gin.Context) {
	var getUserProgressRequest req.GetAnalyticsRequest
	if err := c.ShouldBindJSON(&getUserProgressRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	getUserProgressResponse, err := handler.getTimeSpentDailyUsecase.Execute(ctx, getUserProgressRequest.StartTime, getUserProgressRequest.EndTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getUserProgressResponse)
}
