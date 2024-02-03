package main

import (
	"github.com/SubhranilRaha/go-crud/controllers"
	"github.com/SubhranilRaha/go-crud/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init(){
	// initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/news", controllers.NewsCreate)
	r.GET("/news", controllers.NewsIndex)
	r.POST("/user", controllers.UserCreate)
	r.POST("/login", controllers.LoginHandler)
	r.PATCH("/likes/:id", controllers.LikesUpdate)
	// PORT := 3000
	// r.Run("localhost:3000")

	server := &http.Server{
		Addr:         ":3000", // Set the desired port
		Handler:      r,
		/* ReadTimeout:  15 * time.Second, // Adjust these values based on your application's needs
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second, */
	}

	// Start the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// Handle any errors that occur during server startup
		panic(err)
	}
}