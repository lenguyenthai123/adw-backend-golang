package api

import (
	"backend-golang/api/routes"
	"backend-golang/config"
	"backend-golang/config/models"
	"backend-golang/pkgs/dbs/postgres"
	"backend-golang/pkgs/transport/http/server"
	"backend-golang/utils"
	"time"

	"github.com/openai/openai-go"
)

func NewServer() (*server.HTTPServer, error) {
	appConfig := config.LoadConfig(&models.AppConfig{}).(*models.AppConfig)
	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)
	//openaiConfig := config.LoadConfig(&models.OpenaiConfig{}).(*models.OpenaiConfig)

	connection := postgres.Connection{
		SSLMode:               postgres.Require,
		Host:                  dbConfig.DBHost,
		Port:                  dbConfig.DBPort,
		Database:              dbConfig.DBName,
		User:                  dbConfig.DBUserName,
		Password:              dbConfig.DBPassword,
		MaxIdleConnections:    dbConfig.MaxIdleConnections,
		MaxOpenConnections:    dbConfig.MaxOpenConnections,
		ConnectionMaxIdleTime: time.Duration(dbConfig.ConnectionMaxIdleTime),
		ConnectionMaxLifeTime: time.Duration(dbConfig.ConnectionMaxLifeTime),
	}

	db := postgres.InitDatabase(connection)
	openaiClient := openai.NewClient() // Thay bằng API key của bạn

	s := server.NewHTTPServer(
		server.AddName(appConfig.ServiceName),
		server.AddPort(appConfig.ServicePort),
		server.SetGracefulShutdownTimeout(time.Duration(appConfig.ServiceTimeout)),
	)

	requestValidator := utils.NewValidator()

	srv := &routes.RouteHandler{
		UserHandler:      routes.NewUserHandler(db, requestValidator),
		TaskHandler:      routes.NewTaskHandler(db, openaiClient, requestValidator),
		AnalyticsHandler: routes.NewAnalyticsHandler(db, openaiClient, requestValidator),
	}
	s.AddGroupRoutes(srv.InitGroupRoutes())

	return s, nil
}
