package controllers

import (
	"github.com/SubhranilRaha/go-crud/initializers"
	"github.com/SubhranilRaha/go-crud/models"
	"github.com/gin-gonic/gin"
)

func NewsCreate(c *gin.Context) {
	//Get data off req body'
	var body struct {
		Title            string
		Content          string
		ShortDescription string
		User             string
		Likes            uint
	}
	c.Bind(&body)
	//Create a post
	post := models.NewsPost{Title: body.Title, Content: body.Content, ShortDescription: body.ShortDescription, User: body.User, Likes: 0}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func NewsIndex(c *gin.Context) {
	var news []models.NewsPost
	initializers.DB.Find(&news)

	c.JSON(200, gin.H{
		"news": news,
	})
}

func LatestNewsIndex(c *gin.Context) {
	var news []models.NewsPost
	initializers.DB.Order("created_at desc").Find(&news)

	c.JSON(200, gin.H{
		"news": news,
	})
}

func LikesUpdate(c *gin.Context) {
	//Get data off req body'
	var body struct {
		IsLike bool
	}
	c.Bind(&body)
	var record models.NewsPost
	result := initializers.DB.First(&record, c.Param("id"))
	if result.Error != nil {
		c.JSON(400, "Post not found")
		return
	}
	// increment like
	if body.IsLike {
		record.Likes++
	} else { // decrement like
		record.Likes--
	}
	result = initializers.DB.Save(&record)

	if result.Error != nil {
		c.JSON(400, "Could not update DB")
	}

	c.JSON(200, gin.H{
		"post": record,
	})
}
