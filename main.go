package main

import (
	"net/http"

	"github.com/SubhranilRaha/go-crud/controllers"
	"github.com/SubhranilRaha/go-crud/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/news", controllers.NewsCreate)
	r.GET("/news", controllers.NewsIndex)
	r.POST("/user", controllers.UserCreate)
	r.POST("/login", controllers.LoginHandler)
	r.PATCH("/likes/:id", controllers.LikesUpdate)

	server := &http.Server{
		Addr:    ":3000", // Set the desired port
		Handler: r,
	}

	// Start the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// Handle any errors that occur during server startup
		panic(err)
	}
}
