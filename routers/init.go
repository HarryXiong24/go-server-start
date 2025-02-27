package routers

import (
	"github.com/gin-gonic/gin"
	"go-server-start/handlers"
)

func Init(router *gin.RouterGroup) {
	router.GET("/get-user-info", handlers.GetUserInfo)
}
