package organizer

// Parent is the Client who is being Organizer of some events
type Parent interface {
	GetID() uint
}

// Controller for Organizer
type Controller interface {
	SetParent(Parent)
}
