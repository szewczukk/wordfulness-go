package main

import (
	"html/template"
	"log"
	"net/http"
	"wordfulness/api"
	"wordfulness/storage"
)

func main() {
	storage := &storage.MemoryStorage{}
	tmpl := template.Must(template.ParseFiles("templates/homepage.html"))

	http.HandleFunc("/", api.GetCourses(storage, tmpl))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
