package res

import "net/http"

func ErrFieldValidationFailed(err error) *ErrorResponse {
	return NewErrorResponse(
		http.StatusBadRequest,
		err,
		err.Error(),
		"ERR_FIELD_VALIDATION",
	)
}

func ErrInvalidRequest(err error) *ErrorResponse {
	return NewErrorResponse(
		http.StatusBadRequest,
		err,
		"invalid request",
		"ERR_INVALID_REQUEST",
	)
}

func ErrInternalServerError(err error) *ErrorResponse {
	return NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"internal server error",
		"ERR_INTERNAL_SERVER",
	)
}
