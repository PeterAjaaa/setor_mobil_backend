package models

import (
	"reflect"
	"time"
)

type Cars struct {
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
	CreatedByID     uint
	Orders          *[]Orders  `gorm:"foreignKey:CarID"`
	Ratings         *[]Ratings `gorm:"foreignKey:CarID"`
}

func (car Cars) TableName() string {
	return reflect.TypeOf(car).Name()
}
