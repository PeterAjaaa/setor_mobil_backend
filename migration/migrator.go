package migration

import (
	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

func DoMigrateTables() {
	helper.MigrateTable(&models.Users{})
}
