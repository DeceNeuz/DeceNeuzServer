package main

import (
	"github.com/SubhranilRaha/go-crud/initializers"
	"github.com/SubhranilRaha/go-crud/models"
)

func init(){
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.NewsPost{})
	initializers.DB.AutoMigrate(&models.UserDetails{})
}