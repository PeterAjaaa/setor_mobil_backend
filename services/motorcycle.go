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
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting motorcycle in GetMotorcycleById() function...")

	w.Header().Set("Content-Type", "application/json")

	var motor models.Motorcycles
	var response dto.MotorcycleResponseRequest

	motorID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		helper.SendHttpResponse(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.Preload("Orders").Preload("Ratings").First(&motor, motorID)

	if result.Error != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.SendHttpResponse(w, http.StatusNotFound, fmt.Sprintf("No motorcycle with ID %d found", motorID), nil)
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

	helper.SendHttpResponse(w, http.StatusOK, fmt.Sprintf("Motorcycle ID %d is found!", motorID), response)
}

func (h *ServiceHandler) GetAllMotorcycles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting all motorcycles in GetAllMotorcycles() function...")

	w.Header().Set("Content-Type", "application/json")

	var motors []models.Motorcycles
	var response []dto.MotorcycleResponseRequest

	result := h.DB.Preload("Orders").Preload("Ratings").Find(&motors)

	if result.Error != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.SendHttpResponse(w, http.StatusNotFound, "No motorcycles found", nil)
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

	helper.SendHttpResponse(w, http.StatusOK, fmt.Sprintf("Found %d motorcycles", len(response)), response)
}

func (h *ServiceHandler) CreateNewMotorcycle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Creating new motorcycle in CreateNewMotorcycle() function...")

	var req dto.MotorcyleCreationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.SendHttpResponse(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	if err := validator.Validate(&req); err != nil {
		helper.SendHttpResponse(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error("Validation error: " + err.Error())
		return
	}

	adminID, ok := helper.GetUserID(r.Context())

	if !ok {
		helper.SendHttpResponse(w, http.StatusUnauthorized, "Unauthorized", nil)
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
		helper.SendHttpResponse(w, http.StatusInternalServerError, txErr.Error(), nil)
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

	helper.SendHttpResponse(w, http.StatusCreated, "Successfully created a new motorcycle!", response)
}

func (h *ServiceHandler) UpdateMotorcycleDetailById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Updating motorcycle details in UpdateMotorcycleDetailById() function...")
	w.Header().Set("Content-Type", "application/json")

	var motorcycle models.Motorcycles
	var req dto.MotorcycleDetailUpdateRequest

	motorcycleID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		helper.SendHttpResponse(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.SendHttpResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		logger.LOG.Error(err.Error())
		return
	}

	if err := validator.Validate(&req); err != nil {
		helper.SendHttpResponse(w, http.StatusBadRequest, "Validation failed: "+err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.First(&motorcycle, motorcycleID)
	if result.Error != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.SendHttpResponse(w, http.StatusNotFound, fmt.Sprintf("No motorcycle with ID %d found", motorcycleID), nil)
		return
	}

	err = h.DB.Transaction(func(tx *gorm.DB) error {
		// Update motorcycle details
		if req.RegistrationNum != "" {
			motorcycle.RegistrationNum = req.RegistrationNum
		}
		if req.Brand != "" {
			motorcycle.Brand = req.Brand
		}
		if req.Model != "" {
			motorcycle.Model = req.Model
		}
		if req.Year > 0 {
			motorcycle.Year = req.Year
		}
		if req.PricePerDay > 0 {
			motorcycle.PricePerDay = req.PricePerDay
		}
		if req.Status != "" {
			motorcycle.Status = req.Status
		}
		if req.ImageURL != "" {
			motorcycle.ImageURL = req.ImageURL
		}
		if req.Description != "" {
			motorcycle.Description = req.Description
		}

		if err := tx.Save(&motorcycle).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, "Failed to update motorcycle details", nil)
		logger.LOG.Error(err.Error())
		return
	}

	helper.SendHttpResponse(w, http.StatusOK, fmt.Sprintf("Motorcycle #%d details updated successfully", motorcycleID), nil)
}
