package event

import (
	"context"
	"log"

	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	eventGW "github.com/PabloGamiz/SafeEvents-Backend/gateway/event"
	eventMOD "github.com/PabloGamiz/SafeEvents-Backend/model/event"
)

// txPublicaEvent represents an
type txPublicaEvent struct {
	request eventDTO.DTO
}

func (tx *txPublicaEvent) Precondition() (err error) { //Comprova que no existeixi l'event
	return
}

func (tx *txPublicaEvent) Postcondition(ctx context.Context) (interface{}, error) {
	log.Printf("Got a Publica Event request for event %s", tx.request.Title)

	evnt := &eventMOD.Event{
		Title:       tx.request.Title,
		Description: tx.request.Description,
		Capacity:    tx.request.Capacity,
		CheckInDate: tx.request.CheckInDate,
		ClosureDate: tx.request.ClosureDate,
		LocationID:  uint64(tx.request.Locations.ID),
	}
	gw := eventGW.NewEventGateway(ctx, evnt)
	err := gw.Insert()

	return gw, err
}

func (tx *txPublicaEvent) Commit() error {
	return nil
}

func (tx *txPublicaEvent) Rollback() {

}