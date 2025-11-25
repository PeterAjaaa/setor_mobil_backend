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

func (h *ServiceHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var user models.Users
	var response dto.UserResponseRequest

	logger.LOG.Debug(fmt.Sprintf("Getting user by ID:%d in GetUserById() function...", user.ID))

	userID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {
		helper.SendHttpResponse(w, http.StatusBadRequest, err.Error(), nil)
		logger.LOG.Error(err.Error())
		return
	}

	result := h.DB.Preload("Orders").Preload("Ratings").First(&user, userID)

	if result.Error != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.SendHttpResponse(w, http.StatusNotFound, fmt.Sprintf("No user with ID %d found", userID), nil)
		return
	}

	response = dto.UserResponseRequest{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	helper.SendHttpResponse(w, http.StatusOK, fmt.Sprintf("User ID %d is found!", user.ID), response)

	logger.LOG.Debug(fmt.Sprintf("User ID %d is found!", user.ID))
}

func (h *ServiceHandler) GetUserCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var count int64

	logger.LOG.Debug("Getting user count in GetTotalUsers() function")

	result := h.DB.Model(&models.Users{}).Count(&count)

	if result.Error != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		logger.LOG.Error(result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		helper.SendHttpResponse(w, http.StatusNotFound, "No users found", nil)
		return
	}

	helper.SendHttpResponse(w, http.StatusOK, fmt.Sprintf("Found %d users", count), count)

	logger.LOG.Debug(fmt.Sprintf("Found %d users", count))
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.SendHttpResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Registering user in RegisterUser() function...")

	var req dto.UserCreationRequest

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

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, "Failed to hash password", nil)
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
		helper.SendHttpResponse(w, http.StatusInternalServerError, txErr.Error(), nil)
		return
	} else {
		logger.LOG.Info("Successfully registered a new user with transaction!")
	}

	response := dto.UserResponseRequest{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	helper.SendHttpResponse(w, http.StatusCreated, "Successfully registered a new user!", response)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.SendHttpResponse(w, http.StatusInternalServerError, "Method not allowed", nil)
		return
	}

	logger.LOG.Debug("Logging in user in LoginUser() function...")

	var req dto.LoginRequest

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

	var user models.Users

	db := helper.GetDB()
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		helper.SendHttpResponse(w, http.StatusUnauthorized, "Invalid email or password", nil)
		logger.LOG.Error(err.Error())
		return
	}

	if !helper.CheckPasswordHash(req.Password, user.Password) {
		helper.SendHttpResponse(w, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	token, err := helper.GenerateJWT(user)
	if err != nil {
		helper.SendHttpResponse(w, http.StatusInternalServerError, "Failed to create token", nil)
		logger.LOG.Error(err.Error())
		return
	}

	response := dto.LoginResponse{Token: token}

	helper.SendHttpResponse(w, http.StatusOK, fmt.Sprintf("User %s successfully has successfully logged in!", user.Name), response)

	logger.LOG.Debug(fmt.Sprintf("User %s successfully has successfully logged in!", user.Name))
}
