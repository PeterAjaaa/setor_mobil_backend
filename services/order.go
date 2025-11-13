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

func (h *ServiceHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.LOG.Debug("Getting order in GetOrderById() function...")

	w.Header().Set("Content-Type", "application/json")

	var order models.Orders
	var response dto.OrderResponseRequest

	orderID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.First(&order, orderID)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		json.NewEncoder(w).Encode(
			models.APIResponse{
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("No order with ID %s found", r.PathValue("id")),
				Data:    nil,
			},
		)
		return
	}

	response = dto.OrderResponseRequest{
		ID:             order.ID,
		CreatedAt:      order.CreatedAt,
		Duration:       order.Duration,
		Rating:         order.Rating,
		PickupTime:     order.PickupTime,
		PickupLocation: order.PickupLocation,
		Price:          order.Price,
		Status:         order.Status,
		CarID:          helper.UintValue(order.CarID),
		MotorcycleID:   helper.UintValue(order.MotorcycleID),
		UserID:         order.UserID,
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Found order ID %d", orderID),
		Data:    response,
	})
}

func (h *ServiceHandler) GetAllOrdersByUserId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.LOG.Debug("Getting orders by user id in GetAllOrdersByUserId() function...")

	w.Header().Set("Content-Type", "application/json")

	var orders []models.Orders
	var response []dto.OrderResponseRequest

	userID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.Where("user_id = ?", userID).Find(&orders)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		json.NewEncoder(w).Encode(
			models.APIResponse{
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("No order for user ID %d found", userID),
				Data:    nil,
			},
		)
		return
	}

	for _, order := range orders {
		response = append(response, dto.OrderResponseRequest{
			ID:             order.ID,
			CreatedAt:      order.CreatedAt,
			Duration:       order.Duration,
			Rating:         order.Rating,
			PickupTime:     order.PickupTime,
			PickupLocation: order.PickupLocation,
			Price:          order.Price,
			Status:         order.Status,
			CarID:          helper.UintValue(order.CarID),
			MotorcycleID:   helper.UintValue(order.MotorcycleID),
			UserID:         order.UserID,
		})
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Found %d orders by user ID %d", len(response), userID),
		Data:    response,
	})
}

func (h *ServiceHandler) CreateNewOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.LOG.Debug("Creating new order in CreateNewOrder() function...")

	var req dto.OrderCreationRequest

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

	userID, ok := helper.GetUserID(r.Context())

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	order := models.Orders{
		Duration:       req.Duration,
		PickupTime:     req.PickupTime,
		PickupLocation: req.PickupLocation,
		Price:          req.Price,
		Status:         req.Status,
		CarID:          req.CarID,
		MotorcycleID:   req.MotorcycleID,
		UserID:         userID,
	}

	db := helper.GetDB()
	txErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}
		return nil
	})

	if txErr != nil {
		logger.LOG.Error("Error inserting data, transaction rolled back:")
		http.Error(w, txErr.Error(), http.StatusInternalServerError)
		return
	} else {
		logger.LOG.Info("Successfully created a new order with transaction!")
	}

	response := dto.OrderResponseRequest{
		ID:             order.ID,
		CreatedAt:      order.CreatedAt,
		Duration:       order.Duration,
		Rating:         order.Rating,
		PickupTime:     order.PickupTime,
		PickupLocation: order.PickupLocation,
		Price:          order.Price,
		Status:         order.Status,
		CarID:          helper.UintValue(order.CarID),
		MotorcycleID:   helper.UintValue(order.MotorcycleID),
		UserID:         order.UserID,
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: "Successfully created a new order with transaction",
		Data:    response,
	})

}
