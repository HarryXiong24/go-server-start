package handlers

import (
	"go-server-start/internal/services"
	api "go-server-start/internal/types"
	"go-server-start/pkg/errors"
	"go-server-start/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserInfo handles requests to get user information
func GetUserInfo(c *gin.Context) {
	var req api.GetUserInfoRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Sugar.Errorw("Failed to bind request", "error", err.Error())
		c.Error(errors.NewBadRequest("Invalid request parameters", err))
		return
	}

	res, err := services.GetUserInfo(c.Request.Context(), &req)
	if err != nil {
		logger.Sugar.Errorw("Failed to get user info", "error", err.Error())
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}
