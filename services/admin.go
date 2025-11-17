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

func RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.HttpErrorHelper(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Registering admin in RegisterAdmin() function...")

	var req dto.AdminCreationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.HttpErrorHelper(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		helper.HttpErrorHelper(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error("Validation error: " + err.Error())
		return
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		helper.HttpErrorHelper(w, http.StatusInternalServerError, "Failed to hash passwowrd", nil)
		logger.LOG.Error(err.Error())
		return
	}

	admin := models.Admins{
		Name:       req.Name,
		Email:      req.Email,
		Password:   hashedPassword,
		DateJoined: time.Now(),
	}

	db := helper.GetDB()
	txErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&admin).Error; err != nil {
			return err
		}
		return nil
	})

	if txErr != nil {
		logger.LOG.Error("Error inserting data, transaction rolled back:")
		helper.HttpErrorHelper(w, http.StatusInternalServerError, txErr.Error(), nil)
		return
	} else {
		logger.LOG.Info("Successfully registered a new admin with transaction!")
	}

	response := dto.AdminResponseRequest{
		ID:          admin.ID,
		Name:        admin.Name,
		Email:       admin.Email,
		CarsCreated: admin.CarsCreated,
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusCreated,
		Message: "Successfully registered a new admin",
		Data:    response,
	})
}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.HttpErrorHelper(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Logging in admin in LoginAdmin() function...")

	var req dto.LoginRequest

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

	var admin models.Admins

	db := helper.GetDB()

	if err := db.Where("email = ?", req.Email).First(&admin).Error; err != nil {
		helper.HttpErrorHelper(w, http.StatusUnauthorized, "Invalid email or password", nil)
		logger.LOG.Error(err.Error())
		return
	}

	if !helper.CheckPasswordHash(req.Password, admin.Password) {
		helper.HttpErrorHelper(w, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	token, err := helper.GenerateJWT(admin)
	if err != nil {
		helper.HttpErrorHelper(w, http.StatusInternalServerError, "Failed to create token", nil)
		logger.LOG.Error(err.Error())
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Admin %s successfully has successfully logged in!", admin.Name),
		Data:    dto.LoginResponse{Token: token},
	})

	logger.LOG.Debug(fmt.Sprintf("Admin %s successfully has successfully logged in!", admin.Name))
}

func (h *ServiceHandler) GetAdminById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.HttpErrorHelper(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var admin models.Admins
	var response dto.AdminResponseRequest

	logger.LOG.Debug(fmt.Sprintf("Getting admin by ID:%d in GetAdminById() function...", admin.ID))

	adminID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		helper.HttpErrorHelper(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.Preload("CarsCreated").Preload("MotorcyclesCreated").First(&admin, adminID)

	if result.Error != nil {
		helper.HttpErrorHelper(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.HttpErrorHelper(w, http.StatusNotFound, fmt.Sprintf("No admin with ID %d found", adminID), nil)
		return
	}

	response = dto.AdminResponseRequest{
		ID:                 admin.ID,
		Name:               admin.Name,
		Email:              admin.Email,
		CarsCreated:        admin.CarsCreated,
		MotorcyclesCreated: admin.MotorcyclesCreated,
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Admin ID %d is found!", admin.ID),
		Data:    response,
	})

	logger.LOG.Debug(fmt.Sprintf("Admin ID %d is found!", admin.ID))
}
