package router

import (
	"net/http"
	"os"

	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/middleware"
	"github.com/PeterAjaaa/setor_mobil_backend/services"
)

var NewMux *http.ServeMux

func Router() {
	logger.LOG.Info("Setting up routes...")

	key, err := helper.ReadEnvIfExists("JWT_KEY")
	if err != nil {
		logger.LOG.Error(err.Error())
		os.Exit(1)
	}

	authHandler := services.ServiceHandler{Auth: &middleware.AuthHandler{JwtKey: []byte(key), DB: helper.GetDB()}}
	NewMux = http.NewServeMux()

	NewMux.HandleFunc("/register/admin", services.RegisterAdmin)
	NewMux.HandleFunc("/login/admin", services.LoginAdmin)
	NewMux.Handle("/admins/{id}", authHandler.Auth.AuthMiddleware(http.HandlerFunc(authHandler.GetAdminById)))

	NewMux.HandleFunc("/register", services.RegisterUser)
	NewMux.HandleFunc("/login", services.LoginUser)
	NewMux.Handle("/users/{id}", authHandler.Auth.AuthMiddleware(http.HandlerFunc(authHandler.GetUserById)))

	NewMux.Handle("/cars", authHandler.Auth.AuthMiddleware(http.HandlerFunc(authHandler.GetAllCars)))
	NewMux.Handle("/cars/{id}", authHandler.Auth.AuthMiddleware(http.HandlerFunc(authHandler.GetCarById)))
	NewMux.Handle("/cars/create", authHandler.Auth.AuthMiddleware(http.HandlerFunc(authHandler.CreateNewCar)))
}
