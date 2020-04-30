package models

import "time"

type Post struct {
	ID        uint `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Content     string `json:"content"`
	UserId        uint  `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}
