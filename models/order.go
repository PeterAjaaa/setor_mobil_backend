package models

import (
	"reflect"
	"time"
)

type Orders struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Duration     uint8
	PickupTime   time.Time
	StartDate    time.Time
	ReturnDate   time.Time
	Price        uint32
	Status       string
	CarID        *uint
	MotorcycleID *uint
	UserID       uint
	Rating       *Ratings `gorm:"foreignKey:OrderID"`
}

func (order Orders) TableName() string {
	return reflect.TypeOf(order).Name()
}
