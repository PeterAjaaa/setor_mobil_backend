package seeder

import (
	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
)

func SeedDatabase() {
	logger.LOG.Info("Seeding database...")

	db := helper.GetDB()

	// SEED ORDER:
	// USER -> ADMIN -> CAR / MOTORCYCLE -> ORDER -> RATING
	helper.SeedTable(db, userSeed)
	helper.SeedTable(db, adminSeed)
	helper.SeedTable(db, carSeed)
	helper.SeedTable(db, motorcycleSeed)
	helper.SeedTable(db, orderSeed)
	helper.SeedTable(db, ratingSeed)
}
