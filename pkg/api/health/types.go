package health

// Health is the model
type Health struct {

	// Status of the system
	Status string `json:"status"`

	// Version of the application
	Version string `json:"version"`
}
