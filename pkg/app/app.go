package app

import (
	"os"

	"github.com/kashyaprahul94/go-rest-api-starter/pkg/api/server"
)

// Start method starts the app
func Start() {

	port := os.Getenv("PORT")

	if port == "" {
		panic("no http port defined")
	}

	server := server.New(port)

	server.Listen()
}
