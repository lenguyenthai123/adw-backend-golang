package constant

import (
	res "backend-golang/core/response"
	"net/http"
)

func ErrorGetUserProgressFail(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusNotFound,
		err,
		"Cannot found user progress",
		"ERR_TASK_PROGRESS_NOT_FOUND",
	)
}

func ErrorNotAnyProgressToAnalyze(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusBadRequest,
		err,
		"Not found any task progress to analyze",
		"ERR_NOT_ANY_TASK_PROGRESS_TO_ANALYZE",
	)
}
