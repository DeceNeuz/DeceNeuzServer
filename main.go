package main

import (
	"github.com/SubhranilRaha/go-crud/controllers"
	"github.com/SubhranilRaha/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	// initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	r := gin.Default()
	r.POST("/news", controllers.NewsCreate)
	r.GET("/news", controllers.NewsIndex)
	r.POST("/user", controllers.UserCreate)
	r.POST("/login", controllers.LoginHandler)
	r.PATCH("/likes/:id", controllers.LikesUpdate)
	// PORT := 3000
	r.Run("localhost:3000")
}