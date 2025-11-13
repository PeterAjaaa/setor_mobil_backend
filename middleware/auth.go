package middleware

import (
	"net/http"
	"strings"

	"github.com/PeterAjaaa/setor_mobil_backend/helper"
	"github.com/PeterAjaaa/setor_mobil_backend/logger"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	JwtKey []byte
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
			helper.HttpErrorHelper(w, http.StatusUnauthorized, "Missing authorization header", nil)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			return []byte(secret), nil
		}, jwt.WithValidMethods([]string{"HS256"}))

		if err != nil || !token.Valid {
			helper.HttpErrorHelper(w, http.StatusUnauthorized, "Invalid token", nil)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			helper.HttpErrorHelper(w, http.StatusUnauthorized, "Invalid claims", nil)
			return
		}

		ctx := r.Context()
		ctx = helper.SetUserID(ctx, claims["user_id"])
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
