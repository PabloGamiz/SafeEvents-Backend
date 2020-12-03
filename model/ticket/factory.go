package ticket

import (
	"log"
	"sync"

	"github.com/PabloGamiz/SafeEvents-Backend/mysql"
	"gorm.io/gorm"
)

var once sync.Once

// OpenStream ensuring the Event's table does exists
func OpenStream() (db *gorm.DB, err error) {
	if db, err = mysql.OpenStream(); err != nil {
		log.Fatalf("Got %v error while OpenStream", err.Error())
		return
	}

	once.Do(func() {
		// Automigrate must be called only once for each gateway, and allways on the stream's opening call.
		// This makes sure the Event struct has its own table on the database. So model updates are only
		// migrable to the database rebooting the server (not on-the-run).
		db.AutoMigrate(&Ticket{})
	})

	return
}

// GetTicketsByEventID return all current tickets for a given event
func GetTicketsByEventID(id uint) (tickets []Controller, err error) {
	var db *gorm.DB
	if db, err = OpenStream(); err != nil {
		return
	}

	db.Table("tickets").Where(queryFindByEventID, id).Find(&tickets)
	return
}

// GetTicketsByEventIDAndClientID return all current tickets for a given event and client
func GetTicketsByEventIDAndClientID(eid uint, cid uint) (tickets []Controller, err error) {
	var db *gorm.DB
	if db, err = OpenStream(); err != nil {
		return
	}

	db.Table("tickets").Where(queryFindByEventIDAndClientID, eid, cid).Find(&tickets)
	return
}
