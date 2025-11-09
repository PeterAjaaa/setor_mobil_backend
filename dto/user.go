package dto

import (
	"time"
)

type UserCreationRequest struct {
	Name       string    `json:"name" validate:"required,min=1"`
	Password   string    `json:"password" validate:"required,min=8"`
	Birthdate  time.Time `json:"birthdate" validate:"required"`
	Birthplace string    `json:"birthplace" validate:"required"`
	Gender     Genders   `json:"gender" validate:"required,oneof=Laki-laki Perempuan"`
	// Assuming the possible shortest address as 'Jl a no 1'
	Address    string `json:"address" validate:"required,min=9"`
	RT         string `json:"rt" validate:"required,min=1,max=3"`
	RW         string `json:"rw" validate:"required,min=1,max=3"`
	Keluarahan string `json:"kelurahan" validate:"required,min=2"`
	Kecamatan  string `json:"kecamatan" validate:"required,min=2"`
	Occupation string `json:"pekerjaan" validate:"required,min=2"`
	// Assuming the possible shortest email address as a@b.cd
	Email string `json:"email" validate:"required,min=6"`
}

type UserResponseRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Genders string

const (
	LAKI_LAKI Genders = "Laki-laki"
	PEREMPUAN Genders = "Perempuan"
)
