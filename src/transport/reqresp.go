package transport

import "github.com/BaytoorJr/sso/src/domain"

type (
	CreateUserRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
	}
)

type (
	AddUserFieldsRequest struct {
		Login    string            `json:"-"`
		Password string            `json:"-"`
		Data     map[string]string `json:"data"`
	}
	AddUserFieldsResponse struct {
	}
)

type (
	GetUserRequest struct {
		Login    string `json:"-"`
		Password string `json:"password"`
	}
	GetUserResponse struct {
		domain.User
	}
)