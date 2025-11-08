package router

import (
	"net/http"

	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/services"
)

func Router() {
	logger.LOG.Debug("Calling main() function")
	http.HandleFunc("/users", services.GetUsers)
}
