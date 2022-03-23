package http

import (
	"context"
	"encoding/json"
	"github.com/BaytoorJr/sso/src/transport"
	"net/http"
)

func createUserDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func addUserFieldsDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.AddUserFieldsRequest

	req.Login = r.Header.Get("Login")
	req.Password = r.Header.Get("Password")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func getUserDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetUserRequest

	req.Login = r.Header.Get("Login")
	req.Password = r.Header.Get("Password")

	return req, nil
}
