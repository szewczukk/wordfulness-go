package api

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"wordfulness/storage"
	"wordfulness/types"
)

type GetMultipleCoursesData struct {
	Courses []*types.Course
	Error   error
}

func GetMultipleCourses(storage storage.IStorage, template *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &GetMultipleCoursesData{}

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

type GetSingleCourseData struct {
	Course *types.Course
	Error  error
}

func GetSingleCourse(storage storage.IStorage, template *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &GetSingleCourseData{}
		id := r.URL.Query().Get("id")

		parsedId, err := strconv.ParseInt(id, 10, 32)

		if err != nil {
			log.Fatal(err)
		}

		data.Course, data.Error = storage.GetCourse(parsedId)

		template.Execute(w, data)
	}
}

func DeleteCourseData(storage storage.IStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		parsedId, err := strconv.ParseInt(id, 10, 32)

		if err != nil {
			log.Fatal(err)
		}

		err = storage.DeleteCourse(parsedId)

		if err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}
