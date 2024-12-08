package handler

import (
	"backend-golang/modules/user/domain/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler interface {
	HandleCreateUser(c *gin.Context)
	HandleLoginUser(c *gin.Context)
}

type userHandler struct {
	requestValidator  *validator.Validate
	createUserUseCase usecase.CreateUserUseCase
	loginUserUseCase  usecase.LoginUserUseCase
}

func NewUserHandler(
	requestValidator *validator.Validate,
	createUserUseCase usecase.CreateUserUseCase,
	loginUserUseCase usecase.LoginUserUseCase,

) UserHandler {
	return &userHandler{
		requestValidator:  requestValidator,
		createUserUseCase: createUserUseCase,
		loginUserUseCase:  loginUserUseCase,
	}
}
