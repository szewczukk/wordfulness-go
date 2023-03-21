package routes

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"wordfulness/types"
)

type HomePageStorage interface {
	GetAllCourses() ([]*types.Course, error)
}

func HomePage(storage HomePageStorage, template *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		courses, err := storage.GetAllCourses()
		if err != nil {
			log.Fatal(err)
		}

		template.Execute(w, courses)
	})
}

type CreateCourseStorage interface {
	CreateCourse(string) error
	GetAllCourses() ([]*types.Course, error)
}

func CreateCourse(storage CreateCourseStorage, template *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		name := r.Form.Get("name")

		err := storage.CreateCourse(name)
		if err != nil {
			log.Fatal(err)
		}

		courses, err := storage.GetAllCourses()
		if err != nil {
			log.Fatal(err)
		}

		template.Execute(w, courses)
	})
}

type DetailedCourseStorage interface {
	GetCourse(int64) (*types.Course, error)
}

func DetailedCourse(storage DetailedCourseStorage, template *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
}

type DeleteCourseStorage interface {
	DeleteCourse(int64) error
}

func DeleteCourse(storage DeleteCourseStorage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
}
