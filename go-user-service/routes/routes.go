package routes

import (
	"go-user-service/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
) {

	api := r.Group("/api")

	{
		api.GET("/users", userHandler.GetUsers)
		api.POST("/users", userHandler.CreateUser)
		api.POST("/login", authHandler.Login)
	}
}
