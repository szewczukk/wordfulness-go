package main

import (
	"html/template"
	"log"
	"net/http"
	"wordfulness/storage"
)

func homepage(w http.ResponseWriter, r *http.Request, storage storage.IStorage) {
	template := template.Must(template.ParseFiles("templates/homepage.html"))
	courses := storage.GetAllCourses()

	template.Execute(w, courses)
}

func createCourse(w http.ResponseWriter, r *http.Request, storage storage.IStorage) {
	r.ParseForm()

	storage.CreateCourse(r.Form.Get("name"))

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func withStorage(
	storage storage.IStorage,
	handler func(http.ResponseWriter, *http.Request, storage.IStorage),
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, storage)
	}
}

func main() {
	storage := &storage.MemoryStorage{}

	http.HandleFunc("/", withStorage(storage, homepage))
	http.HandleFunc("/create-course", withStorage(storage, createCourse))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
