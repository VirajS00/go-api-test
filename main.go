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

	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.PostsGetAll)
	r.GET("/post/:id", controllers.FindPostById)
	r.PATCH("/post/:id", controllers.PatchPostById)
	r.DELETE("/post/:id", controllers.DeletePostById)

	r.Run()
}
