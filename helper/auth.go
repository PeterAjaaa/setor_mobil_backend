package helper

import (
	"context"

	"github.com/PeterAjaaa/setor_mobil_backend/logger"
)

type userIDKey struct{}

func SetUserID(ctx context.Context, id any) context.Context {
	logger.LOG.Debug("Entering SetUserID() function")
	return context.WithValue(ctx, userIDKey{}, id)
}

func GetUserID(ctx context.Context) any {
	logger.LOG.Debug("Entering GetUserID() function")
	return ctx.Value(userIDKey{})
}
