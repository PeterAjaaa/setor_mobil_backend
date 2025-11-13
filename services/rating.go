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

func (h *ServiceHandler) GetRatingById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.HttpErrorHelper(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting rating in GetRatingById() function...")

	w.Header().Set("Content-Type", "application/json")

	var rating models.Ratings
	var response dto.RatingResponseRequest

	ratingID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		helper.HttpErrorHelper(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.First(&rating, ratingID)

	if result.Error != nil {
		helper.HttpErrorHelper(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.HttpErrorHelper(w, http.StatusNotFound, result.Error.Error(), nil)
		json.NewEncoder(w).Encode(
			models.APIResponse{
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("No rating with ID %d found", ratingID),
				Data:    nil,
			},
		)
		return
	}

	response = dto.RatingResponseRequest{
		ID:           rating.ID,
		Rating:       rating.Rating,
		UserID:       rating.UserID,
		OrderID:      rating.OrderID,
		CarID:        helper.UintValue(rating.CarID),
		MotorcycleID: helper.UintValue(rating.MotorcycleID),
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Found order ID %d", ratingID),
		Data:    response,
	})
}

func (h *ServiceHandler) GetAllRatingByOrderID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.HttpErrorHelper(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting ratings by order id in GetAllRatingByOrderID() function...")

	w.Header().Set("Content-Type", "application/json")

	var ratings []models.Ratings
	var response []dto.RatingResponseRequest

	orderID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		helper.HttpErrorHelper(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.Where("order_id = ?", orderID).Find(&ratings)

	if result.Error != nil {
		helper.HttpErrorHelper(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.HttpErrorHelper(w, http.StatusNotFound, result.Error.Error(), nil)
		json.NewEncoder(w).Encode(
			models.APIResponse{
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("No ratings for order ID %d found", orderID),
				Data:    nil,
			},
		)
		return
	}

	for _, rating := range ratings {
		response = append(response, dto.RatingResponseRequest{
			ID:           rating.ID,
			Rating:       rating.Rating,
			UserID:       rating.UserID,
			OrderID:      rating.OrderID,
			CarID:        helper.UintValue(rating.CarID),
			MotorcycleID: helper.UintValue(rating.MotorcycleID),
		})
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Found %d ratings for order ID %d", len(response), orderID),
		Data:    response,
	})
}

func (h *ServiceHandler) GetAllRatingByUserID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.HttpErrorHelper(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting ratings by user id in GetAllRatingByUserID() function...")

	w.Header().Set("Content-Type", "application/json")

	var ratings []models.Ratings
	var response []dto.RatingResponseRequest

	userID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		helper.HttpErrorHelper(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.Where("user_id = ?", userID).Find(&ratings)

	if result.Error != nil {
		helper.HttpErrorHelper(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.HttpErrorHelper(w, http.StatusNotFound, result.Error.Error(), nil)
		json.NewEncoder(w).Encode(
			models.APIResponse{
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("No ratings by user ID %d found", userID),
				Data:    nil,
			},
		)
		return
	}

	for _, rating := range ratings {
		response = append(response, dto.RatingResponseRequest{
			ID:           rating.ID,
			Rating:       rating.Rating,
			UserID:       rating.UserID,
			OrderID:      rating.OrderID,
			CarID:        helper.UintValue(rating.CarID),
			MotorcycleID: helper.UintValue(rating.MotorcycleID),
		})
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Found %d ratings by user ID %d", len(response), userID),
		Data:    response,
	})
}

func (h *ServiceHandler) GetAllRatingByCarID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.HttpErrorHelper(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting ratings for car id in GetAllRatingByCarID() function...")

	w.Header().Set("Content-Type", "application/json")

	var ratings []models.Ratings
	var response []dto.RatingResponseRequest

	carID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		helper.HttpErrorHelper(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.Where("car_id = ?", carID).Find(&ratings)

	if result.Error != nil {
		helper.HttpErrorHelper(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.HttpErrorHelper(w, http.StatusNotFound, result.Error.Error(), nil)
		json.NewEncoder(w).Encode(
			models.APIResponse{
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("No ratings for car ID %d found", carID),
				Data:    nil,
			},
		)
		return
	}

	for _, rating := range ratings {
		response = append(response, dto.RatingResponseRequest{
			ID:           rating.ID,
			Rating:       rating.Rating,
			UserID:       rating.UserID,
			OrderID:      rating.OrderID,
			CarID:        helper.UintValue(rating.CarID),
			MotorcycleID: helper.UintValue(rating.MotorcycleID),
		})
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Found %d ratings for car ID %d", len(response), carID),
		Data:    response,
	})
}

func (h *ServiceHandler) GetAllRatingByMotorcycleID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.HttpErrorHelper(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting ratings for motorcycle id in GetAllRatingByMotorcycleID() function...")

	w.Header().Set("Content-Type", "application/json")

	var ratings []models.Ratings
	var response []dto.RatingResponseRequest

	motorID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		helper.HttpErrorHelper(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.Where("motorcycle_id = ?", motorID).Find(&ratings)

	if result.Error != nil {
		helper.HttpErrorHelper(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			models.APIResponse{
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("No ratings for motorcycle ID %d found", motorID),
				Data:    nil,
			},
		)
		return
	}

	for _, rating := range ratings {
		response = append(response, dto.RatingResponseRequest{
			ID:           rating.ID,
			Rating:       rating.Rating,
			UserID:       rating.UserID,
			OrderID:      rating.OrderID,
			CarID:        helper.UintValue(rating.CarID),
			MotorcycleID: helper.UintValue(rating.MotorcycleID),
		})
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Found %d ratings for motorcycle ID %d", len(response), motorID),
		Data:    response,
	})
}

func (h *ServiceHandler) CreateNewRating(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.HttpErrorHelper(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Creating new rating in CreateNewRating() function...")

	var req dto.RatingCreationRequest

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

	userID, ok := helper.GetUserID(r.Context())

	if !ok {
		helper.HttpErrorHelper(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	rating := models.Ratings{
		Rating:       req.Rating,
		UserID:       userID,
		OrderID:      req.OrderID,
		CarID:        req.CarID,
		MotorcycleID: req.MotorcycleID,
	}

	db := helper.GetDB()
	txErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&rating).Error; err != nil {
			return err
		}
		return nil
	})

	if txErr != nil {
		logger.LOG.Error("Error inserting data, transaction rolled back:")
		helper.HttpErrorHelper(w, http.StatusInternalServerError, txErr.Error(), nil)
		return
	} else {
		logger.LOG.Info("Successfully created a new rating with transaction!")
	}

	response := dto.RatingResponseRequest{
		ID:           rating.ID,
		Rating:       rating.Rating,
		UserID:       rating.UserID,
		OrderID:      rating.OrderID,
		CarID:        helper.UintValue(rating.CarID),
		MotorcycleID: helper.UintValue(rating.MotorcycleID),
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: "Successfully created a new rating",
		Data:    response,
	})

}
