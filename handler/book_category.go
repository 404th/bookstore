package handler

import (
	"net/http"
	"time"

	"github.com/404th/bookstore/helper"

	"github.com/404th/bookstore/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateBookCategory(ctx *gin.Context) {
	var bookCatCreate *models.CreateBookCategory
	var bookCat *models.BookCategory

	if err := ctx.ShouldBindJSON(&bookCatCreate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => book cat creating",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()

	bookCat.ID = helper.UUIDMaker()
	bookCat.CreatedAt = dt
	bookCat.UpdatedAt = dt
	bookCat.CategoryName = bookCatCreate.CategoryName

	res, err := h.strg.BookCategoryRepo().CreateBookCategory(bookCat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => book cat creating",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"response": models.Response{
			Error:   "",
			Message: "Book Cat created!",
			Data:    res,
		},
	})
}

func (h *handler) GetAllBookCategories(ctx *gin.Context) {
	bookCats, err := h.strg.BookCategoryRepo().GetAllBookCategories()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Some error has been caught in postgres:author getting all bookcategories",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": models.Response{
			Error:   "",
			Message: "Everything is good!",
			Data:    bookCats,
		},
	})
}

func (h *handler) GetBookCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := h.strg.BookCategoryRepo().GetBookCategory(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Some error has been caught in postgres:author getting a book cat",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": models.Response{
			Error:   "",
			Message: "Everything is good!",
			Data:    res,
		},
	})
}

func (h *handler) UpdateBookCategory(ctx *gin.Context) {
	var bookCatModel *models.UpdateBookCategory
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&bookCatModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => UpdateBookCategory",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()
	bookCatModel.UpdatedAt = dt

	res, err := h.strg.BookCategoryRepo().UpdateBookCategory(bookCatModel, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could get answer from pg => UpdateBookCategory",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": models.Response{
			Error:   "",
			Message: "Everything is good => UpdateBookCategory",
			Data:    res,
		},
	})
}

func (h *handler) DeleteBookCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := h.strg.BookCategoryRepo().DeleteBookCategory(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Some error has been caught in postgres:book cat deleting",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": models.Response{
			Error:   "",
			Message: "Everything is good!",
			Data:    res,
		},
	})
}
