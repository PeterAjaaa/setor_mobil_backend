package dto

import "github.com/PeterAjaaa/setor_mobil_backend/models"

type AdminCreationRequest struct {
	Name     string `json:"name" validate:"required,min=1"`
	Email    string `json:"email" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=8"`
}

type AdminResponseRequest struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	CarsCreated *[]models.Cars `json:"cars_created"`
}
