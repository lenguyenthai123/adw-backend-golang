package handler

import (
	"backend-golang/modules/analytics/domain/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AnalyticsHandler interface {
	HandleGetUserProgress(c *gin.Context)
	HandleGetRatioTotalTimeSpent(c *gin.Context)
	HandleGetTimeSpentDaily(c *gin.Context)
	HandleGetTaskOfEachStatus(c *gin.Context)
	HandleGetAIFeedback(c *gin.Context)
}

type AnalyticsHandlerImpl struct {
	requestValidator              *validator.Validate
	getUserProgressUsecase        usecase.GetUserProgressUsecase
	getRatioTotalTimeSpentUsecase usecase.GetRatioTotalTimeSpentUsecase
	getTimeSpentDailyUsecase      usecase.GetTimeSpentDailyUsecase
	getTaskOfEachStatusUsecase    usecase.GetTaskOfEachStatusUsecase
	getAIFeedbackUsecase          usecase.GetAIFeedbackUsecase
}

func NewAnalyticsHandler(
	requestValidator *validator.Validate,
	getUserProgressUsecase usecase.GetUserProgressUsecase,
	getRatioTotalTimeSpentUsecase usecase.GetRatioTotalTimeSpentUsecase,
	getTimeSpentDailyUsecase usecase.GetTimeSpentDailyUsecase,
	getTaskOfEachStatusUsecase usecase.GetTaskOfEachStatusUsecase,
	getAIFeedbackUsecase usecase.GetAIFeedbackUsecase,
) AnalyticsHandler {
	return &AnalyticsHandlerImpl{
		requestValidator:              requestValidator,
		getUserProgressUsecase:        getUserProgressUsecase,
		getRatioTotalTimeSpentUsecase: getRatioTotalTimeSpentUsecase,
		getTimeSpentDailyUsecase:      getTimeSpentDailyUsecase,
		getTaskOfEachStatusUsecase:    getTaskOfEachStatusUsecase,
		getAIFeedbackUsecase:          getAIFeedbackUsecase,
	}
}
