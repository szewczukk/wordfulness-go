package core

import (
	"net/http"
)

type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
}

func (server *Router) Get(endpoint string, handler http.HandlerFunc) {
	_, exists := server.routes[endpoint]
	if !exists {
		server.routes[endpoint] = make(map[string]http.HandlerFunc)
	}

	server.routes[endpoint][http.MethodGet] = handler
}

func (server *Router) Post(endpoint string, handler http.HandlerFunc) {
	_, exists := server.routes[endpoint]
	if !exists {
		server.routes[endpoint] = make(map[string]http.HandlerFunc)
	}

	server.routes[endpoint][http.MethodPost] = handler
}

func (server *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, exists := server.routes[r.URL.Path][r.Method]

	if !exists {
		http.NotFound(w, r)
		return
	}

	handler(w, r)
}
