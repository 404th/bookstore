package main

import (
	"fmt"
	"log"

	"github.com/404th/bookstore/config"
	"github.com/404th/bookstore/handler"
	"github.com/404th/bookstore/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	str := fmt.Sprintf("port=%d host=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.PostgresPort, cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresDatabase, cfg.PostgresPassword, cfg.PostgresSSLMode,
	)

	strg := postgres.NewPostgres(str)
	defer strg.CloseDB()

	handler := handler.NewHandler(strg)

	switch cfg.Environment {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	ge := gin.Default()

	api := ge.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			book_category := v1.Group("/book_category")
			{
				book_category.POST("/", handler.CreateBookCategory)
				book_category.GET("/", handler.GetAllBookCategories)
				book_category.GET("/:id", handler.GetBookCategory)
				book_category.PUT("/:id", handler.UpdateBookCategory)
				book_category.DELETE("/:id", handler.DeleteBookCategory)
			}

			books := v1.Group("/books")
			{
				books.POST("/", handler.CreateBook)
				books.GET("/", handler.GetAllBooks)
				books.GET("/:id", handler.GetBook)
				books.PUT("/:id", handler.UpdateBook)
				books.DELETE("/:id", handler.DeleteBook)
			}

			authors := v1.Group("/authors")
			{
				authors.POST("/", handler.CreateAuthor)
				authors.GET("/", handler.GetAllAuthors)
				authors.GET("/:id", handler.GetAuthor)
				authors.PUT("/:id", handler.UpdateAuthor)
				authors.DELETE("/:id", handler.DeleteAuthor)
			}
		}
	}

	if err := ge.Run(cfg.HTTPPort); err != nil {
		log.Fatalf("error while running app on port: %s => %e", cfg.HTTPPort, err)
		return
	}
}
