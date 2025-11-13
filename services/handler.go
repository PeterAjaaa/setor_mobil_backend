package services

import (
	"github.com/PeterAjaaa/setor_mobil_backend/middleware"
	"gorm.io/gorm"
)

type ServiceHandler struct {
	Auth *middleware.AuthHandler
	DB   *gorm.DB
}
