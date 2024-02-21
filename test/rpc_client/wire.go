//go:build wireinject
// +build wireinject

package rpc_client

import (
	"context"

	"github.com/castbox/go-guru/pkg/goguru/conf"
	"github.com/google/wire"
)

func wireClient(context.Context, *conf.Client, *conf.Log, *conf.Otlp) (*Clients, func(), error) {
	panic(wire.Build(ProviderSet))
}
