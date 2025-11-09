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
	logger.LOG.Info(fmt.Sprintf("Seeding %s table...", tableName))

	txErr := db.Transaction(func(tx *gorm.DB) error {
		for _, user := range data {
			if err := tx.Create(&user).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if txErr != nil {
		logger.LOG.Warn(fmt.Sprintf("Error seeding data to table %s, transaction rolled back. ERROR: %s", tableName, txErr))
	} else {
		logger.LOG.Info(fmt.Sprintf("Successfully seeded %s table data", tableName))
	}
}
