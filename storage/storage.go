package storage

import "github.com/404th/bookstore/models"

type StorageI interface {
	CloseDB() error
	BookCategoryRepo() BookCategoryI
	BookRepo() BookI
	AuthorRepo() AuthorI
}

type BookCategoryI interface {
	GetBookCategory(id string) (*models.BookCategory, error)
	GetAllBookCategories() ([]models.BookCategory, error)
	CreateBookCategory(details *models.BookCategory) (string, error)
	UpdateBookCategory(details *models.UpdateBookCategory, id string) (string, error)
	DeleteBookCategory(id string) (string, error)
}

type BookI interface {
	GetBook(id string) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	CreateBook(details *models.Book) (string, error)
	UpdateBook(details *models.UpdateBook, id string) (string, error)
	DeleteBook(id string) (string, error)
}

type AuthorI interface {
	GetAuthor(id string) (*models.Author, error)
	GetAllAuthors() ([]models.Author, error)
	CreateAuthor(details *models.Author) (string, error)
	UpdateAuthor(details *models.UpdateAuthor, id string) (string, error)
	DeleteAuthor(id string) (string, error)
}
