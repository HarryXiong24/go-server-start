package main

import (
	"fmt"
	"go-server-start/db"
	"go-server-start/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	err := db.Init()
	if err != nil {
		fmt.Println("Failed to initialize database:", err)
		return
	}

	router := server.Group("/")
	routers.Init(router)

	err = server.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
