package main

import (
	"log"
	"net/http"
	"wordfulness/api"
	"wordfulness/storage"
)

func main() {
	storage := &storage.MemoryStorage{}

	http.HandleFunc("/", api.GetCourses(storage))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
