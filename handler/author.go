package handler

import (
	"net/http"
	"time"

	"github.com/404th/bookstore/helper"
	"github.com/404th/bookstore/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateAuthor(ctx *gin.Context) {
	var ar *models.CreateAuthor
	var new_ar *models.Author

	if err := ctx.ShouldBindJSON(&ar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not convert data to json",
				Data:    nil,
			},
		})
		return
	}

	new_id := helper.UUIDMaker()
	dt := time.Now()

	new_ar.ID = new_id
	new_ar.Age = ar.Age
	new_ar.Email = ar.Email
	new_ar.Firstname = ar.Firstname
	new_ar.Secondname = ar.Secondname
	new_ar.CreatedAt = dt
	new_ar.UpdatedAt = dt

	res, err := h.strg.AuthorRepo().CreateAuthor(new_ar)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Some error has been caught in postgres:author",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"response": models.Response{
			Error:   "",
			Message: "Everything is good!",
			Data:    res,
		},
	})
}

func (h *handler) GetAllAuthors(ctx *gin.Context) {
	res, err := h.strg.AuthorRepo().GetAllAuthors()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Some error has been caught in postgres:author getting all authors",
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

func (h *handler) GetAuthor(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := h.strg.AuthorRepo().GetAuthor(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Some error has been caught in postgres:author getting an author",
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

func (h *handler) UpdateAuthor(ctx *gin.Context) {
	var ar *models.UpdateAuthor
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&ar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Could not convert data into json",
				Data:    nil,
			},
		})
		return
	}

	res, err := h.strg.AuthorRepo().UpdateAuthor(ar, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Some error has been caught in postgres:author update",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"response": models.Response{
			Error:   "",
			Message: "Everything is good!",
			Data:    res,
		},
	})
	return
}

func (h *handler) DeleteAuthor(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := h.strg.AuthorRepo().DeleteAuthor(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Some error has been caught in postgres:author deleting an author",
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
