package server

import "github.com/gorilla/mux"

// HTTPServer is the instance for web server
type HTTPServer struct {

	// port on which web seever will listen to
	port string

	// router instance to provide routing
	router *mux.Router
}
