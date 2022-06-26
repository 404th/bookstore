package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/404th/bookstore/helper"
	"github.com/404th/bookstore/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateBook(ctx *gin.Context) {
	var bookCreate *models.CreateBook
	var book *models.Book

	if err := ctx.ShouldBindJSON(&bookCreate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => book creating",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()

	book.ID = helper.UUIDMaker()
	book.CreatedAt = dt
	book.UpdatedAt = dt
	book.CategoryID = bookCreate.CategoryID
	book.AuthorID = bookCreate.AuthorID
	book.Name = bookCreate.Name
	book.Price = bookCreate.Price
	book.Definition = bookCreate.Definition

	res, err := h.strg.BookRepo().CreateBook(book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => book creating",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"response": models.Response{
			Error:   "",
			Message: "Book created!",
			Data:    res,
		},
	})
}

func (h *handler) GetAllBooks(ctx *gin.Context) {
	var qP models.ApplicationQueryParamModel

	offset, offset_exists := ctx.GetQuery("offset")
	if offset_exists {
		res_offset, err := strconv.Atoi(offset)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"response": models.Response{
					Error:   err.Error(),
					Message: "Some error has been caught in postgres:author getting all books",
					Data:    nil,
				},
			})
			return
		}

		qP.Offset = res_offset
	}

	limit, limit_exists := ctx.GetQuery("limit")
	if limit_exists {
		res_limit, err := strconv.Atoi(limit)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"response": models.Response{
					Error:   err.Error(),
					Message: "Some error has been caught in postgres:author getting all books",
					Data:    nil,
				},
			})
			return
		}

		qP.Limit = res_limit
	}

	search, search_exists := ctx.GetQuery("search")
	if search_exists {
		qP.Search = search
	}

	books, err := h.strg.BookRepo().GetAllBooks(qP)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Some error has been caught in postgres:author getting all books",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": models.Response{
			Error:   "",
			Message: "Everything is good!",
			Data:    books,
		},
	})
}

func (h *handler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := h.strg.BookRepo().GetBook(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Some error has been caught in postgres:author getting a book",
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

func (h *handler) UpdateBook(ctx *gin.Context) {
	var bookModel *models.UpdateBook
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&bookModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not bind json => UpdateBook",
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()
	bookModel.UpdatedAt = dt

	res, err := h.strg.BookRepo().UpdateBook(bookModel, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could get answer from pg => UpdateBook",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": models.Response{
			Error:   "",
			Message: "Everything is good => UpdateBook",
			Data:    res,
		},
	})
}

func (h *handler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := h.strg.BookRepo().DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Some error has been caught in postgres:book deleting",
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
