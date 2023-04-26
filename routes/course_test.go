package routes_test

import (
	"errors"
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"wordfulness/routes"
	"wordfulness/storage"
	"wordfulness/types"
)

type ErrorStorage struct{}

func (s *ErrorStorage) GetAllCourses() ([]*types.Course, error) {
	return nil, errors.New("error")
}

func TestHomePage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	template, _ := template.New("homepage").Parse("{{range .}}{{.Id}} {{.Name}}{{end}}")

	routes.HomePage(storage, template)(w, req)

	body := w.Body.String()

	if body != "0 German" {
		t.Errorf("Wrong body returned %v", body)
	}
}

func TestErrorOnHomePage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	storage := &ErrorStorage{}
	template, _ := template.New("homepage").Parse("{{range .}}{{.Id}} {{.Name}}{{end}}")

	routes.HomePage(storage, template)(w, req)

	statusCode := w.Result().StatusCode
	body := w.Body.String()

	if body != "error\n" {
		t.Errorf("Wrong body returned %v", body)
	}

	if statusCode != 400 {
		t.Errorf("Wrong status code returned %v", statusCode)
	}
}

func TestCreateCourse(t *testing.T) {
	form := url.Values{}
	form.Add("name", "Spanish")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	template, _ := template.New("homepage").Parse("{{range .}}{{.Id}} {{.Name}} {{end}}")

	routes.CreateCourse(storage, template)(w, req)

	body := w.Body.String()

	if body != "0 German 1 Spanish " {
		t.Errorf("Wrong body returned %v", body)
	}
}

func TestExistingDetailedCourse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/course?id=0", nil)
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	template, _ := template.New("homepage").Parse("{{.Id}} {{.Name}}")

	routes.DetailedCourse(storage, template)(w, req)

	body := w.Body.String()

	if body != "0 German" {
		t.Errorf("Wrong body returned %v", body)
	}
}

func TestDeleteExistingCourse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/delete-course?id=0", nil)
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})

	routes.DeleteCourse(storage)(w, req)

	body := w.Result().StatusCode

	if body != 308 {
		t.Errorf("Wrong status code returned, got: %v", body)
	}
}
