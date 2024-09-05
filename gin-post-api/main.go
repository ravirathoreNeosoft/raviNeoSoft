package main

import (
	"gin-post-api/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitializeRoutes(router)
	router.Run(":8080")
}
