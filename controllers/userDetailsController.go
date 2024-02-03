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

func LoginHandler(c *gin.Context) {
	var body struct {
		Email     string
		Password   string
	}
	c.Bind(&body)

	var user models.UserDetails
	result := initializers.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Email not found"})
		return
	}

	if body.Password == user.Password {
		c.JSON(200, gin.H{"message": "Successful login"})
	} else {
		c.JSON(401, gin.H{"error": "Incorrect password"})
	}
}