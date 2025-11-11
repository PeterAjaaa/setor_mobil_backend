package models

import (
	"reflect"
	"time"
)

type Motorcycles struct {
	ID              uint `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	RegistrationNum string
	Brand           string
	Model           string
	Year            uint16
	PricePerDay     uint32
	Status          string
	Description     string
	ImageURL        string
	// FK Field
	CreatedByID uint
}

func (motor Motorcycles) TableName() string {
	return reflect.TypeOf(motor).Name()
}
