package middleware

import (
	"context"
	"time"

	"github.com/BaytoorJr/sso/src/service"
	"github.com/BaytoorJr/sso/src/transport"
	"github.com/go-kit/kit/log"
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

func (lm *loggingMiddleware) CreateUser(ctx context.Context, req *transport.CreateUserRequest) (_ *transport.CreateUserResponse, err error) {
	defer lm.logging(time.Now(), "CreateUser", err)
	return lm.next.CreateUser(ctx, req)
}

func (lm *loggingMiddleware) AddUserFields(ctx context.Context, req *transport.AddUserFieldsRequest) (_ *transport.AddUserFieldsResponse, err error) {
	defer lm.logging(time.Now(), "AddUserFields", err)
	return lm.next.AddUserFields(ctx, req)
}

func (lm *loggingMiddleware) GetUser(ctx context.Context, req *transport.GetUserRequest) (_ *transport.GetUserResponse, err error) {
	defer lm.logging(time.Now(), "GetUser", err)
	return lm.next.GetUser(ctx, req)
}

func (lm *loggingMiddleware) DeleteUser(ctx context.Context, req *transport.DeleteUserRequest) (_ *transport.DeleteUserResponse, err error) {
	defer lm.logging(time.Now(), "DeleteUser", err)
	return lm.next.DeleteUser(ctx, req)
}