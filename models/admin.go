package models

import (
	"reflect"
	"time"
)

type Admins struct {
	ID                 uint `gorm:"primaryKey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Name               string
	Email              string
	Password           string
	DateJoined         time.Time
	CarsCreated        *[]Cars        `gorm:"foreignKey:CreatedByID"`
	MotorcyclesCreated *[]Motorcycles `gorm:"foreignKey:CreatedByID"`
}

func (admin Admins) TableName() string {
	return reflect.TypeOf(admin).Name()
}

func (admin Admins) GetID() uint {
	return admin.ID
}

func (admin Admins) GetEmail() string {
	return admin.Email
}

func (admin Admins) GetName() string {
	return admin.Name
}
