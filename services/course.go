package services

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"wordfulness/middleware"
	"wordfulness/types"
)

type CoursesStorage interface {
	GetAllCourses() ([]*types.Course, error)
	GetCourse(int) (*types.Course, error)
	CreateCourse(string) error
	UpdateCourse(int, string) error
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

	user, ok := middleware.UserFromContext(r.Context())
	if !ok {
		http.Error(w, "User not present in the context", http.StatusInternalServerError)
		return
	}

	payload := makeHomePagePayload(user, courses)

	s.templates["HomePage"].Execute(w, payload)
}

func (s *CoursesService) CreateCourse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.Form.Get("name")

	err := s.storage.CreateCourse(name)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
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

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (s *CoursesService) UpdateCourseGET(w http.ResponseWriter, r *http.Request) {
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

	s.templates["UpdateCourse"].Execute(w, course)
}

func (s *CoursesService) UpdateCoursePOST(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := r.Form.Get("id")
	name := r.Form.Get("name")

	parsedId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = s.storage.UpdateCourse(int(parsedId), name)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	redirectUrl := fmt.Sprint("/courses?id=", id)

	http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
}

type HomePagePayload struct {
	UserName string
	Courses  []*types.Course
}

func makeHomePagePayload(
	user *types.User,
	courses []*types.Course,
) HomePagePayload {
	return HomePagePayload{
		UserName: user.Username,
		Courses:  courses,
	}
}
