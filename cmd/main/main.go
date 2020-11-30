package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/buaazp/fasthttprouter"
	fasthttpprometheus "github.com/flf2ko/fasthttp-prometheus"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"

	"github.com/LuLStackCoder/mts-assignment/pkg/service"
	"github.com/LuLStackCoder/mts-assignment/pkg/service/httperrors"
	"github.com/LuLStackCoder/mts-assignment/pkg/service/httpserver"
	"github.com/LuLStackCoder/mts-assignment/pkg/urlclient"

	"github.com/kelseyhightower/envconfig"
)

var (
	serviceVersion = "dev"
	methodError    = []string{"method", "error"}
)

type configuration struct {
	Debug              bool   `envconfig:"DEBUG" default:"true"`
	Port               string `envconfig:"PORT" default:"8080"`
	MaxRequestBodySize int    `envconfig:"MAX_REQUEST_BODY_SIZE" default:"10485760"` // 10 MB
	ReadBufferSize     int    `envconfig:"READ_BUFFER_SIZE" default:"16384"`
	// 6.сервер не обслуживает больше чем 100 одновременных входящих http-запросов
	MaxSimultaneousConns int `envconfig:"MAX_SIM_CONNS" default:"100"`
	//  5.сервер не принимает запрос если количество url в нем больше 20
	MaxURLs int `envconfig:"MAX_URLS" default:"20"`
	// 7.таймаут на обработку одного входящего запроса - 10000 миллисекунд
	InTimeout time.Duration `envconfig:"IN_TIMEOUT" default:"10s"`
	// 8.таймаут на запрос одного url - 500 миллисекунд
	GetTimeout time.Duration `envconfig:"GET_TIMEOUT" default:"500ms"`

	MetricsNamespace string `envconfig:"METRICS_NAMESPACE" default:"mts"`
	MetricsSubsystem string `envconfig:"METRICS_SUBSYSTEM" default:"mts_assignment"`

	MetricsNameCount    string `envconfig:"METRICS_NAME_COUNT" default:"request_count"`
	MetricsNameDuration string `envconfig:"METRICS_NAME_DURATION" default:"request_duration"`
	MetricsHelpCount    string `envconfig:"METRICS_HELP_COUNT" default:"Request count"`
	MetricsHelpDuration string `envconfig:"METRICS_HELP_DURATION" default:"Request duration"`
}

func main() {
	printVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *printVersion {
		fmt.Println(serviceVersion)
		os.Exit(0)
	}

	// logger
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	_ = level.Info(logger).Log("msg", "initializing", "version", serviceVersion)

	// processing configuration
	var cfg configuration
	if err := envconfig.Process("", &cfg); err != nil {
		_ = level.Error(logger).Log("msg", "failed to load configuration", "err", err)
		os.Exit(1)
	}

	if !cfg.Debug {
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	// client for getting data from urls
	httpClient := http.Client{Timeout: cfg.GetTimeout}

	urlClient := urlclient.NewClient(&httpClient)

	//  5.сервер не принимает запрос если количество url в нем больше 20
	// service creation
	svc := service.NewService(urlClient, cfg.MaxURLs)
	svc = service.NewLoggingMiddleware(logger, svc)
	svc = service.NewInstrumentingMiddleware(
		// expose metrics
		kitprometheus.NewCounterFrom(
			prometheus.CounterOpts{
				Namespace: cfg.MetricsNamespace,
				Subsystem: cfg.MetricsSubsystem,
				Name:      cfg.MetricsNameCount,
				Help:      cfg.MetricsHelpCount,
			},
			methodError,
		),
		kitprometheus.NewSummaryFrom(
			prometheus.SummaryOpts{
				Namespace: cfg.MetricsNamespace,
				Subsystem: cfg.MetricsSubsystem,
				Name:      cfg.MetricsNameDuration,
				Help:      cfg.MetricsHelpDuration,
			},
			methodError,
		),
		svc,
	)

	// creation router
	var router = new(fasthttprouter.Router)

	errorProcessor := httperrors.NewErrorProcessor(httperrors.StatusMap)

	// init router
	httpserver.New(
		router,
		svc,
		// 7.таймаут на обработку одного входящего запроса - 10000 миллисекунд
		cfg.InTimeout,
		errorProcessor,
	)

	p := fasthttpprometheus.NewPrometheus(cfg.MetricsSubsystem)
	serverHandler := p.WrapHandler(router)

	fasthttpServer := &fasthttp.Server{
		// 6.сервер не обслуживает больше чем 100 одновременных входящих http-запросов
		Concurrency:        cfg.MaxSimultaneousConns,
		Handler:            serverHandler,
		MaxRequestBodySize: cfg.MaxRequestBodySize,
	}

	// listening port
	go func() {
		_ = level.Info(logger).Log("msg", "starting http server", "port", cfg.Port)
		if err := fasthttpServer.ListenAndServe(":" + cfg.Port); err != nil {
			_ = level.Error(logger).Log("msg", "server run failure", "err", err)
			os.Exit(1)
		}
	}()
	// 9.сервис должен поддерживать 'graceful shutdown': при получении
	//сигнала от OS перестать принимать входящие запросы, завершить текущие запросы и остановиться
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	defer func(sig os.Signal) {
		_ = level.Info(logger).Log("msg", "received signal, exiting", "signal", sig)

		if err := fasthttpServer.Shutdown(); err != nil {
			_ = level.Error(logger).Log("msg", "server shutdown failure", "err", err)
		}

		_ = level.Info(logger).Log("msg", "server stopped")
	}(<-c)
}
