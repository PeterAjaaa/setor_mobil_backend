package seeder

import "github.com/PeterAjaaa/setor_mobil_backend/helper"

func SeedDatabase() {
	db := helper.GetDB()
	seedUserTable(db)
}
