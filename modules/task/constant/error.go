package constant

import (
	res "backend-golang/core/response"
	"net/http"
)

func ErrorNotFoundTask(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusNotFound,
		err,
		"Not found task",
		"ERR_TASK_NOT_FOUND",
	)
}

func ErrorNotFoundTaskList(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusNotFound,
		err,
		"Not found any task",
		"ERR_TASK_LIST_NOT_FOUND",
	)
}

func ErrorCreateTaskFail(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"Create task not success",
		"ERR_CREATE_TASK_FAIL",
	)
}

func ErrorMissingTaskIDWhenUpdate() *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		nil,
		"Missing task id when update",
		"ERR_MISSING_TASK_ID_WHEN_UPDATE",
	)
}

func ErrrorUpdateTaskListFailed(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"Update task list failed",
		"ERR_UPDATE_TASK_LIST_FAILED",
	)
}

func ErrrorDeleteTaskListFailed(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"Delete task list failed",
		"ERR_DELETE_TASK_LIST_FAILED",
	)
}
func ErrrorTaskIDNotInteger(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"Task ID must be integer",
		"ERR_TASK_ID_NOT_INTEGER",
	)
}

func ErrorNotAnyTaskToAnalyze(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"Not found any task to analyze",
		"ERR_NOT_ANY_TASK_TO_ANALYZE",
	)
}

func ErrorSystem(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		nil,
		"System error",
		"ERR_SYSTEM",
	)
}
