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

func (h *ServiceHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user models.Users
	var response dto.UserResponseRequest

	logger.LOG.Debug(fmt.Sprintf("Getting user by ID:%d in GetUserById() function...", user.ID))

	result := h.Auth.DB.First(&user, r.PathValue("id"))

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("No user with ID %s found", r.PathValue("id")),
		})
		return
	}

	response = dto.UserResponseRequest{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	json.NewEncoder(w).Encode(response)
	logger.LOG.Debug(fmt.Sprintf("User ID %d is found!", user.ID))
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	logger.LOG.Debug("Registering user in RegisterUser() function...")

	var req dto.UserCreationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LOG.Error(err.Error())
		return
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		logger.LOG.Error(err.Error())
		return
	}

	user := models.Users{
		Name:       req.Name,
		Email:      req.Email,
		Password:   hashedPassword,
		Birthdate:  req.Birthdate,
		Birthplace: req.Birthplace,
		Gender:     string(req.Gender),
		Address:    req.Address,
		RT:         req.RT,
		RW:         req.RW,
		Keluarahan: req.Keluarahan,
		Kecamatan:  req.Kecamatan,
		Occupation: req.Occupation,
	}

	db := helper.GetDB()
	txErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
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

	response := dto.UserResponseRequest{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	json.NewEncoder(w).Encode(response)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	logger.LOG.Debug("Logging in user in LoginUser() function...")

	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.LOG.Error(err.Error())
		return
	}

	var user models.Users
	db := helper.GetDB()
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		logger.LOG.Error(err.Error())
		return
	}

	if !CheckPasswordHash(req.Password, user.Password) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(user)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		logger.LOG.Error(err.Error())
		return
	}

	json.NewEncoder(w).Encode(dto.LoginResponse{Token: token})
	logger.LOG.Debug(fmt.Sprintf("User %s successfully has successfully logged in!", user.Name))
}
