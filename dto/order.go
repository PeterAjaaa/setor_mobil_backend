package dto

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

type OrderCreationRequest struct {
	Duration     uint8     `json:"duration" validate:"required,min=1"`
	PickupTime   time.Time `json:"pickup_time" validate:"required"`
	StartDate    time.Time `json:"start_date" validate:"required"`
	ReturnDate   time.Time `json:"return_date" validate:"required"`
	Price        uint32    `json:"price" validate:"required"`
	Status       string    `json:"status" validate:"required,oneof=Active Pending Completed"`
	CarID        *uint     `json:"car_id,omitempty"`
	MotorcycleID *uint     `json:"motorcycle_id,omitempty"`
	// There isn't any UserID here, because UserID is automatically set up with context inside of the service
}

type OrderResponseRequest struct {
	ID           uint            `json:"id"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	Duration     uint8           `json:"duration"`
	Rating       *models.Ratings `json:"rating"`
	PickupTime   time.Time       `json:"pickup_time"`
	StartDate    time.Time       `json:"start_date"`
	ReturnDate   time.Time       `json:"return_date"`
	Price        uint32          `json:"price"`
	Status       string          `json:"status"`
	CarID        *uint           `json:"car_id"`
	MotorcycleID *uint           `json:"motorcycle_id"`
	UserID       uint            `json:"user_id"`
}

type OrderStatusUpdateRequest struct {
	Status string `json:"status" validate:"required,oneof=Active Pending Completed"`
}

func (o *OrderCreationRequest) Validate() error {
	if o.CarID == nil && o.MotorcycleID == nil {
		return errors.New("either car_id or motorcycle_id must be provided")
	}
	if o.CarID != nil && o.MotorcycleID != nil {
		return errors.New("cannot provide both car_id and motorcycle_id")
	}

	return nil
}

func (o OrderResponseRequest) MarshalJSON() ([]byte, error) {
	type Alias OrderResponseRequest

	return json.Marshal(&struct {
		PickupTime string `json:"pickup_time"`
		StartDate  string `json:"start_date"`
		ReturnDate string `json:"return_date"`
		*Alias
	}{
		PickupTime: o.PickupTime.Format(time.RFC3339),
		StartDate:  o.StartDate.Format(time.RFC3339),
		ReturnDate: o.ReturnDate.Format(time.RFC3339),
		Alias:      (*Alias)(&o),
	})
}
