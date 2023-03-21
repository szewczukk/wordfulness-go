package api

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"wordfulness/types"
)

type HomePageStorage interface {
	CreateCourse(string) error
	GetAllCourses() ([]*types.Course, error)
}

func HomePage(storage HomePageStorage, template *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			err := storage.CreateCourse(r.Form.Get("name"))
			if err != nil {
				log.Fatal(err)
			}
		}

		courses, err := storage.GetAllCourses()

		if err != nil {
			log.Fatal(err)
		}

		template.Execute(w, courses)
	}
}

type DetailedCourseStorage interface {
	GetCourse(int64) (*types.Course, error)
}

func DetailedCourse(storage DetailedCourseStorage, template *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		parsedId, err := strconv.ParseInt(id, 10, 32)

		if err != nil {
			log.Fatal(err)
		}

		course, err := storage.GetCourse(parsedId)

		if err != nil {
			log.Fatal(err)
		}

		template.Execute(w, course)
	}
}

type DeleteCourseStorage interface {
	DeleteCourse(int64) error
}

func DeleteCourse(storage DeleteCourseStorage) http.HandlerFunc {
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
