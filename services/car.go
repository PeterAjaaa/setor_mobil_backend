package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PeterAjaaa/setor_mobil_backend/dto"
	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
	"gorm.io/gorm"
)

func (h *ServiceHandler) GetCarById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.LOG.Debug("Getting car in GetCarById() function...")

	w.Header().Set("Content-Type", "application/json")

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
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.LOG.Debug("Getting all cars in GetAllCars() function...")

	w.Header().Set("Content-Type", "application/json")

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

func (h *ServiceHandler) CreateNewCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.LOG.Debug("Creating new car in CreateNewCar() function...")

	var req dto.CarCreationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LOG.Error(err.Error())
		return
	}

	adminID, ok := helper.GetUserID(r.Context())

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	car := models.Cars{
		RegistrationNum: req.RegistrationNum,
		Brand:           req.Brand,
		Model:           req.Model,
		Year:            req.Year,
		PricePerDay:     req.PricePerDay,
		Status:          req.Status,
		Description:     req.Description,
		ImageURL:        req.ImageURL,
		CreatedByID:     adminID,
	}

	db := helper.GetDB()
	txErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&car).Error; err != nil {
			return err
		}
		return nil
	})

	if txErr != nil {
		logger.LOG.Error("Error inserting data, transaction rolled back:")
		http.Error(w, txErr.Error(), http.StatusInternalServerError)
		return
	} else {
		logger.LOG.Info("Successfully created new car with transaction!")
	}

	response := dto.CarResponseRequest{
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
