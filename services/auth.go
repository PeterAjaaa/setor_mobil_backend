package services

import (
	"fmt"
	"time"

	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey []byte

func GenerateJWT(user models.User) (string, error) {
	logger.LOG.Debug(fmt.Sprintf("Generating JWT for user %s", user.Name))

	key, err := helper.ReadEnvIfExists("JWT_KEY")

	if err != nil {
		logger.LOG.Error(err.Error())
		return "", err
	}

	jwtKey = []byte(key)

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func HashPassword(password string) (string, error) {
	logger.LOG.Debug("Hashing password...")

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	logger.LOG.Debug("Comparing password hashes...")

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
