package main

import (
	"github.com/gin-gonic/gin"
	"github.com/breno600/go-crud/initializers"
	"github.com/breno600/go-crud/controller"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){

	r := gin.Default()
	r.POST("/create", controller.PostsCreate)
	r.GET("/getAll", controller.GetAll)
	r.GET("/getByTitle", controller.GetByTitle)
	r.GET("/get/:id", controller.GetByID)
	r.Run()
	//fmt.Println("HEllo Word")
}