package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tomyalberdi/go-rest-api/config"
	"github.com/tomyalberdi/go-rest-api/models"
	"net/http"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Preload("User").Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func CreatePost(c *gin.Context) {
	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Validate user
	var user models.User
	if err := config.DB.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	post := models.Post{Title: input.Title, Body: input.Body, UserID: user.ID}
	config.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
	// Get the post ID from the path
	var post models.Post
	id := c.Param("id")
	// Find the post
	if err := config.DB.Where("id = ?", id).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	// Delete the post
	config.DB.Delete(&post)
	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
