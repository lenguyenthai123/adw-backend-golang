package handler

import (
	res "backend-golang/core/response"
	"backend-golang/modules/user/api/mapper"
	"backend-golang/modules/user/api/model/req"
	"backend-golang/pkgs/log"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

// HandleLoginUser godoc
// @Summary      Login user
// @Description  Login user by email and password
// @Tags         User
// @Produce      json
// @Param        LoginUserReq  body	req.LoginUserReq  true  "LoginUserReq JSON"
// @Success      200 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /user/login [post]
func (h *userHandler) HandleLoginUser(c *gin.Context) {
	// Bind request
	var loginUserReq req.LoginUserReq
	if err := c.ShouldBindJSON(&loginUserReq); err != nil {
		log.JsonLogger.Error("HandleLoginUser.bind_json",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		panic(res.ErrInvalidRequest(err))
	}

	// Validate request
	if err := h.requestValidator.Struct(&loginUserReq); err != nil {
		log.JsonLogger.Error("HandleLoginUser.validate_request",
			slog.String("error", err.Error()),
			slog.String("request_id", c.Request.Context().Value("X-Request-ID").(string)),
		)

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				if e.Field() == "Email" {
					panic(res.ErrFieldValidationFailed(errors.New("invalid email")))
				}
			}
			// If no field matched, return default error
			panic(res.ErrFieldValidationFailed(err))
		}
	}

	// Login user
	user, err := h.loginUserUseCase.ExecLoginUser(c.Request.Context(),
		mapper.ConvertLoginUserReqToUserEntity(loginUserReq))
	if err != nil {
		panic(err)
	}

	// Convert user entity to response
	userResponse := mapper.CovertUserEntityToLoginUserRes(*user)

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", userResponse))
}
