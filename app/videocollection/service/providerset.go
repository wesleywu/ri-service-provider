package service

import (
	"github.com/google/wire"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/data"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/service"
)

var ProviderSet = wire.NewSet(
	data.NewVideoCollectionRepo,
	service.NewVideoCollectionService,
	service.RegisterToHTTPServer,
)
