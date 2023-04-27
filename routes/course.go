package routes

import (
	"html/template"
	"net/http"
	"strconv"
	"wordfulness/types"
)

type CoursesStorage interface {
	GetAllCourses() ([]*types.Course, error)
	GetCourse(int) (*types.Course, error)
	CreateCourse(string) error
	DeleteCourse(int) error
}

type CoursesController struct {
	storage   CoursesStorage
	templates map[string]*template.Template
}

func NewCoursesController(
	storage CoursesStorage,
	templates map[string]*template.Template,
) *CoursesController {
	return &CoursesController{
		storage:   storage,
		templates: templates,
	}
}

func (c *CoursesController) HomePage(w http.ResponseWriter, r *http.Request) {
	courses, err := c.storage.GetAllCourses()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	c.templates["HomePage"].Execute(w, courses)
}

func (c *CoursesController) CreateCourse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.Form.Get("name")

	err := c.storage.CreateCourse(name)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func (c *CoursesController) DetailedCourse(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	parsedId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	course, err := c.storage.GetCourse(int(parsedId))
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	c.templates["DetailedCourse"].Execute(w, course)
}

func (c *CoursesController) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	parsedId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = c.storage.DeleteCourse(int(parsedId))
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

type HomePageStorage interface {
	GetAllCourses() ([]*types.Course, error)
}

func HomePage(storage HomePageStorage, template *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		courses, err := storage.GetAllCourses()
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
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
			http.Error(w, err.Error(), 400)
			return
		}

		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	})
}

type DetailedCourseStorage interface {
	GetCourse(int) (*types.Course, error)
}

func DetailedCourse(storage DetailedCourseStorage, template *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		parsedId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		course, err := storage.GetCourse(int(parsedId))
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		template.Execute(w, course)
	})
}

type DeleteCourseStorage interface {
	DeleteCourse(int) error
}

func DeleteCourse(storage DeleteCourseStorage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		parsedId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = storage.DeleteCourse(int(parsedId))
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	})
}
