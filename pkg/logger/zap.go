package logger

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer logger.Sync() // flushes buffer, if any
	return logger, nil
}
