package main

import (
	"fmt"
	"golang-backend/config"
	"golang-backend/database"
	"golang-backend/handler"
	"golang-backend/models"
	"golang-backend/repository"
	"golang-backend/routes"
	"golang-backend/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, _ := config.LoadConfig()

	db, err := database.ConnectDB(cfg)

	if err != nil {
		panic(err)
	}

	// AUTO MIGRATE
	db.AutoMigrate(&models.User{})

	r := gin.Default()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	routes.SetupRoutes(r, userHandler)
	fmt.Println("Server Running on port ", cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}
