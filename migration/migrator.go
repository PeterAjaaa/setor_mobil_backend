package migration

import (
	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

func DoMigrateTables() {
	helper.MigrateTable(&models.Users{})
	helper.MigrateTable(&models.Cars{})
	helper.MigrateTable(&models.Admins{})
	helper.MigrateTable(&models.Motorcycles{})
	helper.MigrateTable(&models.Orders{})
}
