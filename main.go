package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud/controllers"
	"github.com/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsGetAll)
	r.GET("/post/:id", controllers.PostShow)
	r.PATCH("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()
}
