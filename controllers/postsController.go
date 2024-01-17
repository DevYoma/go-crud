package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yoma/go-crud/initializers"
	"github.com/yoma/go-crud/models"
)

// Create Post
func PostsCreate(c *gin.Context) {
	// Get data off req body
	
	// to get it off it, we create a struct to store the data
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
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

// Get Posts
func PostsIndex(c *gin.Context) {
	// Get all posts
	var posts []models.Post

	initializers.DB.Find(&posts)

	// Return it
	c.JSON(200, gin.H{ 
		"posts": posts,
	})
}

// Get Single Post
func PostsShow(c *gin.Context) {
	// Get the post
	var post models.Post

	// getting the id of the URL
	id := c.Param("id")

	initializers.DB.First(&post, id)

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

// Update Post
func PostsUpdate(c *gin.Context) {
	// get id 
	id := c.Param("id")

	// get data off req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// find post we're updating
	var post models.Post
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// send response
	c.JSON(200, gin.H{
		"post": post,
	})
}

// delete post

func PostsDelete(c *gin.Context) {
	// get id
	id := c.Param("id")

	// find post
	var post models.Post
	initializers.DB.First(&post, id)

	// delete post
	initializers.DB.Delete(&post)

	// send response
	c.JSON(200, gin.H{
		"post": post,
	})
}