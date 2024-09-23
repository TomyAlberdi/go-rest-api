package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tomyalberdi/go-rest-api/config"
	"github.com/tomyalberdi/go-rest-api/models"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Name: input.Name, Email: input.Email}
	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// Delete associated posts before deleting the user (manual cascading delete)
	config.DB.Where("user_id = ?", user.ID).Delete(&models.Post{})

	config.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User and associated posts deleted"})
}
