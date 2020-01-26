package app

import (
	"log"
	"os"

	"github.com/kashyaprahul94/go-rest-api-starter/pkg/api/health"
	"github.com/kashyaprahul94/go-rest-api-starter/pkg/api/server"
)

// Start method starts the app
func Start() {

	port := os.Getenv("PORT")

	if port == "" {
		panic("no http port defined")
	}

	// Create new instance of web server
	serverInstance := server.New(port)

	// Register all the modules against the webserver
	registerModules(serverInstance)

	// Start listening on the webserver
	serverInstance.Listen(func() {
		log.Printf("> Server stated ---> http://localhost:%v", port)
	})
}

func registerModules(s *server.Server) {

	// Health module
	health.New(s.GetSubRouter(health.BasePath))

}
