package server

import (
	"fmt"
	"net/http"
	"strings"

	router "github.com/julienschmidt/httprouter"
)

// RouteGroup is group of routes for sub-router implementation
type RouteGroup struct {
	r *Router
	g string
}

func newRouteGroup(r *Router, g string) *RouteGroup {
	if g[0] != '/' {
		panic("path must begin with '/' in path '" + g + "'")
	}

	//Strip traling / (if present) as all added sub paths must start with a /
	if g[len(g)-1] == '/' {
		g = g[:len(g)-1]
	}

	return &RouteGroup{r: r, g: g}
}

// NewGroup is used to create a new route group
func (r *RouteGroup) NewGroup(path string) *RouteGroup {
	return newRouteGroup(r.r, r.subPath(path))
}

// Handle is proxy method
func (r *RouteGroup) Handle(method, path string, handle router.Handle) {
	r.r.Handle(method, r.subPath(path), handle)
}

// Handler is proxy method
func (r *RouteGroup) Handler(method, path string, handler http.Handler) {
	r.r.Handler(method, r.subPath(path), handler)
}

// HandlerFunc is proxy method
func (r *RouteGroup) HandlerFunc(method, path string, handler http.HandlerFunc) {
	r.r.HandlerFunc(method, r.subPath(path), handler)
}

// GET is proxy method
func (r *RouteGroup) GET(path string, handle router.Handle) {
	r.Handle("GET", path, handle)
}

// HEAD is proxy method
func (r *RouteGroup) HEAD(path string, handle router.Handle) {
	r.Handle("HEAD", path, handle)
}

// OPTIONS is proxy method
func (r *RouteGroup) OPTIONS(path string, handle router.Handle) {
	r.Handle("OPTIONS", path, handle)
}

// POST is proxy method
func (r *RouteGroup) POST(path string, handle router.Handle) {
	r.Handle("POST", path, handle)
}

// PUT is proxy method
func (r *RouteGroup) PUT(path string, handle router.Handle) {
	r.Handle("PUT", path, handle)
}

// PATCH is proxy method
func (r *RouteGroup) PATCH(path string, handle router.Handle) {
	r.Handle("PATCH", path, handle)
}

// DELETE is proxy method
func (r *RouteGroup) DELETE(path string, handle router.Handle) {
	r.Handle("DELETE", path, handle)
}

func (r *RouteGroup) subPath(path string) string {
	if path[0] != '/' {
		panic("path must start with a '/'")
	}

	// the path is only '/' trim it
	if path == "/" {
		path = ""
	}

	fullPath := strings.Join([]string{r.g, path}, "")

	fmt.Println(fullPath)

	return fullPath
}

// NewGroup is used to create a new route group
func (r *Router) NewGroup(path string) *RouteGroup {
	return newRouteGroup(r, path)
}
