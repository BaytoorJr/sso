package service

import (
	"context"
	"errors"

	"github.com/BaytoorJr/sso/src/domain"
	"github.com/BaytoorJr/sso/src/transport"
)

func (s *service) CreateUser(ctx context.Context, req *transport.CreateUserRequest) (*transport.CreateUserResponse, error) {
	var user domain.User
	user.Init(req.Login, req.Password)

	err := s.store.Users().CreateUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &transport.CreateUserResponse{
		User: user,
	}, nil
}

func (s *service) AddUserFields(ctx context.Context, req *transport.AddUserFieldsRequest) (*transport.AddUserFieldsResponse, error) {
	user, err := s.store.Users().GetUser(ctx, req.Login)
	if err != nil {
		return nil, err
	}

	if user.Password != req.Password {
		return nil, errors.New("incorrect password")
	}

	user.Data = req.Data

	err = s.store.Users().AddProfileFields(ctx, user)
	if err != nil {
		return nil, err
	}

	return &transport.AddUserFieldsResponse{}, nil
}

func (s *service) GetUser(ctx context.Context, req *transport.GetUserRequest) (*transport.GetUserResponse, error) {
	user, err := s.store.Users().GetUser(ctx, req.Login)
	if err != nil {
		return nil, err
	}

	if user.Password != req.Password {
		return nil, errors.New("access denied")
	}

	return &transport.GetUserResponse{
		User: *user,
	}, nil
}

func (s *service) DeleteUser(ctx context.Context, req *transport.DeleteUserRequest) (*transport.DeleteUserResponse, error) {
	user, err := s.store.Users().GetUser(ctx, req.Login)
	if err != nil {
		return nil, err
	}

	if user.Password != req.Password {
		return nil, errors.New("access denied")
	}

	err = s.store.Users().DeleteUser(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &transport.DeleteUserResponse{}, nil
}