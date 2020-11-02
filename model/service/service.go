package service

import (
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/location"
	"github.com/PabloGamiz/SafeEvents-Backend/model/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service represents the product class from UML
type Service struct {
	ID          primitive.ObjectID `json:"id" sql:"AUTO_INCREMENT"`
	Name        string             `json:"name" gorm:"primary_key"`
	Description string             `json:"description"`
	Kind        Kind               `json:"kind"`
	Location    location.Location  `json:"location"`
	Products    []product.Product  `json:"products"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// GetID return the ID of the Service.
func (service *Service) GetID() primitive.ObjectID {
	return service.ID
}

// SetID sets a new id to the model
func (service *Service) SetID(id primitive.ObjectID) {
	service.ID = id
}

// GetName return the Name of the Service.
func (service *Service) GetName() string {
	return service.Name
}

// SetName sets the Name of the Service.
func (service *Service) SetName(name string) {
	service.Name = name
}

// GetDescription return the Description of the Service.
func (service *Service) GetDescription() string {
	return service.Description
}

// SetDescription sets the Description of the Service.
func (service *Service) SetDescription(description string) {
	service.Description = description
}

// GetKind return the Kind of the Service.
func (service *Service) GetKind() Kind {
	return service.Kind
}

// SetKind sets the Kind of the Service.
func (service *Service) SetKind(kind Kind) {
	service.Kind = kind
}

// GetLocation return the Location of the Service.
func (service *Service) GetLocation() (loc location.Location) {
	return service.Location
}

// SetLocation sets the Description of the Service.
func (service *Service) SetLocation(location location.Location) {
	service.Location = location
}

// GetProducts return the Products of the Service.
func (service *Service) GetProducts() []product.Product {
	return service.Products
}

// SetProducts sets the Description of the Service.
func (service *Service) SetProducts(products []product.Product) {
	service.Products = products
}
