package handler

import (
	res "backend-golang/core/response"
	"backend-golang/modules/task/domain/entity"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleGetTasks godoc
// @Summary      Get tasks
// @Description  Get tasks with search, filter, and sort functionality
// @Tags         Task
// @Produce      json
// @Param        search  query   string  false  "Search term"
// @Param        status  query   string  false  "Task status (e.g., Todo, InProgress, Completed)"
// @Param        sortBy  query   string  false  "Sort by field (e.g., priority, createdAt)"
// @Param        order   query   string  false  "Sort order (asc or desc)"
// @Param        limit   query   int     false  "Limit per page"
// @Param        page    query   int     false  "Page number"
// @Success      200     {object} res.SuccessResponse{data=[]entity.Task}
// @Failure      400     {object} res.ErrorResponse
// @Failure      500     {object} res.ErrorResponse
// @Router       /tasks [get]

func (h *TaskHandlerImpl) HandleGetTaskList(c *gin.Context) {
	var taskSearchFilter entity.TaskSearchFilterEntity

	// Bind the query parameters to the TaskSearchFilterEntity
	if err := c.ShouldBindQuery(&taskSearchFilter); err != nil {
		panic(res.ErrInvalidRequest(err))
	}

	// Validate and parse Limit and Page to integers
	if taskSearchFilter.Limit != nil && *taskSearchFilter.Limit <= 0 {
		panic(res.ErrInvalidRequest(errors.New("limit must be greater than 0")))
	}
	if taskSearchFilter.Page != nil && *taskSearchFilter.Page <= 0 {
		panic(res.ErrInvalidRequest(errors.New("page must be greater than 0")))
	}

	// Fetch tasks using usecase
	tasks, err := h.getTaskListUsecase.ExecGetTaskList(context.Background(), taskSearchFilter)
	if err != nil {
		panic(res.ErrInternalServerError(err)) // Assuming you have a response helper function
	}

	// Respond with success
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "Tasks retrieved successfully", tasks))
}
