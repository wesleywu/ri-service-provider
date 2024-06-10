package rpc_client

import (
	"context"
	"errors"
	"os"

	grpcclient "github.com/castbox/go-guru/pkg/client/grpc"
	"github.com/castbox/go-guru/pkg/infra/appinfo"
	"github.com/castbox/go-guru/pkg/infra/logger"
	"github.com/castbox/go-guru/pkg/infra/otlp"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/wesleywu/ri-service-provider/api/videocollection/service/v1"
	"google.golang.org/grpc"
)

const clientProfileKey = "video-collection"

var ProviderSet = wire.NewSet(
	newAppMetadata,
	logger.NewConfigsByGuru,
	logger.NewLogger,
	logger.NewLoggerHelper,
	grpcclient.NewConfigsByGuru,
	grpcclient.NewGrpcConnections,
	otlp.NewConfigsByGuru,
	otlp.NewTracerProvider,
	NewClients)

type Clients struct {
	logger          log.Logger
	VideoCollection v1.VideoCollectionClient
}

func newAppMetadata() *appinfo.AppMetadata {
	hostname, _ := os.Hostname()
	return &appinfo.AppMetadata{
		AppName:    "rpc_client_test",
		AppVersion: "v0.0.1",
		HostName:   hostname,
	}
}

// NewClients .
func NewClients(ctx context.Context, conns map[string]*grpc.ClientConn, logger log.Logger) (*Clients, error) {
	if conns == nil {
		return nil, errors.New("没有配置grpc client")
	}
	if conn, ok := conns[clientProfileKey]; ok {
		return &Clients{
			logger:          logger,
			VideoCollection: v1.NewVideoCollectionClient(conn),
		}, nil
	}
	return nil, errors.New("没有配置grpc client")
}
