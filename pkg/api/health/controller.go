package health

import (
	"net/http"

	"github.com/gorilla/mux"
	services "github.com/kashyaprahul94/go-rest-api-starter/pkg/api/health/services"
)

// BasePath on which this route would be mounted
const BasePath = "/health"

// Controller is kind of API Controller responsible for heatlh
type Controller struct {

	// router instance to handle all the incoming http request
	router *mux.Router

	// service which does all the necessary operations
	service *services.Health
}

// New creates instance of Heatlh Controller
func New(router *mux.Router) *Controller {

	// Create new instance of service
	serviceInstance := services.New()

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
func (hc *Controller) GetRouter() *mux.Router {
	return hc.router
}

// registerRoutes is used to register the routes
func (hc *Controller) registerRoutes() {

	// Get the router instance
	router := hc.GetRouter()

	// Register / route to send health of the system
	router.HandleFunc("/", hc.get).Methods(http.MethodGet)
}

// get is used for - GET -> /health
func (hc *Controller) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	service := hc.service

	healthStatus := service.GetHealth()
	response := service.ToJSON(healthStatus)

	w.Write(response)
}
