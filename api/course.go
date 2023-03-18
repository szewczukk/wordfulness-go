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

func GetCourses(
	storage storage.IStorage,
	template *template.Template,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &GetCourseData{}

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
}
