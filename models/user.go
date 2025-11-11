package models

import (
	"reflect"
	"time"
)

type Users struct {
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	Password   string
	Birthdate  time.Time
	Birthplace string
	Gender     string
	Address    string
	RT         uint16
	RW         uint16
	Keluarahan string
	Kecamatan  string
	Occupation string
	Email      string
	// TODO: Turn this back on when Orders has been implemented
	// Orders     *[]Orders
}

func (user Users) TableName() string {
	return reflect.TypeOf(user).Name()
}

func (user Users) GetID() uint {
	return user.ID
}

func (user Users) GetEmail() string {
	return user.Email
}

func (user Users) GetName() string {
	return user.Name
}
