//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/castbox/go-guru/pkg/goguru/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/wesleywu/ri-service-provider/provider/internal/data"
	"github.com/wesleywu/ri-service-provider/provider/internal/server"
	"github.com/wesleywu/ri-service-provider/provider/internal/service"
)

func wireApp(context.Context, *conf.Server, *conf.Data, *conf.Log, *conf.Otlp) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, providerSet))
}
