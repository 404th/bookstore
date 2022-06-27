package models

import "time"

type Author struct {
	ID         string    `json:"id" db:"id" binding:"required" example:"qwerty123"`
	Firstname  string    `json:"firstname" db:"firstname" binding:"required" example:"John"`
	Secondname string    `json:"secondname" db:"secondname" binding:"required" example:"Doe"`
	Email      string    `json:"email" db:"email" example:"example@example.com"`
	Age        uint16    `json:"age" db:"age" example:"25"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type CreateAuthor struct {
	Firstname  string `json:"firstname" db:"firstname" binding:"required" example:"Created John"`
	Secondname string `json:"secondname" db:"secondname" binding:"required" example:"created secondname"`
	Email      string `json:"email" db:"email" example:"created email"`
	Age        uint16 `json:"age" db:"age" example:"22"`
}

type UpdateAuthor struct {
	Firstname  string `json:"firstname" db:"firstname" example:"Updated John"`
	Secondname string `json:"secondname" db:"secondname" example:"Updated Doe"`
	Email      string `json:"email" db:"email" example:"updatedexample@example.com"`
	Age        uint16 `json:"age" db:"age" example:"26"`
}
