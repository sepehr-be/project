package routewrapper

import (
	"net/http"
)

type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRoutr() *Router {
	return &Router{
		routes: map[string]map[string]http.HandlerFunc{},
	}
}

func (r *Router) AddRoute(method string, path string, handler http.HandlerFunc) {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]http.HandlerFunc)
	}
	r.routes[path][method] = handler
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.AddRoute(http.MethodGet, path, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.AddRoute(http.MethodPost, path, handler)
}

func (r *Router) Put(path string, handler http.HandlerFunc) {
	r.AddRoute(http.MethodPut, path, handler)
}

func (r *Router) Delete(path string, handler http.HandlerFunc) {
	r.AddRoute(http.MethodDelete, path, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handlers, ok := r.routes[req.URL.Path]; ok {
		if handler, ok := handlers[req.Method]; ok {
			handler(w, req)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, req)
}
