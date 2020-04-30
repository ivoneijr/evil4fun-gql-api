package db

import (
	"fmt"
	"graphql-api/models"
)

func Seed() {
	user := models.User{
		Email:    "ivoneijr@gmail.com",
		Name:     "ivo junior",
		Password: "password",
	}

	db, _ := Connection()
	db.Create(&user)

	fmt.Println(user)

	post := models.Post{
		Name:    "Post Name",
		Content: "<htmnl>content here</html>",
		UserId:  user.ID,
	}

	db.Create(&post)
}
