package logger

import (
	"context"
	reqContext "github.com/fdistorted/gokeeper/context"
	"go.uber.org/zap"
)

type RequestIdType string

const requestIDKey = "request_id"

var logger *zap.Logger

func Get() *zap.Logger {
	return logger
}

func Load() (err error) {
	logger, err = zap.NewProduction() // todo add log pattern here
	return err
}

func WithCtxValue(ctx context.Context) *zap.Logger {
	return logger.With(zapFieldsFromContext(ctx)...)
}

func zapFieldsFromContext(ctx context.Context) []zap.Field {
	return []zap.Field{
		zap.String(requestIDKey, reqContext.GetRequestID(ctx)),
	}
}
