package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ooatamelbug/blog-task-app/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB() *gorm.DB {
	// load .env file
	err := godotenv.Load("./pkg/common/env/.env")
	if err != nil {
		panic("faild to load env")
	}

	// get out the config data to connect db
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// make uri from .env to connect db
	dns := fmt.Sprint("host=", dbHost, " user=", dbUser, " password=", dbPassword, " dbname=", dbName, " port=", dbPort)

	// open connection with postgres db with dns uri of db and pointer gorm config
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	// check if error occure in connect with db
	if err != nil {
		panic("faild to connect to db")
	}

	// make migrate for all model in app
	db.AutoMigrate(&models.User{}, &models.Post{})

	// return db to use it
	return db
}

// close connection db
func CloseConnectionDB(dbConnection *gorm.DB) {
	db, err := dbConnection.DB()
	if err != nil {
		panic("falid to close db")
	}
	db.Close()
}
