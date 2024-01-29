package logic

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/gwwrapper"
	p "github.com/wesleywu/ri-service-provider/provider/api/video_collection/v1"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateLogic struct {
	metadata   *appinfo.AppMetadata
	helper     *log.Helper
	collection *mongo.Collection
}

func NewCreateLogic(metadata *appinfo.AppMetadata, helper *log.Helper, collection *mongo.Collection) *CreateLogic {
	return &CreateLogic{
		metadata:   metadata,
		helper:     helper,
		collection: collection,
	}
}

func (s *CreateLogic) Create(ctx context.Context, req *p.VideoCollectionCreateReq) (*p.VideoCollectionCreateRes, error) {
	_, err := s.collection.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	return &p.VideoCollectionCreateRes{
		Message:      gwwrapper.WrapString("创建记录成功"),
		RowsAffected: gwwrapper.WrapInt64(1),
	}, err
}
