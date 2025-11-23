package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PeterAjaaa/setor_mobil_backend/dto"
	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
	"github.com/PeterAjaaa/setor_mobil_backend/validator"
	"gorm.io/gorm"
)

func (h *ServiceHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting order in GetOrderById() function...")

	w.Header().Set("Content-Type", "application/json")

	var order models.Orders
	var response dto.OrderResponseRequest

	orderID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		helper.SendHttpResponse(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.Preload("Rating").First(&order, orderID)

	if result.Error != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.SendHttpResponse(w, http.StatusNotFound, fmt.Sprintf("No order with ID %d found", orderID), nil)
		return
	}

	loc, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, "Failed to load timezone", nil)
	}

	response = dto.OrderResponseRequest{
		ID:           order.ID,
		CreatedAt:    order.CreatedAt,
		Duration:     order.Duration,
		PickupTime:   order.PickupTime.In(loc),
		StartDate:    order.StartDate.In(loc),
		ReturnDate:   order.ReturnDate.In(loc),
		Price:        order.Price,
		Status:       order.Status,
		CarID:        order.CarID,
		MotorcycleID: order.MotorcycleID,
		UserID:       order.UserID,
		Rating:       order.Rating,
	}

	helper.SendHttpResponse(w, http.StatusOK, fmt.Sprintf("Found order ID %d", orderID), response)
}

func (h *ServiceHandler) GetAllOrdersByUserId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Getting orders by user id in GetAllOrdersByUserId() function...")

	w.Header().Set("Content-Type", "application/json")

	var orders []models.Orders
	var response []dto.OrderResponseRequest

	userID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		helper.SendHttpResponse(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.Where("user_id = ?", userID).Preload("Rating").Find(&orders)

	if result.Error != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.SendHttpResponse(w, http.StatusNotFound, fmt.Sprintf("No order for user ID %d found", userID), nil)
		return
	}

	loc, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, "Failed to load timezone", nil)
	}

	for _, order := range orders {
		response = append(response, dto.OrderResponseRequest{
			ID:           order.ID,
			CreatedAt:    order.CreatedAt,
			Duration:     order.Duration,
			PickupTime:   order.PickupTime.In(loc),
			StartDate:    order.StartDate.In(loc),
			ReturnDate:   order.ReturnDate.In(loc),
			Price:        order.Price,
			Status:       order.Status,
			CarID:        order.CarID,
			MotorcycleID: order.MotorcycleID,
			UserID:       order.UserID,
			Rating:       order.Rating,
		})
	}

	helper.SendHttpResponse(w, http.StatusOK, fmt.Sprintf("Found %d orders by user ID %d", len(response), userID), response)
}

func (h *ServiceHandler) CreateNewOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Creating new order in CreateNewOrder() function...")

	var req dto.OrderCreationRequest

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

	userID, ok := helper.GetUserID(r.Context())

	if !ok {
		helper.SendHttpResponse(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	loc, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, "Failed to load timezone", nil)
	}

	order := models.Orders{
		Duration:     req.Duration,
		PickupTime:   req.PickupTime.In(loc),
		StartDate:    req.StartDate.In(loc),
		ReturnDate:   req.ReturnDate.In(loc),
		Price:        req.Price,
		Status:       req.Status,
		CarID:        req.CarID,
		MotorcycleID: req.MotorcycleID,
		UserID:       userID,
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
		helper.SendHttpResponse(w, http.StatusInternalServerError, txErr.Error(), nil)
		return
	} else {
		logger.LOG.Info("Successfully created a new order with transaction!")
	}

	response := dto.OrderResponseRequest{
		ID:           order.ID,
		CreatedAt:    order.CreatedAt,
		Duration:     order.Duration,
		PickupTime:   order.PickupTime.In(loc),
		StartDate:    order.StartDate.In(loc),
		ReturnDate:   order.ReturnDate.In(loc),
		Price:        order.Price,
		Status:       order.Status,
		CarID:        order.CarID,
		MotorcycleID: order.MotorcycleID,
		UserID:       order.UserID,
	}

	helper.SendHttpResponse(w, http.StatusOK, "Successfully created a new order", response)
}
