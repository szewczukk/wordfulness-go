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

func (r *Router) Get(endpoint string, handler http.HandlerFunc) {
	r.createEndPointSliceIfNotExists(endpoint)
	r.addEndPoint(endpoint, http.MethodGet, handler)
}

func (r *Router) Post(endpoint string, handler http.HandlerFunc) {
	r.createEndPointSliceIfNotExists(endpoint)
	r.addEndPoint(endpoint, http.MethodPost, handler)
}

func (r *Router) createEndPointSliceIfNotExists(endpoint string) {
	_, exists := r.routes[endpoint]
	if !exists {
		r.routes[endpoint] = make(map[string]http.HandlerFunc)
	}
}

func (r *Router) addEndPoint(endpoint string, method string, handler http.HandlerFunc) {
	r.routes[endpoint][method] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, exists := r.routes[req.URL.Path][req.Method]

	if !exists {
		http.NotFound(w, req)
		return
	}

	handler(w, req)
}
