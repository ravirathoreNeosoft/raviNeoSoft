package routes

import (
	controllers "gin-post-api/src/controller"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/create", controllers.CreateUser)
		userRoutes.GET("/users", controllers.GetAllUsers)
	}
}
