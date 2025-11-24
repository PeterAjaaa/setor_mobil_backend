package seeder

import (
	"time"

	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

var adminSeed = []models.Admins{
	{
		Name:        "Dono",
		Email:       "dono@setor.com",
		Password:    "dono1234",
		DateJoined:  time.Date(2025, 11, 11, 2, 13, 00, 00, time.Now().Location()),
		CarsCreated: nil,
	},
	{
		Name:        "Kasino",
		Email:       "kasino@setor.com",
		Password:    "kasino1234",
		DateJoined:  time.Date(2025, 11, 11, 2, 14, 00, 00, time.Now().Location()),
		CarsCreated: nil,
	},
	{
		Name:        "Indro",
		Email:       "indro@setor.com",
		Password:    "indro1234",
		DateJoined:  time.Date(2025, 11, 11, 2, 14, 00, 00, time.Now().Location()),
		CarsCreated: nil,
	},
}
