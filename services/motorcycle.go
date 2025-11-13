package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PeterAjaaa/setor_mobil_backend/dto"
	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
	"github.com/PeterAjaaa/setor_mobil_backend/validator"
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

	motorID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.First(&motor, motorID)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(models.APIResponse{
			Status:  http.StatusNotFound,
			Message: fmt.Sprintf("No motorcycle with ID %d found", motorID),
			Data:    nil,
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

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Motorcycle ID %d is found!", motorID),
		Data:    response,
	})
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

	result := h.DB.Find(&motors)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(models.APIResponse{
			Status:  http.StatusOK,
			Message: "No motorcycles found",
			Data:    nil,
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

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Found %d motorcycles", len(response)),
		Data:    response,
	})
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

	if err := validator.Validate(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LOG.Error("Validation error: " + err.Error())
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
		logger.LOG.Info("Successfully created a new motorcycle with transaction!")
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
	json.NewEncoder(w).Encode(
		models.APIResponse{
			Status:  http.StatusCreated,
			Message: "Successfully created a new motorcycle!",
			Data:    response,
		})

}
