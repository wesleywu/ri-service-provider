package main

import (
	"context"
	"flag"
	"os"

	"github.com/castbox/go-guru/pkg/guru/service/conf"
	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/wesleywu/ri-service-provider/api/codec"
	"github.com/wesleywu/ri-service-provider/api/internal/service"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "helloworld"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
	// flagconf is the config flag.
	flagconf string
	// providerSet app related providers
	providerSet = wire.NewSet(newAppMetadata, newApp)
)

func init() {
	flag.StringVar(&flagconf, "conf", "api/configs", "config path, eg: -conf config.yaml")
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

	app, cleanup, err := wireApp(ctx, cfg.Server, cfg.Client, cfg.Data, cfg.Log, cfg.Otlp)
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

func newApp(ctx context.Context, m *appinfo.AppMetadata, logger log.Logger, server *http.Server, services *service.Services, tp *trace.TracerProvider) (*kratos.App, error) {
	err := services.RegisterToHTTPServer(server)
	if err != nil {
		return nil, err
	}
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
	if tp != nil {
		otel.SetTracerProvider(tp)
	}
	encoding.RegisterCodec(codec.JsonCodec{})
	return app, nil
}