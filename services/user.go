package services

import (
	"html/template"
	"net/http"
	"wordfulness/types"

	"golang.org/x/crypto/bcrypt"
)

type UserStorage interface {
	CreateUser(string, string) error
	GetUserByUserName(string) (*types.User, error)
}

type UserService struct {
	storage   UserStorage
	templates map[string]*template.Template
}

func NewUserService(
	storage UserStorage,
	templates map[string]*template.Template,
) *UserService {
	return &UserService{
		storage:   storage,
		templates: templates,
	}
}

func (s *UserService) CreateUserPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.storage.CreateUser(username, string(hashedPassword))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (s *UserService) CreateUserGet(w http.ResponseWriter, r *http.Request) {
	err := s.templates["CreateUser"].Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *UserService) LogInGet(w http.ResponseWriter, r *http.Request) {
	err := s.templates["LogIn"].Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *UserService) LogInPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := s.storage.GetUserByUserName(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookie := makeAuthenticationCookie(username)

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func makeAuthenticationCookie(username string) *http.Cookie {
	return &http.Cookie{
		Name:     "authentication",
		Value:    username,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
}
