package main

import (
	"fmt"
	"go-user-service/config"
	"go-user-service/database"
	"go-user-service/handler"
	"go-user-service/repository"
	"go-user-service/routes"
	"go-user-service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, _ := config.LoadConfig()

	db, err := database.ConntectDB(cfg)

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(userService)

	// routes
	routes.SetupRoutes(r, userHandler, authHandler)

	fmt.Println("Server running on port", cfg.AppPort)

	r.Run(":" + cfg.AppPort)
}
