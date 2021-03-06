package event

import (
	"time"

	service_api "github.com/PabloGamiz/SafeEvents-Backend/dtos/service"
)

// FullEvent represents the expected data from an Event.
type FullEvent struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Capacity    int               `json:"capacity"`
	Organizer   string            `json:"organizers"`
	CheckInDate time.Time         `json:"checkInDate"`
	ClosureDate time.Time         `json:"closureDate"`
	Price       float32           `json:"price"`
	Location    string            `json:"location"`
	Services    []service_api.DTO `json:"services"`
	Image       string            `json:"image"`
	Tipus       string            `json:"tipus"`
	Mesures     string            `json:"mesures"`
	Faved       bool              `json:"Faved"`
	Taken       int               `json:"Taken"`
	EsOrganize  bool              `json:"esorg"`
}
