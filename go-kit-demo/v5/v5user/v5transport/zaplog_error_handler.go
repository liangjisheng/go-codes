package v5transport

import (
	"context"
	"fmt"

	"v5/v5user/v5service"

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
	h.logger.Warn(fmt.Sprint(ctx.Value(v5service.ContextReqUUID)), zap.Error(err))
}
