package handler

import (
	"net/http"
	"time"

	"github.com/404th/bookstore/helper"
	"github.com/404th/bookstore/models"
	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {
	var bookModel *models.CreateBook
	var book *models.Book

	if err := ctx.ShouldBindJSON(&bookModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => Book",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()

	book.ID = helper.UUIDMaker()
	book.CreatedAt = dt
	book.UpdatedAt = dt
	book.Name = bookModel.Name

}

func GetAllBook(ctx *gin.Context) {
	var bookModel *models.CreateBook
	var book *models.Book

	if err := ctx.ShouldBindJSON(&bookModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => CreateBook",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()

	book.ID = helper.UUIDMaker()
	book.CreatedAt = dt
	book.UpdatedAt = dt
	book.Name = bookModel.Name

	return
}

func GetBook(ctx *gin.Context) {
	var bookModel *models.CreateBook
	var book *models.Book

	if err := ctx.ShouldBindJSON(&bookModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => CreateBook",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()

	book.ID = helper.UUIDMaker()
	book.CreatedAt = dt
	book.UpdatedAt = dt
	book.Name = bookModel.Name

	return
}

func UpdateBook(ctx *gin.Context) {
	var bookModel *models.CreateBook
	var book *models.Book

	if err := ctx.ShouldBindJSON(&bookModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => CreateBook",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()

	book.ID = helper.UUIDMaker()
	book.CreatedAt = dt
	book.UpdatedAt = dt
	book.Name = bookModel.Name

	return
}

func DeleteBook(ctx *gin.Context) {
	var bookModel *models.CreateBook
	var book *models.Book

	if err := ctx.ShouldBindJSON(&bookModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => CreateBook",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()

	book.ID = helper.UUIDMaker()
	book.CreatedAt = dt
	book.UpdatedAt = dt
	book.Name = bookModel.Name

	return
}
