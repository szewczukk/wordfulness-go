package services_test

import (
	"errors"
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

type ErrorStorage struct{}

func (s *ErrorStorage) GetAllCourses() ([]*types.Course, error) {
	return nil, errors.New("error")
}

func (s *ErrorStorage) GetCourse(id int) (*types.Course, error) {
	return nil, errors.New("error")
}

func (s *ErrorStorage) CreateCourse(name string) error {
	return errors.New("error")
}

func (s *ErrorStorage) DeleteCourse(id int) error {
	return errors.New("error")
}

func TestHomePage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	temp, _ := template.New("homepage").Parse("{{range .}}{{.Id}} {{.Name}}{{end}}")
	templates := map[string]*template.Template{
		"HomePage": temp,
	}
	coursesController := services.NewCoursesService(storage, templates)

	coursesController.HomePage(w, req)

	body := w.Body.String()

	if body != "0 German" {
		t.Errorf("Wrong body returned %v", body)
	}
}

func TestErrorOnHomePage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	storage := &ErrorStorage{}
	temp, _ := template.New("homepage").Parse("{{range .}}{{.Id}} {{.Name}}{{end}}")
	templates := map[string]*template.Template{
		"HomePage": temp,
	}
	coursesController := services.NewCoursesService(storage, templates)

	coursesController.HomePage(w, req)

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
	templates := map[string]*template.Template{}
	coursesController := services.NewCoursesService(storage, templates)

	coursesController.CreateCourse(w, req)

	statusCode := w.Result().StatusCode
	url, _ := w.Result().Location()

	if statusCode != 308 {
		t.Errorf("Wrong status code returned, got: %v", statusCode)
	}

	if url.Path != "/" {
		t.Errorf("Wrong redirection url, got: %v", statusCode)
	}
}

func TestCreateCourseWithErrorStorage(t *testing.T) {
	form := url.Values{}
	form.Add("name", "Spanish")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	storage := &ErrorStorage{}
	templates := map[string]*template.Template{}
	coursesController := services.NewCoursesService(storage, templates)

	coursesController.CreateCourse(w, req)

	body := w.Body.String()
	statusCode := w.Result().StatusCode

	if body != "error\n" {
		t.Errorf("Wrong body returned %v", body)
	}

	if statusCode != 400 {
		t.Errorf("Wrong status code returned %v", statusCode)
	}
}

func TestExistingDetailedCourse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/course?id=0", nil)
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	temp, _ := template.New("homepage").Parse("{{.Id}} {{.Name}}")
	templates := map[string]*template.Template{
		"DetailedCourse": temp,
	}
	coursesController := services.NewCoursesService(storage, templates)

	coursesController.DetailedCourse(w, req)

	body := w.Body.String()

	if body != "0 German" {
		t.Errorf("Wrong body returned %v", body)
	}
}

func TestExistingDetailedCourseWithInvalidId(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/course?id=a", nil)
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	temp, _ := template.New("homepage").Parse("{{.Id}} {{.Name}}")
	templates := map[string]*template.Template{
		"DetailedCourse": temp,
	}
	coursesController := services.NewCoursesService(storage, templates)

	coursesController.DetailedCourse(w, req)

	body := w.Body.String()
	statusCode := w.Result().StatusCode

	if body != "strconv.ParseInt: parsing \"a\": invalid syntax\n" {
		t.Errorf("Wrong body returned %v", body)
	}

	if statusCode != 400 {
		t.Errorf("Wrong status code, got: %v", statusCode)
	}
}

func TestNonExistingDetailedCourse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/course?id=0", nil)
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{})
	temp, _ := template.New("homepage").Parse("{{.Id}} {{.Name}}")
	templates := map[string]*template.Template{
		"DetailedCourse": temp,
	}
	coursesController := services.NewCoursesService(storage, templates)

	coursesController.DetailedCourse(w, req)

	body := w.Body.String()
	statusCode := w.Result().StatusCode

	if body != "not found\n" {
		t.Errorf("Wrong body returned %v", body)
	}

	if statusCode != 404 {
		t.Errorf("Wrong status code, got: %v", statusCode)
	}
}

func TestDeleteExistingCourse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/delete-course?id=0", nil)
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	templates := map[string]*template.Template{}
	coursesController := services.NewCoursesService(storage, templates)

	coursesController.DeleteCourse(w, req)

	statusCode := w.Result().StatusCode
	url, _ := w.Result().Location()

	if statusCode != 308 {
		t.Errorf("Wrong status code returned, got: %v", statusCode)
	}

	if url.Path != "/" {
		t.Errorf("Wrong redirection url, got: %v", statusCode)
	}
}

func TestDeleteExistingCourseWithInvalidId(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/delete-course?id=a", nil)
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	templates := map[string]*template.Template{}
	coursesController := services.NewCoursesService(storage, templates)

	coursesController.DeleteCourse(w, req)

	statusCode := w.Result().StatusCode
	body := w.Body.String()

	if body != "strconv.ParseInt: parsing \"a\": invalid syntax\n" {
		t.Errorf("Wrong body, got: %v", body)
	}

	if statusCode != 400 {
		t.Errorf("Wrong status code, got: %v", statusCode)
	}
}

func TestDeleteNonExistingCourse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/delete-course?id=0", nil)
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{})
	templates := map[string]*template.Template{}
	coursesController := services.NewCoursesService(storage, templates)

	coursesController.DeleteCourse(w, req)

	statusCode := w.Result().StatusCode
	body := w.Body.String()

	if body != "not found\n" {
		t.Errorf("Wrong body, got: %v", body)
	}

	if statusCode != 404 {
		t.Errorf("Wrong status code, got: %v", statusCode)
	}
}
