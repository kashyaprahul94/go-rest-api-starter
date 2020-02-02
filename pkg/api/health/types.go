package health

import "github.com/gorilla/mux"

// Health is the model
type Health struct {

	// Status of the system
	Status string `json:"status"`

	// Version of the application
	Version string `json:"version"`
}

// Controller is kind of API Controller responsible for heatlh
type Controller struct {

	// router instance to handle all the incoming http request
	router *mux.Router

	// service which does all the necessary operations
	service *Service
}
