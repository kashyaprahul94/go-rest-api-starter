package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Server is the instance for web server
type Server struct {

	// port on which web seever will listen to
	port string

	// router instance to provide routing
	router *mux.Router
}

// New creates the instance of new http web server \n
// It accepts a port as an argument on which http web server will listen on
func New(port string) *Server {

	router := mux.NewRouter()

	serverInstace := &Server{
		port:   port,
		router: router,
	}

	return serverInstace
}

// GetSubRouter will create an empty route and return the subrouter instance
func (s *Server) GetSubRouter(forPath string) *mux.Router {

	subRouter := s.router.PathPrefix(forPath).Subrouter()

	return subRouter
}

// Listen will start listening
func (s *Server) Listen() {

	port := s.port

	log.Printf("> Server stated ---> http://localhost:%v", port)

	err := http.ListenAndServe(strings.Join([]string{"", port}, ":"), s.router)

	if err != nil {
		log.Println("web server encountered error")
		log.Fatalln(err)
	}
}
