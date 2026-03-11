package handler

import (
	"go-user-service/models"
	"go-user-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(c *gin.Context) {

	users, err := h.service.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var user models.Users

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})

		return
	}

	err := h.service.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed create user",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
