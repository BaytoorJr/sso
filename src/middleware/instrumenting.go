package middleware

import (
	"github.com/BaytoorJr/sso/src/service"
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingMiddleware struct {
	next           service.MainService
	requestCount   metrics.Counter
	requestError   metrics.Counter
	requestLatency metrics.Histogram
}

func (im *instrumentingMiddleware) instrumenting(begin time.Time, method string, err error) {
	im.requestCount.With("method", method).Add(1)
	if err != nil {
		im.requestError.With("method", method).Add(1)
	}

	im.requestLatency.With("method", method).Observe(time.Since(begin).Seconds())
}

func NewInstrumentingMiddleware(counter, errCounter metrics.Counter, latency metrics.Histogram) Middleware {
	return func(next service.MainService) service.MainService {
		return &instrumentingMiddleware{
			next:           next,
			requestLatency: latency,
			requestError:   errCounter,
			requestCount:   counter,
		}
	}
}
