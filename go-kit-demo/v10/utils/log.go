package utils

import (
	"go.uber.org/zap"

	logzap "local.com/log-zap"
)

var logger *zap.Logger

// NewLoggerServer ...
func NewLoggerServer() {
	logger = logzap.NewLogger(
		logzap.SetAppName("go-kit"),
		logzap.SetDevelopment(true),
		logzap.SetLevel(zap.DebugLevel),
	)
}

// GetLogger ...
func GetLogger() *zap.Logger {
	return logger
}
