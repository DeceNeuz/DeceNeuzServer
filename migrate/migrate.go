package main

import (
	"github.com/SubhranilRaha/go-crud/initializers"
	"github.com/SubhranilRaha/go-crud/models"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.NewsPost{})
}