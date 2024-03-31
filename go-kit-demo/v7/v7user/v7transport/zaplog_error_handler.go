package v7transport

import (
	"context"
	"fmt"

	"v7/v7user/v7service"

	"go.uber.org/zap"
)

// LogErrorHandler ...
type LogErrorHandler struct {
	logger *zap.Logger
}

// NewZapLogErrorHandler ...
func NewZapLogErrorHandler(logger *zap.Logger) *LogErrorHandler {
	return &LogErrorHandler{
		logger: logger,
	}
}

// Handle ...
func (h *LogErrorHandler) Handle(ctx context.Context, err error) {
	h.logger.Warn(fmt.Sprint(ctx.Value(v7service.ContextReqUUID)), zap.Error(err))
}
