package models

// Article
type Article struct {
	Id        int    `json:"id" db:"id" validate:"omitempty"`
	Title     string `json:"title" db:"title" validate:"required,gte=20"`
	Content   string `json:"content" db:"content" validate:"required,gte=200"`
	Category  string `json:"category" db:"category" validate:"required,gte=3"`
	Status    string `json:"status" db:"status" validate:"required,oneof=publish draft trash"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
}

type ArticleList struct {
	Articles []*Article `json:"articles"`
	Meta     Meta       `json:"meta"`
}
