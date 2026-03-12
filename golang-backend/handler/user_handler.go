package handler

import (
	"golang-backend/models"
	"golang-backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
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

	var user models.User

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
