package main

import (
	"gin-post-api/src/config"
	"gin-post-api/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	router := gin.Default()
	routes.InitializeRoutes(router)
	router.Run(":8080")
}
