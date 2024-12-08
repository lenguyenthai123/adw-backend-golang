package res

import (
	"errors"
)

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	ErrorKey   string `json:"error_key"`
}

// RootError returns the root error of the ResError.
//
// It does not take any parameters.
// It returns an error.
func (e *ErrorResponse) RootError() error {
	var err *ErrorResponse

	if errors.As(e.RootErr, &err) {
		return err.RootError()
	}

	return e.RootErr
}

// Error returns the error message of the ResError.
//
// It returns a string.
func (e *ErrorResponse) Error() string {
	return e.RootError().Error()
}

// NewErrorResponse creates a new custom error with the given code, root error, message, and key.
//
// Parameters:
//
//   - code: The error code.
//
//   - root: The root error.
//
//   - msg: The error message.
//
//   - key: The error key.
//
// Returns:
//   - *ResError: The new custom error.
func NewErrorResponse(statusCode int, root error, msg, errorKey string) *ErrorResponse {
	if root != nil {
		return &ErrorResponse{
			StatusCode: statusCode,
			RootErr:    root,
			Message:    msg,
			Log:        root.Error(),
			ErrorKey:   errorKey,
		}
	}

	return &ErrorResponse{
		StatusCode: statusCode,
		RootErr:    errors.New(msg),
		Message:    msg,
		Log:        msg,
		ErrorKey:   errorKey,
	}
}
