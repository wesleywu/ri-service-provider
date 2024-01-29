package client

import (
	"context"

	"github.com/castbox/go-guru/pkg/client"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

var ProviderSet = wire.NewSet(client.NewGrpcConnection, NewClients)

type Clients struct {
	VideoCollection p.VideoCollectionClient
}

// NewClients .
func NewClients(ctx context.Context, conn *grpc.ClientConn) (*Clients, error) {
	return &Clients{
		VideoCollection: p.NewVideoCollectionClient(conn),
	}, nil
}
