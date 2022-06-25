package postgres

import (
	"github.com/404th/bookstore/models"
	"github.com/jmoiron/sqlx"
)

type authorPg struct {
	db *sqlx.DB
}

func (ar *authorPg) GetAuthor(id string) (*models.Author, error) {
	// TODO
	return nil, nil
}

func (ar *authorPg) GetAllAuthors() ([]models.Author, error) {
	// TODO
	return nil, nil
}

func (ar *authorPg) CreateAuthor(details *models.Author) (string, error) {
	// TODO
	return "", nil
}

func (ar *authorPg) UpdateAuthor(details *models.UpdateAuthor, id string) (string, error) {
	// TODO
	return "", nil
}

func (ar *authorPg) DeleteAuthor(id string) (string, error) {
	// TODO
	return "", nil
}
