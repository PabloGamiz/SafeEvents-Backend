package buyerdata

import "github.com/PabloGamiz/SafeEvents-Backend/model/buyerdata"

// A Gateway represents the way between a model's object and the database
type Gateway interface {
	buyerdata.Controller
	Insert() error
	Update() error
	Remove() error
}
