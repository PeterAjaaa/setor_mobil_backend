package helper

import (
	"fmt"
	"os"

	"github.com/PeterAjaaa/setor_mobil_backend/logger"
)

func ReadEnvIfExists(key string) (string, error) {
	logger.LOG.Debug("Entering ReadEnvIfExists function")
	value, found := os.LookupEnv(key)
	logger.LOG.Debug(fmt.Sprintf("Reading %s key from the env", key))

	if found {
		logger.LOG.Debug(fmt.Sprintf("Found %s key from the env", key))
		return value, nil
	} else {
		return "", fmt.Errorf("env key didn't exist, make sure key %s exists", key)
	}
}
