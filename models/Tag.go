package models

type Tag struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Posts       []*Post `json:"posts"`
}
