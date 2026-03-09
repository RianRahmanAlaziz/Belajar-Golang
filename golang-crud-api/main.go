package main

import (
	"golang-crud-api/config"
	"golang-crud-api/controllers"
	"golang-crud-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})

	api := r.Group("/api")
	{
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.POST("/users", controllers.CreateUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)
	}

	r.Run(":8080")
}
