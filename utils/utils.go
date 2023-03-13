package utils

import (
	"net/http"
	"wordfulness/storage"
)

func WithStorage(
	storage storage.IStorage,
	handler func(http.ResponseWriter, *http.Request, storage.IStorage),
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, storage)
	}
}
