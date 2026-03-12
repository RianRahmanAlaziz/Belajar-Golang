package routes

import (
	"golang-backend/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	userHanlder *handler.UserHandler) {
	api := r.Group("/api")
	{
		api.GET("/users", userHanlder.GetUsers)
		api.POST("/users", userHanlder.CreateUser)
	}
}
