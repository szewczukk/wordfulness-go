package main

import (
	"html/template"
	"log"
	"net/http"
	"wordfulness/storage"
	"wordfulness/types"
	"wordfulness/utils"
)

type HomepageData struct {
	Courses []types.Course
	Error   error
}

func homepage(w http.ResponseWriter, r *http.Request, storage storage.IStorage) {
	pagedata := &HomepageData{}
	template := template.Must(template.ParseFiles("templates/homepage.html"))

	if r.Method == "POST" {
		r.ParseForm()

		error := storage.CreateCourse(r.Form.Get("name"))
		if error != nil {
			pagedata.Error = error
		}
	}

	pagedata.Courses = storage.GetAllCourses()

	template.Execute(w, pagedata)
}

func main() {
	storage := &storage.MemoryStorage{}

	http.HandleFunc("/", utils.WithStorage(storage, homepage))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
