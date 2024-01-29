package server

import (
	"github.com/castbox/go-guru/pkg/server"
	"github.com/castbox/go-guru/pkg/util/logger"
	"github.com/castbox/go-guru/pkg/util/otlp"
	"github.com/castbox/go-guru/pkg/util/redis"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(logger.NewLogger, otlp.NewTracer, server.NewHTTPServer, redis.NewRedisCache)
