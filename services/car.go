package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PeterAjaaa/setor_mobil_backend/dto"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

func (h *ServiceHandler) GetCarById(w http.ResponseWriter, r *http.Request) {
	logger.LOG.Debug("Getting car in GetCarById() function...")

	w.Header().Set("Content-type", "application/json")
	var car models.Cars
	var response dto.CarResponseRequest

	result := h.Auth.DB.First(&car, r.PathValue("id"))

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("No car with ID %s found", r.PathValue("id")),
		})
		return
	}

	response = dto.CarResponseRequest{
		ID:              car.ID,
		RegistrationNum: car.RegistrationNum,
		Brand:           car.Brand,
		Model:           car.Model,
		Year:            car.Year,
		PricePerDay:     car.PricePerDay,
		Status:          car.Status,
		Description:     car.Description,
		ImageURL:        car.ImageURL,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *ServiceHandler) GetAllCars(w http.ResponseWriter, r *http.Request) {
	logger.LOG.Debug("Getting all cars in GetAllCars() function...")

	w.Header().Set("Content-type", "application/json")
	var cars []models.Cars
	var response []dto.CarResponseRequest

	result := h.Auth.DB.Find(&cars)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "No cars found",
		})
		return
	}

	for _, car := range cars {
		response = append(response, dto.CarResponseRequest{
			ID:              car.ID,
			RegistrationNum: car.RegistrationNum,
			Brand:           car.Brand,
			Model:           car.Model,
			Year:            car.Year,
			PricePerDay:     car.PricePerDay,
			Status:          car.Status,
			Description:     car.Description,
			ImageURL:        car.ImageURL,
		})
	}

	json.NewEncoder(w).Encode(response)
}
