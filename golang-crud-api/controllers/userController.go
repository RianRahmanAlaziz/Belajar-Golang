package controllers

import (
	"golang-crud-api/config"
	"golang-crud-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var users models.User

	config.DB.First(&users, id)

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)
	config.DB.Create(&user)

	c.JSON(http.StatusOK, user)

}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	config.DB.First(&user, id)

	c.BindJSON(&user)

	config.DB.Save(&user)

	c.JSON(http.StatusOK, user)

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	config.DB.Delete(&user, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}
