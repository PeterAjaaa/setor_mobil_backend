package helper

import (
	"fmt"

	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
	"gorm.io/gorm"
)

func SeedTable[T models.Model](db *gorm.DB, data []T) {
	if len(data) == 0 {
		logger.LOG.Error("seedTable() called with 0 seeder data. Skipping...")
		return
	}

	tableName := data[0].TableName()
	logger.LOG.Info(fmt.Sprintf("Checking %s table for existing data...", tableName))

	var count int64
	if err := db.Model(data[0]).Count(&count).Error; err != nil {
		logger.LOG.Error(fmt.Sprintf("Error checking %s table count: %s", tableName, err))
		return
	}

	if count > 0 {
		logger.LOG.Info(fmt.Sprintf("Table %s already contains %d records. Skipping seeding.", tableName, count))
		return
	}

	logger.LOG.Info(fmt.Sprintf("Seeding %s table (empty)...", tableName))

	txErr := db.Transaction(func(tx *gorm.DB) error {
		for _, item := range data {
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if txErr != nil {
		logger.LOG.Warn(fmt.Sprintf("Error seeding data to table %s, transaction rolled back. ERROR: %s", tableName, txErr))
	} else {
		logger.LOG.Info(fmt.Sprintf("Successfully seeded %d records to %s table", len(data), tableName))
	}
}
