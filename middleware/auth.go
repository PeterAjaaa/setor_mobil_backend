package middleware

import (
	"net/http"
	"strings"

	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthHandler struct {
	JwtKey []byte
	DB     *gorm.DB
}

func (h *AuthHandler) AuthMiddleware(next http.Handler) http.Handler {
	logger.LOG.Debug("Entering AuthMiddleware() function...")

	secret, err := helper.ReadEnvIfExists("JWT_KEY")

	if err != nil {
		logger.LOG.Error(err.Error())
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid claims", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = helper.SetUserID(ctx, claims["user_id"])
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
