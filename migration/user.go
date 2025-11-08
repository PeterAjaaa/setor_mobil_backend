package migration

import (
	"fmt"
	"os"

	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

func MigrateUserTable() {
	logger.LOG.Info("Starting user table migration process")
	db := helper.GetDB()
	dbName, err := helper.ReadEnvIfExists("DB_NAME")

	if db == nil {
		logger.LOG.Error(fmt.Sprintf("Connection to database %s to migrate User table failed. Exiting...", dbName))
		os.Exit(1)
	}

	if err != nil {
		logger.LOG.Error(err.Error())
		os.Exit(1)
	}

	logger.LOG.Info(fmt.Sprintf("Migrating User table to %s database", dbName))
	err = db.AutoMigrate(&models.User{})

	if err != nil {
		logger.LOG.Error(err.Error())
		os.Exit(1)
	}

}
