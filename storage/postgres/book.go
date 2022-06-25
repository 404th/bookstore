package postgres

import (
	"github.com/404th/bookstore/models"
	"github.com/jmoiron/sqlx"
)

type bookPg struct {
	db *sqlx.DB
}

func (bk *bookPg) GetBook(id string) (*models.Book, error) {
	// TODO
	return nil, nil
}

func (bk *bookPg) GetAllBooks() ([]*models.Book, error) {
	// TODO
	return nil, nil
}

func (bk *bookPg) CreateBook(details *models.Book) (string, error) {
	// TODO
	return "", nil
}

func (bk *bookPg) UpdateBook(details *models.UpdateBook, id string) (string, error) {
	// TODO
	return "", nil
}

func (bk *bookPg) DeleteBook(id string) (string, error) {
	// TODO
	return "", nil
}
