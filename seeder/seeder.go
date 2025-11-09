package seeder

import (
	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
)

func SeedDatabase() {
	logger.LOG.Info("Seeding database...")

	db := helper.GetDB()

	helper.SeedTable(db, userSeed)
	helper.SeedTable(db, carSeed)
}
