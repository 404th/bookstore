package models

import "time"

type Book struct {
	ID         string    `json:"id" db:"id" example:"123abd"`
	CategoryID string    `json:"category_id" db:"category_id" binding:"required" example:"qwerty123"`
	AuthorID   string    `json:"author_id" db:"author_id" binding:"required"`
	Name       string    `json:"name" db:"name" binding:"required" example:"Start with why"`
	Price      float32   `json:"price" db:"price" binding:"required" example:"18.99"`
	Definition string    `json:"defenition" db:"defenition" example:"very poor book"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type CreateBook struct {
	CategoryID string  `json:"category_id" db:"category_id" binding:"required" example:"qwerty123"`
	AuthorID   string  `json:"author_id" db:"author_id" binding:"required" example:"author_id"`
	Name       string  `json:"name" db:"name" binding:"required" example:"Start with why"`
	Price      float32 `json:"price" db:"price" binding:"required" example:"18.99"`
	Definition string  `json:"defenition" db:"defenition" example:"created very poor book"`
}

type UpdateBook struct {
	CategoryID string    `json:"category_id" db:"category_id" example:"qwerty123"`
	AuthorID   string    `json:"author_id" db:"author_id" example:"qwerty123"`
	Name       string    `json:"name" db:"name" example:"start with why"`
	Price      float32   `json:"price" db:"price" example:"18.99"`
	Definition string    `json:"defenition" db:"defenition" example:"updated very poor book"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
