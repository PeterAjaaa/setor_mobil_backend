package seeder

import (
	"time"

	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

var orderSeed = []models.Orders{
	{
		Duration:       5,
		Rating:         5,
		PickupTime:     time.Now(),
		PickupLocation: "Bandung",
		Price:          300000,
		Status:         "Active",
		CarID:          helper.UintPtr(1),
		UserID:         1,
	},
	{
		Duration:       10,
		Rating:         4,
		PickupTime:     time.Now(),
		PickupLocation: "Jakarta",
		Price:          150000,
		Status:         "Completed",
		MotorcycleID:   helper.UintPtr(2),
		UserID:         2,
	},
}
