package dto

import (
	"errors"
)

type RatingCreationRequest struct {
	// There isn't any UserID here, because UserID is automatically set up with context inside of the service
	Rating       uint8 `json:"rating" validate:"required"`
	OrderID      uint  `json:"order_id" validate:"required"`
	CarID        *uint `json:"car_id,omitempty"`
	MotorcycleID *uint `json:"motorcycle_id,omitempty"`
}

type RatingResponseRequest struct {
	ID           uint
	Rating       uint8
	UserID       uint
	OrderID      uint
	CarID        uint
	MotorcycleID uint
}

func (o *RatingCreationRequest) Validate() error {
	if o.CarID == nil && o.MotorcycleID == nil {
		return errors.New("either car_id or motorcycle_id must be provided")
	}
	if o.CarID != nil && o.MotorcycleID != nil {
		return errors.New("cannot provide both car_id and motorcycle_id")
	}

	return nil
}
