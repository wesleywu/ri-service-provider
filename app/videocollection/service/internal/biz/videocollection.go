package biz

import (
	"context"

	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

type VideoCollectionRepo interface {
	Count(ctx context.Context, req *p.VideoCollectionCountReq) (*p.VideoCollectionCountRes, error)
	One(ctx context.Context, req *p.VideoCollectionOneReq) (*p.VideoCollectionOneRes, error)
	List(ctx context.Context, req *p.VideoCollectionListReq) (*p.VideoCollectionListRes, error)
	Get(ctx context.Context, req *p.VideoCollectionGetReq) (*p.VideoCollectionGetRes, error)
	Create(ctx context.Context, req *p.VideoCollectionCreateReq) (*p.VideoCollectionCreateRes, error)
	Update(ctx context.Context, req *p.VideoCollectionUpdateReq) (*p.VideoCollectionUpdateRes, error)
	Upsert(ctx context.Context, req *p.VideoCollectionUpsertReq) (*p.VideoCollectionUpsertRes, error)
	Delete(ctx context.Context, req *p.VideoCollectionDeleteReq) (*p.VideoCollectionDeleteRes, error)
	DeleteMulti(ctx context.Context, req *p.VideoCollectionDeleteMultiReq) (*p.VideoCollectionDeleteMultiRes, error)
}
