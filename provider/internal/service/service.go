package service

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	videoCollectionV1 "github.com/wesleywu/ri-service-provider/provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/provider/internal/data"
	videoCollection "github.com/wesleywu/ri-service-provider/provider/internal/service/video_collection"
)

var ProviderSet = wire.NewSet(NewServices)

// Services .
type Services struct {
	videoCollection *videoCollection.ServiceImpl
}

func (s *Services) RegisterToGRPCServer(srv *grpc.Server) error {
	if s.videoCollection != nil {
		videoCollectionV1.RegisterVideoCollectionServer(srv, s.videoCollection)
		err := videoCollectionV1.RegisterVideoCollectionGuruServer(s.videoCollection)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewServices .
func NewServices(ctx context.Context, metadata *appinfo.AppMetadata, logger log.Logger, data *data.Data) (*Services, error) {
	helper := log.NewHelper(logger).WithContext(ctx)
	return &Services{
		videoCollection: videoCollection.NewVideoCollectionService(metadata, helper, data),
	}, nil
}
