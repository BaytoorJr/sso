package repository

import (
	"context"

	"github.com/BaytoorJr/sso/src/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, login string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, ID string) error

	AddProfileFields(ctx context.Context, user *domain.User) error
}
