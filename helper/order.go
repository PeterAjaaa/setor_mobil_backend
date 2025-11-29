package helper

import (
	"fmt"

	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
	"gorm.io/gorm"
)

func MapOrderStatusToVehicleStatus(orderStatus string) string {
	switch orderStatus {
	case "Active":
		return "Rented"
	case "Pending":
		return "Pending"
	case "Completed", "Cancelled":
		return "Available"
	default:
		return "Available"
	}
}

func UpdateVehicleStatus(tx *gorm.DB, order *models.Orders, status string) error {
	if order.CarID != nil && *order.CarID > 0 {
		if err := tx.Model(&models.Cars{}).Where("id = ?", *order.CarID).Update("status", status).Error; err != nil {
			logger.LOG.Error(fmt.Sprintf("Failed to update car status: %v", err))
			return err
		}
		logger.LOG.Debug(fmt.Sprintf("Car #%d status updated to %s", *order.CarID, status))
	}

	if order.MotorcycleID != nil && *order.MotorcycleID > 0 {
		if err := tx.Model(&models.Motorcycles{}).Where("id = ?", *order.MotorcycleID).Update("status", status).Error; err != nil {
			logger.LOG.Error(fmt.Sprintf("Failed to update motorcycle status: %v", err))
			return err
		}
		logger.LOG.Debug(fmt.Sprintf("Motorcycle #%d status updated to %s", *order.MotorcycleID, status))
	}

	return nil
}
