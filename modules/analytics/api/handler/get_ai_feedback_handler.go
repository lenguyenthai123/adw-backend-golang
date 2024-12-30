package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *AnalyticsHandlerImpl) HandleGetAIFeedback(c *gin.Context) {
	// var getUserProgressRequest req.GetAnalyticsRequest
	// if err := c.ShouldBindJSON(&getUserProgressRequest); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// getUserProgressResponse, err := handler.getUserProgressUsecase.ExecGetUserProgress(c, getUserProgressRequest)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, "get ai feedback")
}
