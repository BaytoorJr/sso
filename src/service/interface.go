package service

import (
	"context"

	"github.com/BaytoorJr/sso/src/transport"
)

type MainService interface {
	CreateUser(ctx context.Context, req *transport.CreateUserRequest) (*transport.CreateUserResponse, error)
	AddUserFields(ctx context.Context, req *transport.AddUserFieldsRequest) (*transport.AddUserFieldsResponse, error)
	GetUser(ctx context.Context, req *transport.GetUserRequest) (*transport.GetUserResponse, error)
	DeleteUser(ctx context.Context, req *transport.DeleteUserRequest) (*transport.DeleteUserResponse, error)
}
