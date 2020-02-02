package health

import (
	"net/http"

	router "github.com/julienschmidt/httprouter"

	"github.com/kashyaprahul94/go-rest-api-starter/pkg/api/server"
)

// BasePath on which this route would be mounted
const BasePath = "/health"

// NewHealthController creates instance of Heatlh Controller
func NewHealthController(router *server.RouteGroup) *Controller {

	// Create new instance of service
	serviceInstance := NewHealthService()

	// Initialize controller
	controller := &Controller{
		router:  router,
		service: serviceInstance,
	}

	// Register the routes
	controller.registerRoutes()

	// return the controller instance
	return controller
}

// GetRouter return this router of Heatlh Controller
func (hc *Controller) GetRouter() *server.RouteGroup {
	return hc.router
}

// registerRoutes is used to register the routes
func (hc *Controller) registerRoutes() {

	// Get the router instance
	router := hc.GetRouter()

	// Register / route to send health of the system
	router.GET("/", hc.get)
}

// get is used for - GET -> /health
func (hc *Controller) get(w http.ResponseWriter, r *http.Request, _ router.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	service := hc.service

	healthStatus := service.GetHealth()
	response := service.ToJSON(healthStatus)

	w.Write(response)
}
