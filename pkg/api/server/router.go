package server

import (
	"net/http"
	"strings"

	router "github.com/julienschmidt/httprouter"
)

// SubRouter is group of routes for sub-router implementation
type SubRouter struct {
	router *Router
	path   string
}

func newSubRouter(router *Router, path string) *SubRouter {
	if path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}

	//Strip traling / (if present) as all added sub paths must start with a /
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	return &SubRouter{router: router, path: path}
}

// NewSubRouter is used to create a new route group
func (r *SubRouter) NewSubRouter(path string) *SubRouter {
	return newSubRouter(r.router, r.subPath(path))
}

// Handle is proxy method
func (r *SubRouter) Handle(method, path string, handle router.Handle) {
	r.router.Handle(method, r.subPath(path), handle)
}

// Handler is proxy method
func (r *SubRouter) Handler(method, path string, handler http.Handler) {
	r.router.Handler(method, r.subPath(path), handler)
}

// HandlerFunc is proxy method
func (r *SubRouter) HandlerFunc(method, path string, handler http.HandlerFunc) {
	r.router.HandlerFunc(method, r.subPath(path), handler)
}

// GET is proxy method
func (r *SubRouter) GET(path string, handle router.Handle) {
	r.Handle("GET", path, handle)
}

// HEAD is proxy method
func (r *SubRouter) HEAD(path string, handle router.Handle) {
	r.Handle("HEAD", path, handle)
}

// OPTIONS is proxy method
func (r *SubRouter) OPTIONS(path string, handle router.Handle) {
	r.Handle("OPTIONS", path, handle)
}

// POST is proxy method
func (r *SubRouter) POST(path string, handle router.Handle) {
	r.Handle("POST", path, handle)
}

// PUT is proxy method
func (r *SubRouter) PUT(path string, handle router.Handle) {
	r.Handle("PUT", path, handle)
}

// PATCH is proxy method
func (r *SubRouter) PATCH(path string, handle router.Handle) {
	r.Handle("PATCH", path, handle)
}

// DELETE is proxy method
func (r *SubRouter) DELETE(path string, handle router.Handle) {
	r.Handle("DELETE", path, handle)
}

func (r *SubRouter) subPath(path string) string {
	if path[0] != '/' {
		panic("path must start with a '/'")
	}

	// the path is only '/' trim it
	if path == "/" {
		path = ""
	}

	fullPath := strings.Join([]string{r.path, path}, "")

	return fullPath
}

// NewSubRouter is used to create a new route group
func (r *Router) NewSubRouter(path string) *SubRouter {
	return newSubRouter(r, path)
}
