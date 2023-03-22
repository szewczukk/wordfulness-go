package routes_test

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"wordfulness/routes"
	"wordfulness/types"
)

type MockHomePageStorage struct {
	courses []*types.Course
	nextId  int
}

func (s *MockHomePageStorage) GetAllCourses() ([]*types.Course, error) {
	return s.courses, nil
}

func (s *MockHomePageStorage) CreateCourse(name string) error {
	s.courses = append(s.courses, &types.Course{Id: s.nextId, Name: name})
	s.nextId++
	return nil
}

func TestHomePageGET(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	storage := &MockHomePageStorage{
		nextId:  1,
		courses: []*types.Course{{Id: 0, Name: "German 101"}},
	}
	template, _ := template.New("homepage").Parse("{{range .}}{{.Id}} {{.Name}}{{end}}")

	routes.HomePage(storage, template)(w, req)

	result := w.Body.String()
	if result != "0 German 101" {
		t.Errorf("Wrong body returned %v", result)
	}
}

func TestHomePagePOST(t *testing.T) {
	form := url.Values{}
	form.Add("name", "Spanish")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	storage := &MockHomePageStorage{
		nextId:  1,
		courses: []*types.Course{{Id: 0, Name: "German"}},
	}
	template, _ := template.New("homepage").Parse("{{range .}}{{.Id}} {{.Name}} {{end}}")

	routes.CreateCourse(storage, template)(w, req)

	result := w.Body.String()
	if result != "0 German 1 Spanish " {
		t.Errorf("Wrong body returned %v", result)
	}
}
