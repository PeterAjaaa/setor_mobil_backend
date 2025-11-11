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

func (h *ServiceHandler) GetMotorcyleById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.LOG.Debug("Getting motorcycle in GetMotorcycleById() function...")

	w.Header().Set("Content-Type", "application/json")

	var motor models.Motorcycles
	var response dto.MotorcycleResponseRequest

	result := h.Auth.DB.First(&motor, r.PathValue("id"))

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("No motorcycle with ID %s found", r.PathValue("id")),
		})
		return
	}

	response = dto.MotorcycleResponseRequest{
		ID:              motor.ID,
		RegistrationNum: motor.RegistrationNum,
		Brand:           motor.Brand,
		Model:           motor.Model,
		Year:            motor.Year,
		PricePerDay:     motor.PricePerDay,
		Status:          motor.Status,
		Description:     motor.Description,
		ImageURL:        motor.ImageURL,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *ServiceHandler) GetAllMotorcycles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.LOG.Debug("Getting all motorcycles in GetAllMotorcycles() function...")

	w.Header().Set("Content-Type", "application/json")

	var motors []models.Motorcycles
	var response []dto.MotorcycleResponseRequest

	result := h.Auth.DB.Find(&motors)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "No motorcycles found",
		})
		return
	}

	for _, motor := range motors {
		response = append(response, dto.MotorcycleResponseRequest{
			ID:              motor.ID,
			RegistrationNum: motor.RegistrationNum,
			Brand:           motor.Brand,
			Model:           motor.Model,
			Year:            motor.Year,
			PricePerDay:     motor.PricePerDay,
			Status:          motor.Status,
			Description:     motor.Description,
			ImageURL:        motor.ImageURL,
		})
	}

	json.NewEncoder(w).Encode(response)
}

func (h *ServiceHandler) CreateNewMotorcycle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.LOG.Debug("Creating new motorcycle in CreateNewMotorcycle() function...")

	var req dto.MotorcyleCreationRequest

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

	motor := models.Motorcycles{
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
		if err := tx.Create(&motor).Error; err != nil {
			return err
		}
		return nil
	})

	if txErr != nil {
		logger.LOG.Error("Error inserting data, transaction rolled back:")
		http.Error(w, txErr.Error(), http.StatusInternalServerError)
		return
	} else {
		logger.LOG.Info("Successfully created new motorcycle with transaction!")
	}

	response := dto.MotorcycleResponseRequest{
		ID:              motor.ID,
		RegistrationNum: motor.RegistrationNum,
		Brand:           motor.Brand,
		Model:           motor.Model,
		Year:            motor.Year,
		PricePerDay:     motor.PricePerDay,
		Status:          motor.Status,
		Description:     motor.Description,
		ImageURL:        motor.ImageURL,
	}
	json.NewEncoder(w).Encode(response)

}
