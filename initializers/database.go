package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(){

	var err error
	dsn := "host=arjuna.db.elephantsql.com user=zpzenhka password=va0hKEJBMlferhvZOvqgll5uv1u-VT90 dbname=zpzenhka port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to db")
	}
}