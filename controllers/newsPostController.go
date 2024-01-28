package controllers

import (
	"github.com/SubhranilRaha/go-crud/initializers"
	"github.com/SubhranilRaha/go-crud/models"
	"github.com/gin-gonic/gin"
)

func NewsCreate(c *gin.Context) {
	//Get data off req body'
	var body struct {
		Title     string    
		Content   string    
		User     string    
		Likes     uint      
	}
	c.Bind(&body)
	//Create a post
	post := models.NewsPost{Title: body.Title, Content:body.Content, User:body.User, Likes:0}
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