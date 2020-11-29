package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/LuLStackCoder/mts-assignment/pkg/api"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger log.Logger
	svc    Service
}

// HandleUrls ...
func (s *loggingMiddleware) HandleUrls(ctx context.Context, urls []string) (data []api.URLData, errorFlag bool, errorText string, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "HandleUrls",
			"urls", urls,
			"data", data,
			"errorFlag", errorFlag,
			"errorText", errorText,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.HandleUrls(ctx, urls)
}

func (s *loggingMiddleware) wrap(err error) log.Logger {
	lvl := level.Debug
	if err != nil {
		lvl = level.Error
	}
	return lvl(s.logger)
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger log.Logger, svc Service) Service {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
