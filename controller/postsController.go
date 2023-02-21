package controller

import (
	"github.com/breno600/go-crud/initializers"
	"github.com/breno600/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate (c *gin.Context) {

	var body struct{
		Body string
		Title string 
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body:body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}