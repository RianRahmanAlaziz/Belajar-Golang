package controllers

import (
	"go-user-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Services services.UserService
}

func (u *UserController) GetUsers(c *gin.Context) {

	users, err := u.Services.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error",
		})
		return
	}
	c.JSON(http.StatusOK, users)

}
