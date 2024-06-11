package service

import (
	"github.com/google/wire"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/service"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

var ProviderSet = wire.NewSet(
	p.NewVideoCollectionRepo,
	service.NewVideoCollectionService,
	service.RegisterToHTTPServer,
	service.RegisterToGRPCServer,
)
