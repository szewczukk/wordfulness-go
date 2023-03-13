package main

import (
	"html/template"
	"log"
	"net/http"
	"wordfulness/storage"
	"wordfulness/utils"
)

func homepage(w http.ResponseWriter, r *http.Request, storage storage.IStorage) {
	template := template.Must(template.ParseFiles("templates/homepage.html"))
	if r.Method == "POST" {
		r.ParseForm()

		storage.CreateCourse(r.Form.Get("name"))
	}

	courses := storage.GetAllCourses()

	template.Execute(w, courses)
}

func main() {
	storage := &storage.MemoryStorage{}

	http.HandleFunc("/", utils.WithStorage(storage, homepage))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
