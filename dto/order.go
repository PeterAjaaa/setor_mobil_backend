package dto

import (
	"errors"
	"time"
)

type OrderCreationRequest struct {
	Duration       uint8     `json:"duration" validate:"required,min=1"`
	PickupTime     time.Time `json:"pickup_time" validate:"required"`
	PickupLocation string    `json:"pickup_location" validate:"required"`
	Price          uint32    `json:"price" validate:"required"`
	Status         string    `json:"status" validate:"required,oneof=Active Pending Completed"`
	CarID          *uint     `json:"car_id,omitempty"`
	MotorcycleID   *uint     `json:"motorcycle_id,omitempty"`
	// There isn't any UserID here, because UserID is automatically set up with context inside of the service
}

type OrderResponseRequest struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Duration       uint8
	Rating         *uint
	PickupTime     time.Time
	PickupLocation string
	Price          uint32
	Status         string
	CarID          uint
	MotorcycleID   uint
	UserID         uint
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
