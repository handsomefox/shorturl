package handlers

import "shorturl/pkg/storage"

type handler struct {
	storage storage.LinkStorage
}

func New(storage storage.LinkStorage) *handler {
	return &handler{storage: storage}
}
