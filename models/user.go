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
	RT         string
	RW         string
	Keluarahan string
	Kecamatan  string
	Occupation string
	Email      string
}

func (user Users) TableName() string {
	return reflect.TypeOf(user).Name()
}
