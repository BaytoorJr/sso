package middleware

import (
	"github.com/BaytoorJr/sso/src/service"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingMiddleware struct {
	next   service.MainService
	logger log.Logger
}

func (lm *loggingMiddleware) logging(begin time.Time, method string, err error) {
	_ = lm.logger.Log("method", method, "took", time.Since(begin), "err", err)
}

func NewLoggingMiddleware(logger log.Logger) Middleware {
	return func(service service.MainService) service.MainService {
		return &loggingMiddleware{
			next:   service,
			logger: logger,
		}
	}
}
