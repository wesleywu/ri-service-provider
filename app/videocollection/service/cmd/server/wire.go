//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/castbox/go-guru/pkg/goguru/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service"
)

func wireApp(context.Context, *conf.Server, *conf.Server_ServiceCache, *conf.Database, *conf.Redis, *conf.Log, *conf.Otlp) (*kratos.App, func(), error) {
	panic(wire.Build(service.ProviderSet, providerSet))
}