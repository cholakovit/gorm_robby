package controllers

import (
	"gorm_robby/initializers"
	"gorm_robby/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Get data off req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// Create a post
	 post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return it

	c.JSON(200, gin.H{
		"post": posts,
	})

	// Respond with them
}

func PostShow(c *gin.Context) {
	//get id off url
	id := c.Param("id")

	// Get the posts
	var post []models.Post
	initializers.DB.First(&post, id)

	// Return it

	c.JSON(200, gin.H{
		"post": post,
	})

	// Respond with them
}

func PostUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// Find the post were updating
	var post []models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).UpdateColumns(models.Post{Title: body.Title, Body: body.Body})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Respond
	initializers.DB.Delete(&models.Post{}, id)
}