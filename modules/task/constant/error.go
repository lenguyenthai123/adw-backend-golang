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
