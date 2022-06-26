package postgres

import (
	"log"

	"github.com/404th/bookstore/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgres struct {
	db               *sqlx.DB
	authorRepo       *authorPg
	bookCategoryRepo *bookCategoryPg
	bookRepo         *bookPg
}

func NewPostgres(str string) storage.StorageI {
	db, err := sqlx.Connect("postgres", str)
	if err != nil {
		log.Fatalf("Could not connect to db: %e", err)
		return nil
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not connect to db error ping: %e", err)
		return nil
	}

	return &postgres{
		db:               db,
		authorRepo:       &authorPg{db},
		bookRepo:         &bookPg{db},
		bookCategoryRepo: &bookCategoryPg{db},
	}
}

func (pg *postgres) CloseDB() error {
	return pg.db.Close()
}

func (pg *postgres) AuthorRepo() storage.AuthorI {
	return pg.authorRepo
}

func (pg *postgres) BookCategoryRepo() storage.BookCategoryI {
	return pg.bookCategoryRepo
}

func (pg *postgres) BookRepo() storage.BookI {
	return pg.bookRepo
}
