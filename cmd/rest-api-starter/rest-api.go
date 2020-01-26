package main

import (
	"log"
	"os"

	"github.com/kashyaprahul94/go-rest-api-starter/pkg/app"
)

func main() {
	if err := run(); err != nil {
		log.Println("error :", err)
		os.Exit(1)
	}
}

func run() error {

	app.Start()

	return nil
}
