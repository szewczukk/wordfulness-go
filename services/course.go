package services

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

type CoursesService struct {
	storage   CoursesStorage
	templates map[string]*template.Template
}

func NewCoursesService(
	storage CoursesStorage,
	templates map[string]*template.Template,
) *CoursesService {
	return &CoursesService{
		storage:   storage,
		templates: templates,
	}
}

func (s *CoursesService) HomePage(w http.ResponseWriter, r *http.Request) {
	courses, err := s.storage.GetAllCourses()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	s.templates["HomePage"].Execute(w, courses)
}

func (s *CoursesService) CreateCourse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.Form.Get("name")

	err := s.storage.CreateCourse(name)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func (s *CoursesService) DetailedCourse(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	parsedId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	course, err := s.storage.GetCourse(int(parsedId))
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	s.templates["DetailedCourse"].Execute(w, course)
}

func (s *CoursesService) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	parsedId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = s.storage.DeleteCourse(int(parsedId))
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
