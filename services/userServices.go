package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PeterAjaaa/setor_mobil_backend/dto"
	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var users []models.User

	db := helper.GetDB()
	result := db.Find(&users)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "No users found",
		})
		return
	}

	var response []dto.UserResponseRequest
	for _, user := range users {
		response = append(response, dto.UserResponseRequest{
			ID:    fmt.Sprintf("%d", user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	json.NewEncoder(w).Encode(response)

}
