package main

import (
	"log"

	"github.com/404th/bookstore/handler"
	"github.com/404th/bookstore/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {

	// db
	db, err := postgres.NewDB()
	if err != nil {
		log.Fatalf("Error while connecting to db: %e", err)
		return
	}

	ge := gin.Default()

	api := ge.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			book_category := v1.Group("/book_category")
			{
				book_category.POST("/", handler.CreateBookCategory)
				book_category.GET("/", handler.GetAllBookCategory)
				book_category.GET("/:id", handler.GetBookCategory)
				book_category.PUT("/:id", handler.UpdateBookCategory)
				book_category.DELETE("/:id", handler.DeleteBookCategory)
			}

			books := v1.Group("/books")
			{
				books.POST("/", handler.CreateBook)
				books.GET("/", handler.GetAllBook)
				books.GET("/:id", handler.GetBook)
				books.PUT("/:id", handler.UpdateBook)
				books.DELETE("/:id", handler.DeleteBook)
			}
		}
	}

	if err := ge.Run(":6767"); err != nil {
		log.Fatalf("error while running app on port:6767 => %r", err)
		return
	}
}
