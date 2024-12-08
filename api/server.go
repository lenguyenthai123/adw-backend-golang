package api

import (
	"backend-golang/api/routes"
	"backend-golang/config"
	"backend-golang/config/models"
	"backend-golang/pkgs/dbs/postgres"
	"backend-golang/pkgs/transport/http/server"
	"backend-golang/utils"
	"time"
)

func NewServer() (*server.HTTPServer, error) {
	appConfig := config.LoadConfig(&models.AppConfig{}).(*models.AppConfig)
	dbConfig := config.LoadConfig(&models.DBConfig{}).(*models.DBConfig)

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

	s := server.NewHTTPServer(
		server.AddName(appConfig.ServiceName),
		server.AddPort(appConfig.ServicePort),
		server.SetGracefulShutdownTimeout(time.Duration(appConfig.ServiceTimeout)),
	)

	requestValidator := utils.NewValidator()

	srv := &routes.RouteHandler{
		UserHandler: routes.NewUserHandler(db, requestValidator),
	}
	s.AddGroupRoutes(srv.InitGroupRoutes())

	return s, nil
}
