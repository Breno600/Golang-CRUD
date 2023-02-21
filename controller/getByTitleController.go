package controller

import (

	"github.com/breno600/go-crud/initializers"
	"github.com/breno600/go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetByTitle (c *gin.Context) {
	posts := models.Post{}

	var body struct{
		Title string 
	}

	c.Bind(&body)

	result := initializers.DB.Where("title = ?",body.Title).Find(&posts)

	if result.Error != nil {
		c.Status(400)
		return
	}
	
	c.JSON(200, gin.H{
		"post": posts,
	})
}