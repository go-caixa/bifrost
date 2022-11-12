package logger

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CorrelationKey string

var (
	CorrelationIDKey      CorrelationKey = "correlation-id"
	correlationIDNotFound                = uuid.Nil.String()
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339})
}

func NewCtx(ctx context.Context) context.Context {
	correlationID := uuid.NewString()
	return context.WithValue(ctx, CorrelationIDKey, correlationID)
}

func Infof(ctx context.Context, msg string, args ...any) {
	logWithField(ctx, nil).Infof(msg, args...)
}

func Fatalf(ctx context.Context, err error, msg string, args ...any) {
	logWithField(ctx, err).Fatalf(msg, args...)
}

func Errorf(ctx context.Context, err error, msg string, args ...any) {
	logWithField(ctx, err).Errorf(msg, args...)
}

func logWithField(ctx context.Context, err error) *logrus.Entry {
	correlationID := getCorrelationID(ctx)
	return logrus.WithFields(logrus.Fields{
		"correlationId": correlationID,
		"error":         err,
	})
}

func getCorrelationID(ctx context.Context) string {
	v, ok := ctx.Value(CorrelationIDKey).(string)
	if !ok {
		return correlationIDNotFound
	}

	return v
}
