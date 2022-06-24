package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/404th/bookstore/helper"

	"github.com/404th/bookstore/models"
	"github.com/gin-gonic/gin"
)

func CreateBookCategory(ctx *gin.Context) {
	var bookCmodel *models.CreateBookCategory
	var bookC *models.BookCategory

	if err := ctx.ShouldBindJSON(&bookCmodel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => CreateBookCategory",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()

	bookC.ID = helper.UUIDMaker()
	bookC.CreatedAt = dt
	bookC.UpdatedAt = dt
	bookC.CategoryName = bookCmodel.CategoryName

	return
}

func GetAllBookCategory(ctx *gin.Context) {
	fmt.Println("Hello world!")
}

func GetBookCategory(ctx *gin.Context) {
	var bookCmodel *models.CreateBookCategory
	var bookC *models.BookCategory

	if err := ctx.ShouldBindJSON(&bookCmodel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => CreateBookCategory",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()

	bookC.ID = helper.UUIDMaker()
	bookC.CreatedAt = dt
	bookC.UpdatedAt = dt
	bookC.CategoryName = bookCmodel.CategoryName

	return
}

func UpdateBookCategory(ctx *gin.Context) {
	var bookCmodel *models.CreateBookCategory
	var bookC *models.BookCategory

	if err := ctx.ShouldBindJSON(&bookCmodel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => CreateBookCategory",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()

	bookC.ID = helper.UUIDMaker()
	bookC.CreatedAt = dt
	bookC.UpdatedAt = dt
	bookC.CategoryName = bookCmodel.CategoryName

	return
}

func DeleteBookCategory(ctx *gin.Context) {
	var bookCmodel *models.CreateBookCategory
	var bookC *models.BookCategory

	if err := ctx.ShouldBindJSON(&bookCmodel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => CreateBookCategory",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()

	bookC.ID = helper.UUIDMaker()
	bookC.CreatedAt = dt
	bookC.UpdatedAt = dt
	bookC.CategoryName = bookCmodel.CategoryName

	return
}
