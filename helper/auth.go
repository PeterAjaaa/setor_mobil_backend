package helper

import (
	"context"
	"fmt"
	"time"

	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/PeterAjaaa/setor_mobil_backend/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userIDKey struct{}

var jwtKey []byte

func SetUserID(ctx context.Context, id any) context.Context {
	logger.LOG.Debug("Entering SetUserID() function")
	return context.WithValue(ctx, userIDKey{}, id)
}

func GetUserID(ctx context.Context) (uint, bool) {
	logger.LOG.Debug("Entering GetUserID() function")

	v := ctx.Value(userIDKey{})
	if v == nil {
		return 0, false
	}

	f, ok := v.(float64) // JWT numeric values are float64
	if !ok {
		return 0, false
	}

	return uint(f), true
}

func GenerateJWT[T models.JwtClaims](actor T) (string, error) {
	logger.LOG.Debug(fmt.Sprintf("Generating JWT for %s", actor.GetName()))

	key, err := ReadEnvIfExists("JWT_KEY")

	if err != nil {
		logger.LOG.Error(err.Error())
		return "", err
	}

	jwtKey = []byte(key)

	claims := jwt.MapClaims{
		"user_id": actor.GetID(),
		"email":   actor.GetEmail(),
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
