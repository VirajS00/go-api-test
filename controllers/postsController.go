package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud/initializers"
	"github.com/go-crud/models"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsGetAll(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func FindPostById(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Post Not Found",
		})

		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PatchPostById(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Post Does not exist",
		})

		return
	}

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePostById(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(204)
}
