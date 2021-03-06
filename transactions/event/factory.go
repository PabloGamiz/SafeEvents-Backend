package event

import (
	eventDTO "github.com/PabloGamiz/SafeEvents-Backend/dtos/event"
	"github.com/alvidir/util/pattern/transaction"
)

// NewTxListEvents builds a brand new transaction for List Events
func NewTxListEvents() transaction.Tx {
	body := &txListEvents{}
	return transaction.NewTransaction(body)
}

// NewTxListEventsByType builds a brand new transaction for List Events
func NewTxListEventsByType(request eventDTO.ListEventsByTypeRequestDTO) transaction.Tx {
	body := &txListEventsByType{request: request}
	return transaction.NewTransaction(body)
}

// NewTxPublicaEvent builds a brand new transaction for Publica_event
func NewTxPublicaEvent(request eventDTO.PublicaEvent) transaction.Tx {
	body := &txPublicaEvent{request: request}
	return transaction.NewTransaction(body)
}

// NewTxModificaEvent builds a brand new transaction for modifica event
func NewTxModificaEvent(request eventDTO.DTO) transaction.Tx {
	body := &txModificaEvent{request: request}
	return transaction.NewTransaction(body)
}

// NewTxGetEvent builds a brand new transaction for Get_Event
func NewTxGetEvent(request eventDTO.GetEvent) transaction.Tx {
	body := &txGetEvent{request: request}
	return transaction.NewTransaction(body)
}

// NewTxListFavorites do this and that ...
func NewTxListFavorites(request eventDTO.ListFavoritesRequestDTO) transaction.Tx {
	body := &txListFavorites{request: request}
	return transaction.NewTransaction(body)
}

// NewTxRecomanaEvents do this and that ...
func NewTxRecomanaEvents(request eventDTO.ListFavoritesRequestDTO) transaction.Tx {
	body := &txRecomanaEvents{request: request}
	return transaction.NewTransaction(body)
}

// NewTxGetEventAnonim builds a brand new transaction for Get_Event
func NewTxGetEventAnonim(request eventDTO.GetEvent) transaction.Tx {
	body := &txGetEventAnonim{request: request}
	return transaction.NewTransaction(body)
}
