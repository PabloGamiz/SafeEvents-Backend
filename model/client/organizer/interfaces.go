package organizer

import "github.com/PabloGamiz/SafeEvents-Backend/model/event"

// Parent is the Client who is being Organizer of some events
type Parent interface {
}

// Controller for Organizer
type Controller interface {
	SetParent(Parent)
	AddEvent(ctrl event.Controller) int
}
