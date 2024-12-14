package handler

import (
	"backend-golang/core"
	res "backend-golang/core/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *userHandler) HandleTest(c *gin.Context) {
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", c.Value(core.CurrentRequesterKeyString).(core.Requester).GetUserID()))
}
