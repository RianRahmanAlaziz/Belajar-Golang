package handler

import (
	"go-user-service/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.UserService
}

func NewAuthHandler(s *services.UserService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Login(c *gin.Context) {

	var request struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	token, err := h.service.Login(request.Email)

	if err != nil {
		c.JSON(401, gin.H{
			"message": "login failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
