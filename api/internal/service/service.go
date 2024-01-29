package service

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/wesleywu/ri-service-provider/api/internal/client"
	videoCollection "github.com/wesleywu/ri-service-provider/api/internal/service/video_collection"
	videoCollectionV1 "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
)

var ProviderSet = wire.NewSet(NewServices)

// Services .
type Services struct {
	videoCollection *videoCollection.ServiceImpl
}

func (s *Services) RegisterToHTTPServer(srv *http.Server) error {
	if s.videoCollection != nil {
		videoCollectionV1.RegisterVideoCollectionHTTPServer(srv, s.videoCollection)
		err := videoCollectionV1.RegisterVideoCollectionGuruServer(s.videoCollection)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewServices .
func NewServices(ctx context.Context, metadata *appinfo.AppMetadata, logger log.Logger, client *client.Clients) (*Services, error) {
	helper := log.NewHelper(logger).WithContext(ctx)
	return &Services{
		videoCollection: videoCollection.NewVideoCollectionService(metadata, helper, client),
	}, nil
}
