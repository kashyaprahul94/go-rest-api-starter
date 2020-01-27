package health

import (
	"encoding/json"
	"fmt"
)

// Service is the basic struct
type Service struct{}

// NewHealthService creates the instance of the health service
func NewHealthService() *Service {
	return &Service{}
}

// ToJSON converts Health model to JSON
func (s *Service) ToJSON(input *Health) []byte {
	result, err := json.Marshal(input)

	if err != nil {
		fmt.Println("cannot convert to JSON")
	}

	return result
}

// FromJSON converts JSON to Health model
func (s *Service) FromJSON(input []byte) Health {

	var health Health

	err := json.Unmarshal(input, &health)

	if err != nil {
		fmt.Println("cannot convert from JSON")
	}

	return health
}

// GetHealth retreives the health of the system
func (s *Service) GetHealth() *Health {

	health := &Health{
		Status:  "OK",
		Version: "1.0.0",
	}

	return health
}
