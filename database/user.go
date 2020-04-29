package database

import (
	"graphql-api/models"

	"github.com/go-pg/pg/v9"
)

type UsersRepo struct {
	DB *pg.DB
}

func (u *UsersRepo) GetUsers() ([]*models.User, error) {
	var users []*models.User

	err := u.DB.Model(&users).Select()

	if err != nil {
		return nil, err
	}

	return users, nil
}
