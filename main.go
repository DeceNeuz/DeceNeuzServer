package main

import (
	"os"

	"github.com/SubhranilRaha/go-crud/controllers"
	"github.com/SubhranilRaha/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	r := gin.Default()
	r.POST("/news", controllers.NewsCreate)
	PORT := os.Getenv("PORT")
	r.Run("localhost:"+PORT)
}