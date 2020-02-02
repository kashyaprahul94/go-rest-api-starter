package server

import (
	router "github.com/julienschmidt/httprouter"
)

// HTTPServer is the instance for web server
type HTTPServer struct {

	// port on which web seever will listen to
	port string

	// router instance to provide routing
	router *Router
}

// Router is the wrapper on top of julienschmidt/httprouter
type Router struct {
	*router.Router
}
