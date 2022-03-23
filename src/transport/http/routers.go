package http

import (
	encoders "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/transport/http"
	"github.com/BaytoorJr/sso/src/middleware"
	kitHTTP "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func initializeRoutes(endpoints *middleware.Endpoints, options []kitHTTP.ServerOption) *mux.Router {

	createUser := kitHTTP.NewServer(
		endpoints.CreateUser,
		createUserDecoder,
		encoders.EncodeResponse,
		options...,
	)

	addUserFields := kitHTTP.NewServer(
		endpoints.AddUserFields,
		addUserFieldsDecoder,
		encoders.EncodeResponse,
		options...,
	)

	getUser := kitHTTP.NewServer(
		endpoints.GetUser,
		getUserDecoder,
		encoders.EncodeResponse,
		options...,
	)

	router := mux.NewRouter()

	router.Path("/sso/v1/user").
		Methods("POST").
		Handler(createUser)

	router.Path("/sso/v1/user").
		Methods("PUT").
		Handler(addUserFields)

	router.Path("/sso/v1/user").
		Methods("GET").
		Handler(getUser)

	return router
}
