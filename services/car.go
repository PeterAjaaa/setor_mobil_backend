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

func (h *ServiceHandler) GetCarById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting car in GetCarById() function...")

	w.Header().Set("Content-Type", "application/json")

	var car models.Cars
	var response dto.CarResponseRequest

	carID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		helper.SendHttpResponse(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.Preload("Orders").Preload("Ratings").First(&car, carID)

	if result.Error != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.SendHttpResponse(w, http.StatusNotFound, fmt.Sprintf("No car with ID %d found", carID), nil)
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
		Orders:          car.Orders,
		Ratings:         car.Ratings,
	}

	helper.SendHttpResponse(w, http.StatusOK, fmt.Sprintf("Car ID %d is found!", carID), response)
}

func (h *ServiceHandler) GetAllCars(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting all cars in GetAllCars() function...")

	w.Header().Set("Content-Type", "application/json")

	var cars []models.Cars
	var response []dto.CarResponseRequest

	result := h.DB.Preload("Orders").Preload("Ratings").Find(&cars)

	if result.Error != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.SendHttpResponse(w, http.StatusNotFound, "No cars found", nil)
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
			Orders:          car.Orders,
			Ratings:         car.Ratings,
		})
	}

	helper.SendHttpResponse(w, http.StatusOK, fmt.Sprintf("Found %d cars", len(response)), response)
}

func (h *ServiceHandler) CreateNewCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Creating new car in CreateNewCar() function...")

	var req dto.CarCreationRequest

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
		helper.SendHttpResponse(w, http.StatusInternalServerError, txErr.Error(), nil)
		return
	} else {
		logger.LOG.Info("Successfully created a new car with transaction!")
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

	helper.SendHttpResponse(w, http.StatusCreated, "Successfully created a new car!", response)
}

func (h *ServiceHandler) UpdateCarDetailById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Updating car details in UpdateCarDetailById() function...")
	w.Header().Set("Content-Type", "application/json")

	var car models.Cars
	var req dto.CarDetailUpdateRequest

	carID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
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

	result := h.DB.First(&car, carID)
	if result.Error != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.SendHttpResponse(w, http.StatusNotFound, fmt.Sprintf("No car with ID %d found", carID), nil)
		return
	}

	err = h.DB.Transaction(func(tx *gorm.DB) error {
		if req.RegistrationNum != "" {
			car.RegistrationNum = req.RegistrationNum
		}
		if req.Brand != "" {
			car.Brand = req.Brand
		}
		if req.Model != "" {
			car.Model = req.Model
		}
		if req.Year > 0 {
			car.Year = uint16(req.Year)
		}
		if req.PricePerDay > 0 {
			car.PricePerDay = uint32(req.PricePerDay)
		}
		if req.Status != "" {
			car.Status = req.Status
		}
		if req.ImageURL != "" {
			car.ImageURL = req.ImageURL
		}
		if req.Description != "" {
			car.Description = req.Description
		}

		if err := tx.Save(&car).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, "Failed to update car details", nil)
		logger.LOG.Error(err.Error())
		return
	}

	helper.SendHttpResponse(w, http.StatusOK, fmt.Sprintf("Car #%d details updated successfully", carID), nil)
}

func (h *ServiceHandler) DeleteCarById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Deleting car in DeleteCarById() function...")
	w.Header().Set("Content-Type", "application/json")

	var car models.Cars

	carID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		helper.SendHttpResponse(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.First(&car, carID)
	if result.Error != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.SendHttpResponse(w, http.StatusNotFound, fmt.Sprintf("No car with ID %d found", carID), nil)
		return
	}

	err = h.DB.Delete(&car).Error
	if err != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, "Failed to delete car", nil)
		logger.LOG.Error(err.Error())
		return
	}

	helper.SendHttpResponse(w, http.StatusOK, fmt.Sprintf("Car #%d deleted successfully", carID), nil)
}
