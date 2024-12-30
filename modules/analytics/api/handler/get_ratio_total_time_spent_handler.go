package handler

import (
	"backend-golang/modules/analytics/api/model/req"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *AnalyticsHandlerImpl) HandleGetRatioTotalTimeSpent(c *gin.Context) {
	var getUserProgressRequest req.GetAnalyticsRequest
	if err := c.ShouldBindJSON(&getUserProgressRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	getUserProgressResponse, err := handler.getRatioTotalTimeSpentUsecase.ExecGetRatioTotalTimeSpent(c, getUserProgressRequest.StartTime, getUserProgressRequest.EndTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getUserProgressResponse)
}
