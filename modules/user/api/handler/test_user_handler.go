package handler

import (
	res "backend-golang/core/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *userHandler) HandleTest(c *gin.Context) {
	res.ResponseSuccess(c, res.NewSuccessResponse(http.StatusOK, "success", "test"))
}
