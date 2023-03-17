package main

import (
	"log"
	"net/http"
	"wordfulness/api"
	"wordfulness/middleware"
	"wordfulness/storage"
)

func main() {
	storage := &storage.MemoryStorage{}

	http.HandleFunc("/", middleware.UseStorage(storage, api.GetCourse))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
