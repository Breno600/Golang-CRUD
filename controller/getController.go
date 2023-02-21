package controller

import (
	"github.com/breno600/go-crud/initializers"
	"github.com/breno600/go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetAll (c *gin.Context) {
	posts := []models.Post{}
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": posts,
	})
}