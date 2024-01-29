//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/castbox/go-guru/pkg/guru/service/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/wesleywu/ri-service-provider/api/internal/client"
	"github.com/wesleywu/ri-service-provider/api/internal/server"
	"github.com/wesleywu/ri-service-provider/api/internal/service"
)

func wireApp(context.Context, *conf.Server, *conf.Client, *conf.Data, *conf.Log, *conf.Otlp) (*kratos.App, func(), error) {
	panic(wire.Build(client.ProviderSet, server.ProviderSet, service.ProviderSet, providerSet))
}
