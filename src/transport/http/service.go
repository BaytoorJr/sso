package http

import (
	httpencoders "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/transport/http"
	"github.com/BaytoorJr/sso/src/middleware"
	"github.com/go-kit/kit/log"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func NewHTTPService(svcEndpoints *middleware.Endpoints, options []kithttp.ServerOption, logger log.Logger) http.Handler {
	errorEncoder := kithttp.ServerErrorEncoder(
		httpencoders.EncodeErrorResponse,
	)

	errorLogger := kithttp.ServerErrorHandler(
		kittransport.NewLogErrorHandler(logger),
	)

	options = append(options, errorEncoder, errorLogger)

	return initializeRoutes(svcEndpoints, options)
}
