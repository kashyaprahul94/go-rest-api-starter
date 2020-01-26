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

// New creates the new webserver
func New(port string) *Server {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	}).Methods(http.MethodGet)

	serverInstace := &Server{
		port:   port,
		router: router,
	}

	return serverInstace
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
