package helper

import (
	"encoding/json"
	"net/http"

	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

func HttpErrorHelper(w http.ResponseWriter, status int, msg string, data any) {
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(
		models.APIResponse{
			Status:  status,
			Message: msg,
			Data:    data,
		},
	)

}
