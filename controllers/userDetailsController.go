package controllers

import (
	"github.com/SubhranilRaha/go-crud/initializers"
	"github.com/SubhranilRaha/go-crud/models"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	//Get data off req body'
	var body struct {
		Email     string
		Password   string
		UserName     string
	}
	c.Bind(&body)
	//Create an user
	user := models.UserDetails{Email: body.Email, Password: body.Password, UserName: body.UserName}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"user": user,
	})
}