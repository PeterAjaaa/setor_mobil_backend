package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PeterAjaaa/setor_mobil_backend/dto"
	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
	"gorm.io/gorm"
)

func RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.LOG.Debug("Registering admin in RegisterAdmin() function...")

	var req dto.AdminCreationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LOG.Error(err.Error())
		return
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
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
		http.Error(w, txErr.Error(), http.StatusInternalServerError)
		return
	} else {
		logger.LOG.Info("Successfully registered new user with transaction!")
	}

	response := dto.AdminResponseRequest{
		ID:          admin.ID,
		Name:        admin.Name,
		Email:       admin.Email,
		CarsCreated: admin.CarsCreated,
	}
	json.NewEncoder(w).Encode(response)
}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.LOG.Debug("Logging in admin in LoginAdmin() function...")

	var req dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LOG.Error(err.Error())
		return
	}

	var admin models.Admins

	db := helper.GetDB()

	if err := db.Where("email = ?", req.Email).First(&admin).Error; err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		logger.LOG.Error(err.Error())
		return
	}

	if !helper.CheckPasswordHash(req.Password, admin.Password) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := helper.GenerateJWT(admin)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		logger.LOG.Error(err.Error())
		return
	}

	json.NewEncoder(w).Encode(dto.LoginResponse{Token: token})
	logger.LOG.Debug(fmt.Sprintf("User %s successfully has successfully logged in!", admin.Name))
}

func (h *ServiceHandler) GetAdminById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var admin models.Admins
	var response dto.AdminResponseRequest

	logger.LOG.Debug(fmt.Sprintf("Getting admin by ID:%d in GetAdminById() function...", admin.ID))

	result := h.Auth.DB.Preload("CarsCreated").Preload("MotorcyclesCreated").First(&admin, r.PathValue("id"))

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("No admin with ID %s found", r.PathValue("id")),
		})
		return
	}

	response = dto.AdminResponseRequest{
		ID:                 admin.ID,
		Name:               admin.Name,
		Email:              admin.Email,
		CarsCreated:        admin.CarsCreated,
		MotorcyclesCreated: admin.MotorcyclesCreated,
	}

	json.NewEncoder(w).Encode(response)
	logger.LOG.Debug(fmt.Sprintf("Admin ID %d is found!", admin.ID))
}
