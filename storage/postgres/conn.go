package postgres

import (
	"fmt"
	"os"

	"github.com/404th/bookstore/models"
	"github.com/jmoiron/sqlx"
)

var postgresConf = models.Config{
	Host:     "localhost",
	Port:     os.Getenv("POSTGRES_PORT"),
	User:     "postgres",
	Password: os.Getenv("POSTGRES_PASSWORD"),
	DBName:   "bookstore",
	SSLMode:  "disable",
}

func NewDB() (*sqlx.DB, error) {

	q := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", postgresConf.Host, postgresConf.Port, postgresConf.User, postgresConf.Password, postgresConf.DBName, postgresConf.SSLMode)

	db, err := sqlx.Connect("postgres", q)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func DBClose(db *sqlx.DB) error {
	return db.Close()
}
