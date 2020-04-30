package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"graphql-api/models"
)
import _ "github.com/jinzhu/gorm/dialects/postgres"

func Connection() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=devil4fun sslmode=disable")
	db.LogMode(true)

	return db, err
}

func Ping() {
	db, err := Connection()

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	database := db.DB()

	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("[db] Connection to PostgreSQL was successful!")
}

func AutoMigrate() {
	db, err := Connection()

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Post{})

	fmt.Println("[db] AutoMigrate was successful")
}
