package handlers

import (
	"fmt"
	"go-server-start/api"
	"go-server-start/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {

	var req api.GetUserInfoRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := services.GetUserInfo(&req)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
