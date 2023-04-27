package core_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"wordfulness/core"
)

func TestGetCreatesNewEndpoint(t *testing.T) {
	router := core.NewRouter()

	router.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "get /")
	}))

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	body := w.Body.String()

	if body != "get /" {
		t.Errorf("Wrong body, got %v", body)
	}
}

func TestPostCreatesNewEndpoint(t *testing.T) {
	router := core.NewRouter()

	router.Post("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "post /")
	}))

	r := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	body := w.Body.String()

	if body != "post /" {
		t.Errorf("Wrong body, got %v", body)
	}
}

func TestNotExistPage(t *testing.T) {
	router := core.NewRouter()

	r := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	body := w.Body.String()

	if body != "404 page not found\n" {
		t.Errorf("Wrong body, got %v", body)
	}
}
