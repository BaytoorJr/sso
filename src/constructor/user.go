package constructor

import (
	"github.com/BaytoorJr/sso/src/middleware"
	"github.com/BaytoorJr/sso/src/repository"
	"github.com/BaytoorJr/sso/src/service"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func NewUserService(mainRepo repository.MainRepo, logger log.Logger) service.MainService {
	svc := service.NewService(mainRepo, logger)
	svc = middleware.NewLoggingMiddleware(logger)(svc)
	svc = middleware.NewInstrumentingMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "sso",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"}),
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "sso",
			Name:      "error_count",
			Help:      "Number of error requests received.",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "sso",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"}),
	)(svc)

	return svc
}
