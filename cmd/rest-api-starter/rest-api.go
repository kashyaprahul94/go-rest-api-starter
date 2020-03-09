package main

import (
	"os"

	"github.com/kashyaprahul94/go-rest-api-starter/pkg/app"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		log.Error("error :", err)
		os.Exit(1)
	}
}

func run() error {

	app.Start()

	return nil
}
