package seeder

import (
	"time"

	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

var adminSeed = []models.Admins{
	{
		Name:        "Dono",
		Email:       "dono@admin.com",
		Password:    "dono",
		DateJoined:  time.Date(2025, 11, 11, 2, 13, 00, 00, time.Now().Location()),
		CarsCreated: nil,
	},
	{
		Name:        "Kasino",
		Email:       "kasino@admin.com",
		Password:    "kasino",
		DateJoined:  time.Date(2025, 11, 11, 2, 14, 00, 00, time.Now().Location()),
		CarsCreated: nil,
	},
	{
		Name:        "Indro",
		Email:       "indro@admin.com",
		Password:    "indro",
		DateJoined:  time.Date(2025, 11, 11, 2, 14, 00, 00, time.Now().Location()),
		CarsCreated: nil,
	},
}
