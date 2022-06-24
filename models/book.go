package models

import "time"

type Book struct {
	ID        string    `json:"id" db:"id" example:"123abd"`
	Name      string    `json:"name" db:"name" binding:"required" example:"Start with why"`
	Price     float32   `json:"price" db:"price" binding:"required" example:"18.99"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateBook struct {
	Name  string  `json:"name" db:"name" binding:"required" example:"Start with why"`
	Price float32 `json:"price" db:"price" binding:"required" example:"18.99"`
}
