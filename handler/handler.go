package handler

import "github.com/404th/bookstore/storage"

type handler struct {
	strg storage.StorageI
}

func NewHandler(strg storage.StorageI) *handler {
	return &handler{
		strg: strg,
	}
}
