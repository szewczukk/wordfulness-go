package services_test

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"wordfulness/services"
	"wordfulness/storage"
	"wordfulness/types"
)

func TestCreateUserGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/register", nil)
	w := httptest.NewRecorder()

	storage := storage.NewUserMemoryStorage([]*types.User{})
	temp, _ := template.New("homepage").Parse("render")
	templates := map[string]*template.Template{
		"CreateUser": temp,
	}
	service := services.NewUserService(storage, templates)

	service.CreateUserGet(w, req)

	body := w.Body.String()

	if body != "render" {
		t.Errorf("Wrong body returned %v", body)
	}
}

func TestCreateUserPost(t *testing.T) {
	form := url.Values{}
	form.Add("username", "jbytnar")
	form.Add("password", "zaq1@WSX")

	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	storage := storage.NewUserMemoryStorage([]*types.User{})
	templates := map[string]*template.Template{}
	service := services.NewUserService(storage, templates)

	service.CreateUserPost(w, req)

	statusCode := w.Result().StatusCode
	url, _ := w.Result().Location()
	user, _ := storage.GetUser(0)

	if statusCode != 301 {
		t.Errorf("Wrong status code returned, got: %v", statusCode)
	}

	if url.Path != "/" {
		t.Errorf("Wrong redirection url, got: %v", statusCode)
	}

	if user.Username != "jbytnar" {
		t.Errorf("Wrong username , got: %v", user.Username)
	}
}

func TestCreateUserDuplicate(t *testing.T) {
	form := url.Values{}
	form.Add("username", "jbytnar")
	form.Add("password", "zaq1@WSX")

	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	storage := storage.NewUserMemoryStorage([]*types.User{{Id: 0, Username: "jbytnar", Password: "zaq1@WSX"}})
	templates := map[string]*template.Template{}
	service := services.NewUserService(storage, templates)

	service.CreateUserPost(w, req)

	body := w.Body.String()
	statusCode := w.Result().StatusCode

	if statusCode != http.StatusBadRequest {
		t.Errorf("Wrong status code returned, got: %v", statusCode)
	}

	if body != "duplicate\n" {
		t.Errorf("Wrong body, got: %v", body)
	}
}
