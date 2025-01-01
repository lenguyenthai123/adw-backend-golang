package handler

import (
	"backend-golang/core"
	res "backend-golang/core/response"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleCreateTask	godoc
// @Summary      Delete task
// @Description  Delete task
// @Tags         Task
// @Produce      json
// @Param
// @Success      201 {object}  	res.SuccessResponse
// @failure		 400 {object} 	res.ErrorResponse
// @failure		 500 {object} 	res.ErrorResponse
// @Router       /task/number-by-status [get]

func (h *TaskHandlerImpl) HandleGetTaskNumberEachStatus(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// Get user from context, require middleware
	ctx := context.WithValue(c.Request.Context(), core.CurrentRequesterKeyStruct{},
		c.MustGet(core.CurrentRequesterKeyString).(core.Requester))

	taskNumber, err := h.getTaskNumberEachStatusUsecase.ExecGetTaskNumberEachStatus(ctx)

	if err != nil {
		panic(err)
	}

	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", taskNumber))
}
