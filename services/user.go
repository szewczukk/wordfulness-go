package services

import (
	"html/template"
	"net/http"
)

type UserStorage interface {
	CreateUser(string, string) error
}

type UserService struct {
	storage   UserStorage
	templates map[string]*template.Template
}

func NewUserService(storage UserStorage, templates map[string]*template.Template) *UserService {
	return &UserService{
		storage:   storage,
		templates: templates,
	}
}

func (s *UserService) CreateUserPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	err := s.storage.CreateUser(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (s *UserService) CreateUserGet(w http.ResponseWriter, r *http.Request) {
	s.templates["CreateUser"].Execute(w, nil)
}