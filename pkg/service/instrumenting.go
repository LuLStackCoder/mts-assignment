package service

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/LuLStackCoder/mts-assignment/pkg/api"
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         Service
}

// HandleUrls ...
func (s *instrumentingMiddleware) HandleUrls(ctx context.Context, urls []string) (data []api.URLData, err error) {
	defer s.recordMetrics("HandleUrls", time.Now(), err)
	return s.svc.HandleUrls(ctx, urls)
}

func (s *instrumentingMiddleware) recordMetrics(method string, startTime time.Time, err error) {
	labels := []string{
		"method", method,
		"error", strconv.FormatBool(err != nil),
	}
	s.reqCount.With(labels...).Add(1)
	s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
}

// NewInstrumentingMiddleware ...
func NewInstrumentingMiddleware(reqCount metrics.Counter, reqDuration metrics.Histogram, svc Service) Service {
	return &instrumentingMiddleware{
		reqCount:    reqCount,
		reqDuration: reqDuration,
		svc:         svc,
	}
}
