package middleware

import (
	"context"

	"github.com/BaytoorJr/sso/src/service"
	"github.com/BaytoorJr/sso/src/transport"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser    endpoint.Endpoint
	AddUserFields endpoint.Endpoint
	GetUser       endpoint.Endpoint
	DeleteUser    endpoint.Endpoint
}

func MakeEndpoints(s service.MainService) *Endpoints {
	return &Endpoints{
		CreateUser:    makeCreateUserEndpoint(s),
		AddUserFields: makeAddUserFieldsEndpoint(s),
		GetUser:       makeGetUserEndpoint(s),
		DeleteUser:    makeDeleteUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s service.MainService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.CreateUserRequest)
		return s.CreateUser(ctx, &req)
	}
}

func makeAddUserFieldsEndpoint(s service.MainService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.AddUserFieldsRequest)
		return s.AddUserFields(ctx, &req)
	}
}

func makeGetUserEndpoint(s service.MainService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.GetUserRequest)
		return s.GetUser(ctx, &req)
	}
}

func makeDeleteUserEndpoint(s service.MainService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.DeleteUserRequest)
		return s.DeleteUser(ctx, &req)
	}
}
