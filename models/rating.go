package models

import (
	"reflect"
	"time"
)

type Ratings struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Rating       uint8
	UserID       uint
	OrderID      uint
	CarID        *uint
	MotorcycleID *uint
}

func (rating Ratings) TableName() string {
	return reflect.TypeOf(rating).Name()
}
