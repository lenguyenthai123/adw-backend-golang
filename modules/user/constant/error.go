package constant

import (
	res "backend-golang/core/response"
	"net/http"
)

func ErrorEmailAlreadyExists(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"email already exists",
		"ERR_EMAIL_ALREADY_EXISTS",
	)
}

func ErrorInternalServerError(err error) error {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"internal server error",
		"ERR_INTERNAL_SERVER_ERROR",
	)
}

func ErrUserNotFound(err error) error {
	return res.NewErrorResponse(
		http.StatusNotFound,
		err,
		"user not found",
		"ERR_USER_NOT_FOUND",
	)
}

func ErrorPhoneAlreadyExists(err error) error {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"phone already exists",
		"ERR_PHONE_ALREADY_EXISTS",
	)
}

func ErrorWrongPassword(err error) error {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"wrong password",
		"ERR_WRONG_PASSWORD",
	)
}

func ErrorHashPassword(err error) error {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"hash password error",
		"ERR_HASH_PASSWORD",
	)
}
