package main

import (
	"baas-fabric/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 22 // 8 MiB
	router.Static("/public/", "public")
	router.GET("/api/v1/ping",controller.Ping)
	router.POST("/api/v1/deployChainCode",controller.DeployCC)
	router.GET("/api/v1/getUser",controller.QueryChainCode)
	router.Run(":8080")
}