package middleware

import (
	"context"
	"github.com/BaytoorJr/sso/src/service"
	"github.com/BaytoorJr/sso/src/transport"
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

func (im *instrumentingMiddleware) CreateUser(ctx context.Context, req *transport.CreateUserRequest) (_ *transport.CreateUserResponse, err error) {
	defer im.instrumenting(time.Now(), "CreateUser", err)
	return im.next.CreateUser(ctx, req)
}

func (im *instrumentingMiddleware) AddUserFields(ctx context.Context, req *transport.AddUserFieldsRequest) (_ *transport.AddUserFieldsResponse, err error) {
	defer im.instrumenting(time.Now(), "AddUserFields", err)
	return im.next.AddUserFields(ctx, req)
}

func (im *instrumentingMiddleware) GetUser(ctx context.Context, req *transport.GetUserRequest) (_ *transport.GetUserResponse, err error) {
	defer im.instrumenting(time.Now(), "GetUser", err)
	return im.next.GetUser(ctx, req)
}
