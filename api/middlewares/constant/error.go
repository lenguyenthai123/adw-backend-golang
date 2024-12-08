package constant

import (
	res "backend-golang/core/response"
	"net/http"
)

func ErrInvalidToken(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusForbidden,
		err,
		"token is invalid signature",
		"ERR_INVALID_TOKEN",
	)
}

func ErrMissingToken(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusUnauthorized,
		err,
		"missing token in header",
		"ERR_UNAUTHORIZED",
	)
}

func ErrInternal(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong with the server",
		"ERR_INTERNAL",
	)
}

func ErrEmailNotFound(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusNotFound,
		err,
		"email not found",
		"ERR_EMAIL_NOT_FOUND",
	)
}

func ErrEmailOrPasswordInvalid(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusUnauthorized,
		err,
		"email or password invalid",
		"ERR_EMAIL_OR_PASSWORD_INVALID",
	)
}
