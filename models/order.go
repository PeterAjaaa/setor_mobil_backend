package models

import (
	"reflect"
	"time"
)

type Orders struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Duration       uint8
	PickupTime     time.Time
	PickupLocation string
	Price          uint32
	Status         string
	Rating         *uint
	CarID          *uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MotorcycleID   *uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID         uint
}

func (order Orders) TableName() string {
	return reflect.TypeOf(order).Name()
}
