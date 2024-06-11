//go:build wireinject
// +build wireinject

package rpc_client

import (
	"context"

	"github.com/google/wire"
)

func wireClient(context.Context) (*Clients, func(), error) {
	panic(wire.Build(ProviderSet))
}
