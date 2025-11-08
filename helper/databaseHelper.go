package helper

import (
	"fmt"
	"os"
	"sync"

	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func InitDB() {
	once.Do(func() {
		dbUser, err1 := ReadEnvIfExists("DB_USER")
		dbPass, err2 := ReadEnvIfExists("DB_PASSWORD")
		dbHost, err3 := ReadEnvIfExists("DB_HOST")
		dbPort, err4 := ReadEnvIfExists("DB_PORT")
		dbName, err5 := ReadEnvIfExists("DB_NAME")

		if err := firstErr(err1, err2, err3, err4, err5); err != nil {
			logger.LOG.Error(fmt.Sprintf("Failed to read DB env: %v", err))
			return
		}

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPass, dbHost, dbPort, dbName)

		logger.LOG.Info(fmt.Sprintf("Connecting to MySQL database: %s", dbName))

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.LOG.Error(fmt.Sprintf("Failed to connect to DB: %v", err))
			os.Exit(1)
		}

		dbInstance = db
		logger.LOG.Info("Database connection established")
	})
}

func GetDB() *gorm.DB {
	if dbInstance == nil {
		logger.LOG.Warn("Database not initialized — calling InitDB() automatically")
		InitDB()
	}
	return dbInstance
}

func firstErr(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
