package seeder

import (
	"time"

	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

var userSeed = []models.Users{
	{
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
	}, {
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
	},
}
