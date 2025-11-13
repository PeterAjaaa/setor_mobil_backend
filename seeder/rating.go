package seeder

import (
	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

var ratingSeed = []models.Ratings{
	{
		Rating:  5,
		UserID:  1,
		OrderID: 1,
		CarID:   helper.UintPtr(1),
	},
	{
		Rating:       4,
		UserID:       2,
		OrderID:      2,
		MotorcycleID: helper.UintPtr(2),
	},
}
