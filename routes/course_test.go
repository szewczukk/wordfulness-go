package routes_test

import (
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

func TestHomePageGET(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{{Id: 0, Name: "German"}}, 1)
	template, _ := template.New("homepage").Parse("{{range .}}{{.Id}} {{.Name}}{{end}}")

	routes.HomePage(storage, template)(w, req)

	result := w.Body.String()
	if result != "0 German" {
		t.Errorf("Wrong body returned %v", result)
	}
}

func TestHomePagePOST(t *testing.T) {
	form := url.Values{}
	form.Add("name", "Spanish")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{{Id: 0, Name: "German"}}, 1)
	template, _ := template.New("homepage").Parse("{{range .}}{{.Id}} {{.Name}} {{end}}")

	routes.CreateCourse(storage, template)(w, req)

	result := w.Body.String()
	if result != "0 German 1 Spanish " {
		t.Errorf("Wrong body returned %v", result)
	}
}

func TestExistingDetailedCourse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/course?id=0", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	storage := storage.NewMemoryStorage([]*types.Course{{Id: 0, Name: "German"}}, 1)
	template, _ := template.New("homepage").Parse("{{.Id}} {{.Name}}")

	routes.DetailedCourse(storage, template)(w, req)

	result := w.Body.String()
	if result != "0 German" {
		t.Errorf("Wrong body returned %v", result)
	}
}
