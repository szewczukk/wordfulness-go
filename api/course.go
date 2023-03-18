package api

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"wordfulness/storage"
	"wordfulness/types"
)

type GetCoursesData struct {
	Courses []*types.Course
	Error   error
}

func GetCourses(
	storage storage.IStorage,
	template *template.Template,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &GetCoursesData{}

		if r.Method == "POST" {
			r.ParseForm()

			error := storage.CreateCourse(r.Form.Get("name"))
			if error != nil {
				data.Error = error
			}
		}

		data.Courses, data.Error = storage.GetAllCourses()

		template.Execute(w, data)
	}
}

type GetCourseData struct {
	Course *types.Course
	Error  error
}

func GetCourse(
	storage storage.IStorage,
	template *template.Template,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &GetCourseData{}
		id := r.URL.Query().Get("id")

		parsedId, err := strconv.ParseInt(id, 10, 16)

		if err != nil {
			log.Fatal(err)
		}

		data.Course, data.Error = storage.GetCourse(int(parsedId))

		template.Execute(w, data)
	}
}
