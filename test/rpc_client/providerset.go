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
	v1 "github.com/wesleywu/ri-service-provider/api/episode/service/v1"
	"google.golang.org/grpc"
)

const (
	grpcClientProfileKey = "video-collection"
	grpcServerEndpoint   = "127.0.0.1:22000"
	otlpHttpEndpoint     = "34.120.15.175"
	otlpBasicAuthToken   = "bWVuZ3llLnd1QGNhc3Rib3guZm06VUxGV3BnQ3ZJNUdDWDhCTA=="
	otlpInsecure         = true
)

var ProviderSet = wire.NewSet(
	newAppMetadata,
	logger.NewConfigs,
	logger.NewLogger,
	logger.NewLoggerHelper,
	newGrpcConfigs,
	grpcclient.NewGrpcConnections,
	newOtlpConfigs,
	otlp.NewTracerProvider,
	NewClients)

type Clients struct {
	logger  log.Logger
	Episode v1.EpisodeClient
}

func newAppMetadata() *appinfo.AppMetadata {
	hostname, _ := os.Hostname()
	return &appinfo.AppMetadata{
		AppName:    "rpc_client_test",
		AppVersion: "v0.0.1",
		HostName:   hostname,
	}
}

func newGrpcConfigs() *grpcclient.Configs {
	profile := grpcclient.NewProfileDefault(grpcClientProfileKey, grpcServerEndpoint)
	profile.ApplyLoggingMiddleware(true)
	profile.ApplyTraceMiddleware(true)
	profile.ApplyRecoverMiddleware(true)
	profile.ApplyPrometheusMiddleware(true)
	return grpcclient.NewConfigs(profile)
}

func newOtlpConfigs(m *appinfo.AppMetadata) *otlp.Configs {
	return otlp.NewHttpTpConfigs(m, otlpHttpEndpoint, otlpBasicAuthToken, otlpInsecure)
}

// NewClients .
func NewClients(ctx context.Context, conns map[string]*grpc.ClientConn, logger log.Logger) (*Clients, error) {
	if conns == nil {
		return nil, errors.New("没有配置grpc client")
	}
	if conn, ok := conns[grpcClientProfileKey]; ok {
		return &Clients{
			logger:  logger,
			Episode: v1.NewEpisodeClient(conn),
		}, nil
	}
	return nil, errors.New("没有配置grpc client")
}
