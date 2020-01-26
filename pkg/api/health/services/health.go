package services

import (
	"encoding/json"
	"fmt"

	"github.com/kashyaprahul94/go-rest-api-starter/pkg/api/health/models"
)

// Health is the basic struct
type Health struct{}

// New creates the instance of the health service
func New() *Health {
	return &Health{}
}

// ToJSON converts Health model to JSON
func (s *Health) ToJSON(input *models.Health) []byte {
	result, err := json.Marshal(input)

	if err != nil {
		fmt.Println("cannot convert to JSON")
	}

	return result
}

// FromJSON converts JSON to Health model
func (s *Health) FromJSON(input []byte) models.Health {

	var health models.Health

	err := json.Unmarshal(input, &health)

	if err != nil {
		fmt.Println("cannot convert from JSON")
	}

	return health
}

// GetHealth retreives the health of the system
func (s *Health) GetHealth() *models.Health {

	health := &models.Health{
		Status:  "OK",
		Version: "1.0.0",
	}

	return health
}
