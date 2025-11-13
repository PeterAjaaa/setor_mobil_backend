package dto

import "github.com/PeterAjaaa/setor_mobil_backend/models"

type CarCreationRequest struct {
	// Assuming the possible shortest commercially available registration number as 'A 1'
	RegistrationNum string `json:"registration_num" validate:"required,min=3"`
	Brand           string `json:"brand" validate:"required,min=3"`
	Model           string `json:"model" validate:"required,min=3"`
	Year            uint16 `json:"year" validate:"required"`
	PricePerDay     uint32 `json:"price_per_day" validate:"required"`
	Status          string `json:"status" validate:"required,oneof=Available Rented Maintenance"`
	Description     string `json:"description"`
	ImageURL        string `json:"image_url" validate:"required"`
}

type CarResponseRequest struct {
	ID              uint             `json:"id"`
	RegistrationNum string           `json:"registration_num"`
	Brand           string           `json:"brand"`
	Model           string           `json:"model"`
	Year            uint16           `json:"year"`
	PricePerDay     uint32           `json:"price_per_day"`
	Status          string           `json:"status"`
	Description     string           `json:"description"`
	ImageURL        string           `json:"image_url"`
	Orders          *[]models.Orders `json:"orders"`
}
