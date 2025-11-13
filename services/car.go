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
		helper.HttpErrorHelper(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting car in GetCarById() function...")

	w.Header().Set("Content-Type", "application/json")

	var car models.Cars
	var response dto.CarResponseRequest

	carID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		helper.HttpErrorHelper(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.Preload("Orders").Preload("Ratings").First(&car, carID)

	if result.Error != nil {
		helper.HttpErrorHelper(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(models.APIResponse{
			Status:  http.StatusNotFound,
			Message: fmt.Sprintf("No car with ID %d found", carID),
			Data:    "",
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

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Car ID %d is found!", carID),
		Data:    response,
	})
}

func (h *ServiceHandler) GetAllCars(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.HttpErrorHelper(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting all cars in GetAllCars() function...")

	w.Header().Set("Content-Type", "application/json")

	var cars []models.Cars
	var response []dto.CarResponseRequest

	result := h.DB.Preload("Orders").Preload("Ratings").Find(&cars)

	if result.Error != nil {
		helper.HttpErrorHelper(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(models.APIResponse{
			Status:  http.StatusNotFound,
			Message: "No cars found",
			Data:    nil,
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

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Found %d cars", len(response)),
		Data:    response,
	})
}

func (h *ServiceHandler) CreateNewCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.HttpErrorHelper(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Creating new car in CreateNewCar() function...")

	var req dto.CarCreationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.HttpErrorHelper(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	if err := validator.Validate(&req); err != nil {
		helper.HttpErrorHelper(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error("Validation error: " + err.Error())
		return
	}

	adminID, ok := helper.GetUserID(r.Context())

	if !ok {
		helper.HttpErrorHelper(w, http.StatusUnauthorized, "Unauthorized", nil)
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
		helper.HttpErrorHelper(w, http.StatusInternalServerError, txErr.Error(), nil)
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

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusCreated,
		Message: "Successfully created a new car!",
		Data:    response,
	})

}
