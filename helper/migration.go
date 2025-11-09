package helper

import (
	"fmt"
	"os"

	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
)

func MigrateTable[T models.Model](target T) {
	tableName := target.TableName()
	logger.LOG.Info(fmt.Sprintf("Starting %s table migration process", tableName))

	db := GetDB()
	dbName, err := ReadEnvIfExists("DB_NAME")

	if db == nil {
		logger.LOG.Error(fmt.Sprintf("Connection to database %s to migrate %s table failed. Exiting...", dbName, tableName))
		os.Exit(1)
	}

	if err != nil {
		logger.LOG.Error(err.Error())
		os.Exit(1)
	}

	logger.LOG.Info(fmt.Sprintf("Migrating %s table to %s database", tableName, dbName))
	err = db.AutoMigrate(&target)

	if err != nil {
		logger.LOG.Error(err.Error())
		os.Exit(1)
	}

}
