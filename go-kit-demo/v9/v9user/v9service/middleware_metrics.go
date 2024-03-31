package v9service

import (
	"context"
	"time"

	"v9/v9user/pb"

	"github.com/go-kit/kit/metrics"
)

type metricsMiddlewareServer struct {
	next      Service
	counter   metrics.Counter
	histogram metrics.Histogram
}

// NewMetricsMiddlewareServer ...
func NewMetricsMiddlewareServer(counter metrics.Counter, histogram metrics.Histogram) NewMiddlewareServer {
	return func(service Service) Service {
		return metricsMiddlewareServer{
			next:      service,
			counter:   counter,
			histogram: histogram,
		}
	}
}

func (m metricsMiddlewareServer) Login(ctx context.Context, in *pb.Login) (out *pb.LoginAck, err error) {
	defer func(start time.Time) {
		method := []string{"method", "login"}
		m.counter.With(method...).Add(1)
		m.histogram.With(method...).Observe(time.Since(start).Seconds())
	}(time.Now())
	out, err = m.next.Login(ctx, in)
	return
}
