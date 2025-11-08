package seeder

import (
	"fmt"
	"time"

	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
	"gorm.io/gorm"
)

func seedUserTable(db *gorm.DB) {
	var users []models.User
	txErr := db.Transaction(func(tx *gorm.DB) error {
		users = append(users, models.User{
			Name:       "Ujang",
			Birthdate:  time.Date(2025, time.November, 11, 8, 0, 0, 0, time.Now().Location()),
			Birthplace: "Bandung",
			Gender:     "Laki-laki",
			Address:    "Jl. Mawar No. 20",
			RT:         002,
			RW:         003,
			Keluarahan: "Kesana",
			Kecamatan:  "Kemari",
			Occupation: "Karyawan",
			Email:      "test@email.tld",
		}, models.User{
			Name:       "Dewi",
			Birthdate:  time.Date(2025, time.November, 5, 3, 0, 0, 0, time.Now().Location()),
			Birthplace: "Jawa Tengah",
			Gender:     "Perempuan",
			Address:    "Jl. Melati No. 50",
			RT:         010,
			RW:         001,
			Keluarahan: "Disini",
			Kecamatan:  "Disana",
			Occupation: "Pelajar",
			Email:      "hello@bye.ok",
		})

		for _, user := range users {
			if err := tx.Create(&user).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if txErr != nil {
		logger.LOG.Warn(fmt.Sprintf("Error seeding data, transaction rolled back. ERROR: %s", txErr))
	} else {
		logger.LOG.Info("Successfully seeded User table data")
	}
}
