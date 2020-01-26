package app

import (
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

	serverInstance := server.New(port)

	registerModules(serverInstance)

	serverInstance.Listen()
}

func registerModules(s *server.Server) {

	// Health
	health.New(s.GetSubRouter(health.BasePath))

}
