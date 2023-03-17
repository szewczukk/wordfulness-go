package api

import (
	"html/template"
	"net/http"
	"wordfulness/storage"
	"wordfulness/types"
)

type GetCourseData struct {
	Courses []types.Course
	Error   error
}

func GetCourse(w http.ResponseWriter, r *http.Request, storage storage.IStorage) {
	data := &GetCourseData{}
	template := template.Must(template.ParseFiles("templates/homepage.html"))

	if r.Method == "POST" {
		r.ParseForm()

		error := storage.CreateCourse(r.Form.Get("name"))
		if error != nil {
			data.Error = error
		}
	}

	data.Courses = storage.GetAllCourses()

	template.Execute(w, data)
}
