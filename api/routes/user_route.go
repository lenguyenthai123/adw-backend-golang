package routes

import (
	"backend-golang/config"
	"backend-golang/config/models"
	userHandler "backend-golang/modules/user/api/handler"
	"backend-golang/modules/user/domain/usecase"
	"backend-golang/modules/user/repository"
	database "backend-golang/pkgs/dbs/postgres"
	"backend-golang/pkgs/jwt"
	"backend-golang/pkgs/transport/http/method"
	"backend-golang/pkgs/transport/http/route"
	"backend-golang/utils"
	"github.com/go-playground/validator/v10"
)

func NewUserHandler(db *database.Database, requestValidator *validator.Validate) userHandler.UserHandler {
	// Config
	//mailConfig := config.LoadConfig(&models.MailConfig{}).(*models.MailConfig)
	jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)

	// User Repository
	userRepoReader := repository.NewUserReaderRepository(*db)
	userRepoWriter := repository.NewUserWriterRepository(*db)

	hashAlgo := utils.NewHashAlgo()
	jwtToken := jwt.NewJWT(jwtConfig.JWTSecretKey)
	//mailer := mail.NewMailer(mailConfig)

	return userHandler.NewUserHandler(requestValidator,
		usecase.NewCreateUserUseCase(hashAlgo, userRepoReader, userRepoWriter),
		usecase.NewLoginUserUseCase(jwtToken, jwtConfig.AccessTokenExpiry,
			jwtConfig.RefreshTokenExpiry, hashAlgo, userRepoReader))
}

func (r *RouteHandler) userRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/v1/user",
		Routes: []route.Route{
			{
				Path:    "/signup",
				Method:  method.POST,
				Handler: r.UserHandler.HandleCreateUser,
			},
			{
				Path:    "/login",
				Method:  method.POST,
				Handler: r.UserHandler.HandleLoginUser,
			},
		},
	}
}
