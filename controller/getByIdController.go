package controller

import (
	"github.com/breno600/go-crud/initializers"
	"github.com/breno600/go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetByID (c *gin.Context) {
	id := c.Param("id")
	posts := models.Post{}
	result := initializers.DB.Find(&posts, id)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": posts,
	})
}