package health

import (
	"github.com/kashyaprahul94/go-rest-api-starter/pkg/api/server"
)

// Module is responsible for health management
type Module struct {

	// controller is the route manager
	controller *Controller
}

// New creates new instance of HealthModule
func New(router *server.SubRouter) *Module {

	// create new instance of service layer
	service := NewHealthService()

	// create new instance of controller lauer
	controller := NewHealthController(router, service)

	// Return the module instance
	return &Module{
		controller: controller,
	}
}

// Controller is used to get controller instance of module
func (m *Module) Controller() *Controller {
	return m.controller
}
