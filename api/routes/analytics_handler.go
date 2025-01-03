package routes

import (
	"backend-golang/api/middlewares"
	analyticsHandler "backend-golang/modules/analytics/api/handler"
	"backend-golang/modules/analytics/domain/usecase"
	"backend-golang/modules/analytics/repository"
	database "backend-golang/pkgs/dbs/postgres"
	"backend-golang/pkgs/transport/http/method"
	"backend-golang/pkgs/transport/http/route"

	"github.com/go-playground/validator/v10"
	"github.com/openai/openai-go"
)

func NewAnalyticsHandler(db *database.Database, openaiClient *openai.Client, requestValidator *validator.Validate) analyticsHandler.AnalyticsHandler {
	timeProgressReaderRepo := repository.NewTimeProgressReaderRepository(*db)

	return analyticsHandler.NewAnalyticsHandler(requestValidator,
		usecase.NewGetUserProgressUsecase(timeProgressReaderRepo),
		usecase.NewGetRatioTotalTimeSpentUsecase(timeProgressReaderRepo),
		usecase.NewGetTimeSpentDailyUsecase(timeProgressReaderRepo),
		usecase.NewGetTaskOfEachStatusUsecase(timeProgressReaderRepo),
		usecase.NewGetAIFeedbackUsecase(timeProgressReaderRepo, openaiClient),
	)
}

func (r *RouteHandler) analyticsRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/analytics",
		Routes: []route.Route{
			{
				Path:    "/daily-spent-time",
				Method:  method.POST,
				Handler: r.AnalyticsHandler.HandleGetTimeSpentDaily,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/user-progress",
				Method:  method.GET,
				Handler: r.AnalyticsHandler.HandleGetUserProgress,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
			{
				Path:    "/ai-feedback",
				Method:  method.GET,
				Handler: r.AnalyticsHandler.HandleGetAIFeedback,
				Middlewares: route.Middlewares(
					middlewares.Authentication(),
				),
			},
		},
	}
}
