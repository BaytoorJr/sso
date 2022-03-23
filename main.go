package main

import (
	"context"
	"flag"
	"fmt"
	httpencoders "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/transport/http"
	corsutil "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/utils/cors"
	healthcheckutil "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/utils/healthcheck"
	"github.com/BaytoorJr/sso/src/config"
	"github.com/BaytoorJr/sso/src/constructor"
	"github.com/BaytoorJr/sso/src/middleware"
	"github.com/BaytoorJr/sso/src/repository/postgres"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	postgresqlConn "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/databases/postgresql"
	liblogger "git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/utils/logger"
	userHTTP "github.com/BaytoorJr/sso/src/transport/http"
)

func main() {
	// main context
	ctx := context.Background()

	// parse flags
	httpPort := flag.String("http.port", ":8080", "HTTP listen address")
	flag.Parse()

	// init structured logger for service
	logger := liblogger.NewServiceLogger("sso")

	// init configs
	err := config.InitConfigs()
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(1)
	}

	// init postgresql connection
	postgresConn, err := postgresqlConn.InitConnect(
		ctx,
		config.MainConfig.DBConfig.PostgresMaxConns,
		config.MainConfig.DBConfig.PostgresHost,
		config.MainConfig.DBConfig.PostgresPort,
		config.MainConfig.DBConfig.PostgresUser,
		config.MainConfig.DBConfig.PostgresPass,
		config.MainConfig.DBConfig.PostgresName,
	)
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(1)
	}

	// init main repository (data access layer)
	mainRepo, err := postgres.NewStore(postgresConn, logger)
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(1)
	}

	// init main service (business logic layer)
	mainSvc := constructor.NewUserService(mainRepo, logger)

	// init endpoints (endpoints layer)
	endpoints := middleware.MakeEndpoints(mainSvc)

	// init HTTP handler (transport layer)
	serverOptions := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(httpencoders.EncodeErrorResponse),
	}
	mainHandler := userHTTP.NewHTTPService(endpoints, serverOptions, logger)

	// add routes, prometheus and health check handler
	http.Handle("/sso/v1/", corsutil.CORS(mainHandler))
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/check", healthcheckutil.HealthCheck)

	// init errors chan
	errs := make(chan error)

	// make chan for syscall
	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// init HTTP server
	go func() {
		fmt.Print()
		_ = level.Info(logger).Log("transport", "HTTP", "port", *httpPort)
		errs <- http.ListenAndServe(*httpPort, nil)
	}()

	defer func() {
		_ = level.Info(logger).Log("msg", "service ended")
	}()

	_ = level.Error(logger).Log("exit", <-errs)
}
