package postgres

import (
	"github.com/404th/bookstore/models"
	"github.com/jmoiron/sqlx"
)

type bookCategoryPg struct {
	db *sqlx.DB
}

func (bc *bookCategoryPg) GetBookCategory(id string) (*models.BookCategory, error) {
	// TODO
	return nil, nil
}

func (bc *bookCategoryPg) GetAllBookCategories(queryParam models.ApplicationQueryParamModel) ([]models.BookCategory, error) {
	// TODO
	return nil, nil
}

func (bc *bookCategoryPg) CreateBookCategory(details *models.BookCategory) (string, error) {
	// TODO
	return "", nil
}

func (bc *bookCategoryPg) UpdateBookCategory(details *models.UpdateBookCategory, id string) (int64, error) {
	// TODO
	return 0, nil
}

func (bc *bookCategoryPg) DeleteBookCategory(id string) (int64, error) {
	// TODO
	return 0, nil
}
