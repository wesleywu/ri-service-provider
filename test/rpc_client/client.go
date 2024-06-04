package rpc_client

import (
	"context"
	"os"

	grpcserver "github.com/castbox/go-guru/pkg/client/grpc"
	"github.com/castbox/go-guru/pkg/infra/appinfo"
	"github.com/castbox/go-guru/pkg/infra/logger"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/wesleywu/ri-service-provider/api/videocollection/service/v1"
	"google.golang.org/grpc"
)

var ProviderSet = wire.NewSet(newAppMetadata, logger.NewLogger, grpcserver.NewGrpcConnections, NewClients)

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
func NewClients(ctx context.Context, conn *grpc.ClientConn, logger log.Logger) (*Clients, error) {
	return &Clients{
		logger:          logger,
		VideoCollection: v1.NewVideoCollectionClient(conn),
	}, nil
}
