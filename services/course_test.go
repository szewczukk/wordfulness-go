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

func (s *ErrorStorage) UpdateCourse(id int, name string) error {
	return errors.New("not found")
}

func TestHomePage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	temp, _ := template.New("homepage").Parse(
		"{{.UserName}} {{range .Courses}}{{.Id}} {{.Name}}{{end}}",
	)
	templates := map[string]*template.Template{
		"HomePage": temp,
	}
	service := services.NewCoursesService(storage, templates)

	service.HomePage(w, req, &types.User{Id: 0, Username: "jbytnar", Password: ""})

	body := w.Body.String()

	if body != "jbytnar 0 German" {
		t.Errorf("Wrong body returned %v", body)
	}
}

func TestHomePageError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	storage := &ErrorStorage{}
	temp, _ := template.New("homepage").Parse(
		"{{.UserName}} {{range .Courses}}{{.Id}} {{.Name}}{{end}}",
	)
	templates := map[string]*template.Template{
		"HomePage": temp,
	}
	service := services.NewCoursesService(storage, templates)

	service.HomePage(w, req, &types.User{Id: 0, Username: "jbytnar", Password: ""})

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

	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	templates := map[string]*template.Template{}
	service := services.NewCoursesService(storage, templates)

	service.CreateCourse(w, req)

	statusCode := w.Result().StatusCode
	url, _ := w.Result().Location()

	if statusCode != 301 {
		t.Errorf("Wrong status code returned, got: %v", statusCode)
	}

	if url.Path != "/" {
		t.Errorf("Wrong redirection url, got: %v", statusCode)
	}
}

func TestCreateCourseError(t *testing.T) {
	form := url.Values{}
	form.Add("name", "Spanish")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	storage := &ErrorStorage{}
	templates := map[string]*template.Template{}
	service := services.NewCoursesService(storage, templates)

	service.CreateCourse(w, req)

	body := w.Body.String()
	statusCode := w.Result().StatusCode

	if body != "error\n" {
		t.Errorf("Wrong body returned %v", body)
	}

	if statusCode != 400 {
		t.Errorf("Wrong status code returned %v", statusCode)
	}
}

func TestDetailedCourse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/course?id=0", nil)
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	temp, _ := template.New("homepage").Parse("{{.Id}} {{.Name}}")
	templates := map[string]*template.Template{
		"DetailedCourse": temp,
	}
	service := services.NewCoursesService(storage, templates)

	service.DetailedCourse(w, req)

	body := w.Body.String()

	if body != "0 German" {
		t.Errorf("Wrong body returned %v", body)
	}
}

func TestDetailedCourseInvalidId(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/course?id=a", nil)
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	temp, _ := template.New("homepage").Parse("{{.Id}} {{.Name}}")
	templates := map[string]*template.Template{
		"DetailedCourse": temp,
	}
	service := services.NewCoursesService(storage, templates)

	service.DetailedCourse(w, req)

	body := w.Body.String()
	statusCode := w.Result().StatusCode

	if body != "strconv.ParseInt: parsing \"a\": invalid syntax\n" {
		t.Errorf("Wrong body returned %v", body)
	}

	if statusCode != 400 {
		t.Errorf("Wrong status code, got: %v", statusCode)
	}
}

func TestDetailedCourseNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/course?id=0", nil)
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{})
	temp, _ := template.New("homepage").Parse("{{.Id}} {{.Name}}")
	templates := map[string]*template.Template{
		"DetailedCourse": temp,
	}
	service := services.NewCoursesService(storage, templates)

	service.DetailedCourse(w, req)

	body := w.Body.String()
	statusCode := w.Result().StatusCode

	if body != "not found\n" {
		t.Errorf("Wrong body returned %v", body)
	}

	if statusCode != 404 {
		t.Errorf("Wrong status code, got: %v", statusCode)
	}
}

func TestDeleteCourse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/delete-course?id=0", nil)
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	templates := map[string]*template.Template{}
	service := services.NewCoursesService(storage, templates)

	service.DeleteCourse(w, req)

	statusCode := w.Result().StatusCode
	url, _ := w.Result().Location()

	if statusCode != 301 {
		t.Errorf("Wrong status code returned, got: %v", statusCode)
	}

	if url.Path != "/" {
		t.Errorf("Wrong redirection url, got: %v", url.Path)
	}
}

func TestDeleteCourseInvalidId(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/delete-course?id=a", nil)
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	templates := map[string]*template.Template{}
	service := services.NewCoursesService(storage, templates)

	service.DeleteCourse(w, req)

	statusCode := w.Result().StatusCode
	body := w.Body.String()

	if body != "strconv.ParseInt: parsing \"a\": invalid syntax\n" {
		t.Errorf("Wrong body, got: %v", body)
	}

	if statusCode != 400 {
		t.Errorf("Wrong status code, got: %v", statusCode)
	}
}

func TestDeleteCourseNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/delete-course?id=0", nil)
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{})
	templates := map[string]*template.Template{}
	service := services.NewCoursesService(storage, templates)

	service.DeleteCourse(w, req)

	statusCode := w.Result().StatusCode
	body := w.Body.String()

	if body != "not found\n" {
		t.Errorf("Wrong body, got: %v", body)
	}

	if statusCode != 404 {
		t.Errorf("Wrong status code, got: %v", statusCode)
	}
}

func TestUpdateCourseGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/update-course?id=0", nil)
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "Spanish"}})
	temp, _ := template.New("homepage").Parse("{{.Id}} {{.Name}}")
	templates := map[string]*template.Template{
		"UpdateCourse": temp,
	}
	service := services.NewCoursesService(storage, templates)

	service.UpdateCourseGET(w, req)

	body := w.Body.String()

	if body != "0 Spanish" {
		t.Errorf("Wrong body, got: %v", body)
	}
}

func TestUpdateCouseGetNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/update-course?id=0", nil)
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{})
	temp, _ := template.New("homepage").Parse("{{.Id}} {{.Name}}")
	templates := map[string]*template.Template{
		"DetailedCourse": temp,
	}
	service := services.NewCoursesService(storage, templates)

	service.UpdateCourseGET(w, req)

	body := w.Body.String()

	if body != "not found\n" {
		t.Errorf("Wrong body, got: %v", body)
	}
}

func TestUpdateCouseGetInvalidId(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/update-course?id=a", nil)
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{})
	temp, _ := template.New("homepage").Parse("{{.Id}} {{.Name}}")
	templates := map[string]*template.Template{
		"DetailedCourse": temp,
	}
	service := services.NewCoursesService(storage, templates)

	service.UpdateCourseGET(w, req)

	body := w.Body.String()
	statusCode := w.Result().StatusCode

	if body != "strconv.ParseInt: parsing \"a\": invalid syntax\n" {
		t.Errorf("Wrong body, got: %v", body)
	}

	if statusCode != 400 {
		t.Errorf("Wrong status code, got: %v", statusCode)
	}
}

func TestUpdateCoursePost(t *testing.T) {
	form := url.Values{}
	form.Add("id", "0")
	form.Add("name", "Spanish")

	req := httptest.NewRequest(
		http.MethodPost,
		"/update-course",
		strings.NewReader(form.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	templates := map[string]*template.Template{}
	service := services.NewCoursesService(storage, templates)

	service.UpdateCoursePOST(w, req)

	statusCode := w.Result().StatusCode
	if statusCode != 301 {
		t.Errorf("Wrong status code returned, got: %v", statusCode)
	}

	course, err := storage.GetCourse(0)
	if err != nil {
		t.Errorf("Error occurred, got: %v", course.Name)
	}

	if course.Name != "Spanish" {
		t.Errorf("Wrong name, got: %v", course.Name)
	}

	url, _ := w.Result().Location()
	if url.Path != "/courses" {
		t.Errorf("Wrong redirection url, got: %v", url.Path)
	}

	id := url.Query().Get("id")
	if id != "0" {
		t.Errorf("Wrong id in the query, got: %v", id)
	}
}

func TestUpdateCoursePostInvalidId(t *testing.T) {
	form := url.Values{}
	form.Add("id", "a")
	form.Add("name", "Spanish")

	req := httptest.NewRequest(
		http.MethodPost,
		"/update-course",
		strings.NewReader(form.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "German"}})
	templates := map[string]*template.Template{}
	service := services.NewCoursesService(storage, templates)

	service.UpdateCoursePOST(w, req)

	body := w.Body.String()
	statusCode := w.Result().StatusCode
	url, _ := w.Result().Location()

	if statusCode != 400 {
		t.Errorf("Wrong status code returned, got: %v", statusCode)
	}

	if body != "strconv.ParseInt: parsing \"a\": invalid syntax\n" {
		t.Errorf("Wrong redirection url, got: %v", url.Path)
	}
}

func TestUpdateCoursePostNotFound(t *testing.T) {
	form := url.Values{}
	form.Add("id", "0")
	form.Add("name", "Spanish")

	req := httptest.NewRequest(
		http.MethodPost,
		"/update-course",
		strings.NewReader(form.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{})
	templates := map[string]*template.Template{}
	service := services.NewCoursesService(storage, templates)

	service.UpdateCoursePOST(w, req)

	body := w.Body.String()
	statusCode := w.Result().StatusCode
	url, _ := w.Result().Location()

	if statusCode != 404 {
		t.Errorf("Wrong status code returned, got: %v", statusCode)
	}

	if body != "not found\n" {
		t.Errorf("Wrong redirection url, got: %v", url.Path)
	}
}

func TestUpdateCoursePostDuplicate(t *testing.T) {
	form := url.Values{}
	form.Add("id", "1")
	form.Add("name", "Spanish")

	req := httptest.NewRequest(
		http.MethodPost,
		"/update-course",
		strings.NewReader(form.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "Spanish"}, {Id: 1, Name: "German"}})
	templates := map[string]*template.Template{}
	service := services.NewCoursesService(storage, templates)

	service.UpdateCoursePOST(w, req)

	body := w.Body.String()
	statusCode := w.Result().StatusCode
	url, _ := w.Result().Location()

	if statusCode != 404 {
		t.Errorf("Wrong status code returned, got: %v", statusCode)
	}

	if body != "duplicate\n" {
		t.Errorf("Wrong redirection url, got: %v", url.Path)
	}
}
