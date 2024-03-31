package v6transport

import (
	"context"
	"fmt"

	"v6/v6user/v6service"

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
	h.logger.Warn(fmt.Sprint(ctx.Value(v6service.ContextReqUUID)), zap.Error(err))
}
