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

	serviceHandler := services.ServiceHandler{Auth: &middleware.AuthHandler{JwtKey: []byte(key)}, DB: helper.GetDB()}
	NewMux = http.NewServeMux()

	NewMux.HandleFunc("/register/admin", services.RegisterAdmin)
	NewMux.HandleFunc("/login/admin", services.LoginAdmin)
	NewMux.Handle("/admins/{id}", serviceHandler.Auth.AuthMiddleware(http.HandlerFunc(serviceHandler.GetAdminById)))

	NewMux.HandleFunc("/register", services.RegisterUser)
	NewMux.HandleFunc("/login", services.LoginUser)
	NewMux.Handle("/users/{id}", serviceHandler.Auth.AuthMiddleware(http.HandlerFunc(serviceHandler.GetUserById)))

	NewMux.Handle("/cars", serviceHandler.Auth.AuthMiddleware(http.HandlerFunc(serviceHandler.GetAllCars)))
	NewMux.Handle("/cars/{id}", serviceHandler.Auth.AuthMiddleware(http.HandlerFunc(serviceHandler.GetCarById)))
	NewMux.Handle("/cars/create", serviceHandler.Auth.AuthMiddleware(http.HandlerFunc(serviceHandler.CreateNewCar)))

	NewMux.Handle("/motorcycles", serviceHandler.Auth.AuthMiddleware(http.HandlerFunc(serviceHandler.GetAllMotorcycles)))
	NewMux.Handle("/motorcycles/{id}", serviceHandler.Auth.AuthMiddleware(http.HandlerFunc(serviceHandler.GetMotorcyleById)))
	NewMux.Handle("/motorcycles/create", serviceHandler.Auth.AuthMiddleware(http.HandlerFunc(serviceHandler.CreateNewMotorcycle)))

	NewMux.Handle("/orders/create", serviceHandler.Auth.AuthMiddleware(http.HandlerFunc(serviceHandler.CreateNewOrder)))
	NewMux.Handle("/orders/{id}", serviceHandler.Auth.AuthMiddleware(http.HandlerFunc(serviceHandler.GetOrderById)))
	NewMux.Handle("/orders/user/{id}", serviceHandler.Auth.AuthMiddleware(http.HandlerFunc(serviceHandler.GetAllOrdersByUserId)))

}
