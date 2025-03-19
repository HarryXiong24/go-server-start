package routers

import (
	"go-server-start/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	router.GET("/get-user-info", handlers.GetUserInfo)
}
