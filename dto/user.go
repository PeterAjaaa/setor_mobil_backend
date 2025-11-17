package dto

import (
	"time"

	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

type UserCreationRequest struct {
	Name       string    `json:"name" validate:"required,min=1"`
	Password   string    `json:"password" validate:"required,min=8"`
	Birthdate  time.Time `json:"birthdate" validate:"required"`
	Birthplace string    `json:"birthplace" validate:"required"`
	Gender     string    `json:"gender" validate:"required,oneof=Laki-laki Perempuan"`
	// Assuming the possible shortest address as 'Jl a no 1'
	Address    string `json:"address" validate:"required,min=9"`
	RT         uint16 `json:"rt" validate:"required,min=1,max=999"`
	RW         uint16 `json:"rw" validate:"required,min=1,max=999"`
	Keluarahan string `json:"kelurahan" validate:"required,min=2"`
	Kecamatan  string `json:"kecamatan" validate:"required,min=2"`
	Occupation string `json:"pekerjaan" validate:"required,min=2"`
	// Assuming the possible shortest email address as a@b.cd
	Email string `json:"email" validate:"required,min=6"`
}

type UserResponseRequest struct {
	ID     uint             `json:"id"`
	Name   string           `json:"name"`
	Email  string           `json:"email"`
	Orders *[]models.Orders `json:"orders"`
}
