package main

import (
	"context"
	"flag"
	"os"

	"github.com/castbox/go-guru/pkg/goguru/conf"
	"github.com/castbox/go-guru/pkg/infra/appinfo"
	"github.com/castbox/go-guru/pkg/infra/logger"
	"github.com/castbox/go-guru/pkg/infra/mongodb"
	"github.com/castbox/go-guru/pkg/infra/otlp"
	"github.com/castbox/go-guru/pkg/infra/redis"
	"github.com/castbox/go-guru/pkg/middleware/servicecache"
	httpserver "github.com/castbox/go-guru/pkg/server/http"
	"github.com/castbox/go-guru/pkg/util/codec"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/service"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "ri-service-provider"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
	// flagconf is the config flag.
	flagconf string
	// providerSet app related providers
	providerSet = wire.NewSet(
		logger.NewConfigsByGuru,
		logger.NewLogger,
		logger.NewLoggerHelper,
		mongodb.NewConfigsByGuru,
		mongodb.NewMongoClient,
		redis.NewConfigsByGuru,
		redis.NewRedisCache,
		redis.NewRedisLock,
		servicecache.NewCacheProvider,
		otlp.NewConfigsByGuru,
		otlp.NewTracerProvider,
		useMiddlewares,
		httpserver.NewConfigsByGuru,
		httpserver.DefaultResponseEncoderFunc,
		httpserver.NewHTTPServer,
		newAppMetadata,
		newApp)
)

func init() {
	flag.StringVar(&flagconf, "conf", "configs/config.dev.yaml", "config path, eg: -conf config.dev.yaml")
}

func main() {
	ctx := context.Background()
	flag.Parse()
	c := config.New(
		config.WithSource(
			env.NewSource("GURU_"),
			file.NewSource(flagconf),
		),
	)
	defer func(c config.Config) {
		_ = c.Close()
	}(c)

	if err := c.Load(); err != nil {
		panic(err)
	}
	var cfg conf.Application
	if err := c.Scan(&cfg); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(ctx, cfg.Server.Http, cfg.Server.ServiceCache, cfg.Data.Database, cfg.Data.Redis, cfg.Log, cfg.Otlp)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()
	if err = app.Run(); err != nil {
		log.Fatal(err)
	}
}

func newAppMetadata() *appinfo.AppMetadata {
	hostname, _ := os.Hostname()
	return &appinfo.AppMetadata{
		AppName:    Name,
		AppVersion: Version,
		HostName:   hostname,
	}
}

func newApp(ctx context.Context, m *appinfo.AppMetadata, logger log.Logger, server *http.Server, registerInfo *service.RegisterInfo) (*kratos.App, error) {
	server.Handle("/metrics", promhttp.Handler())
	app := kratos.New(
		kratos.Context(ctx),
		kratos.Name(m.HostName),
		kratos.Version(m.AppVersion),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			server,
		),
	)
	encoding.RegisterCodec(codec.JsonCodec{})
	return app, nil
}

func useMiddlewares(cacheProvider servicecache.CacheProvider, logger log.Logger) []middleware.Middleware {
	return []middleware.Middleware{
		servicecache.NewCacheMiddleware(cacheProvider, logger),
	}
}
