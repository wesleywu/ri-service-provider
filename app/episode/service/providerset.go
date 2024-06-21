package service

import (
	"github.com/google/wire"
	"github.com/wesleywu/ri-service-provider/app/episode/service/internal/service"
	p "github.com/wesleywu/ri-service-provider/app/episode/service/proto"
)

var ProviderSet = wire.NewSet(
	p.NewEpisodeRepo,
	service.NewEpisodeService,
	service.RegisterToHTTPServer,
	service.RegisterToGRPCServer,
)
