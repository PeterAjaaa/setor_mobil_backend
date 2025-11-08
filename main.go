package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/migration"
	"github.com/PeterAjaaa/setor_mobil_backend/router"
	"github.com/PeterAjaaa/setor_mobil_backend/seeder"
)

func main() {
	logger.SetLogLevel()
	logger.LOG.Debug("Calling main() function")

	helper.InitDB()

	migration.MigrateUserTable()
	seeder.SeedDatabase()
	router.Router()

	port, err := helper.ReadEnvIfExists("HOST_PORT")

	if err != nil {
		logger.LOG.Error(err.Error())
		os.Exit(1)
	}

	logger.LOG.Info(fmt.Sprintf("Serving API at %s", port))
	http.ListenAndServe(":"+port, nil)
}
