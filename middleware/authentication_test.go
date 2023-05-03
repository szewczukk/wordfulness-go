package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"wordfulness/middleware"
	"wordfulness/storage"
	"wordfulness/types"
)

func AuthenticatedHelloHandler(w http.ResponseWriter, r *http.Request, u *types.User) {
	w.Write([]byte("Hello"))
}

func TestWithAuthentication(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	storage := storage.NewUserMemoryStorage([]*types.User{{Id: 0, Username: "jbytnar", Password: "zaq"}})
	req.AddCookie(makeAuthenticationCookie("jbytnar"))

	middleware.WithAuthentication(AuthenticatedHelloHandler, storage, "/login")(w, req)

	body := w.Body.String()

	if body != "Hello" {
		t.Errorf("Wrong body returned %v", body)
	}
}

func TestWithAuthenticationRedirectToLogin(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	storage := storage.NewUserMemoryStorage([]*types.User{{Id: 0, Username: "jbytnar", Password: "zaq"}})

	middleware.WithAuthentication(AuthenticatedHelloHandler, storage, "/login")(w, req)

	statusCode := w.Result().StatusCode
	url, _ := w.Result().Location()

	if statusCode != 301 {
		t.Errorf("Wrong status code returned, got: %v", statusCode)
	}

	if url.Path != "/login" {
		t.Errorf("Wrong redirection url, got: %v", statusCode)
	}
}

func TestWithAuthenticationUserNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	storage := storage.NewUserMemoryStorage([]*types.User{})
	req.AddCookie(makeAuthenticationCookie("jbytnar"))

	middleware.WithAuthentication(AuthenticatedHelloHandler, storage, "/login")(w, req)

	body := w.Body.String()
	statusCode := w.Result().StatusCode

	if statusCode != 500 {
		t.Errorf("Wrong status code returned, got: %v", statusCode)
	}

	if body != "not found\n" {
		t.Errorf("Wrong body returned %v", body)
	}
}

func makeAuthenticationCookie(userName string) *http.Cookie {
	return &http.Cookie{
		Name:     "authentication",
		Value:    "jbytnar",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
}
