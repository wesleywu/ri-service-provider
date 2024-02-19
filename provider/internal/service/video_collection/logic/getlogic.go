package logic

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/api/errors"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/gworm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetLogic struct {
	metadata   *appinfo.AppMetadata
	helper     *log.Helper
	collection *mongo.Collection
}

func NewGetLogic(metadata *appinfo.AppMetadata, helper *log.Helper, collection *mongo.Collection) *GetLogic {
	return &GetLogic{
		metadata:   metadata,
		helper:     helper,
		collection: collection,
	}
}

func (s *GetLogic) Get(ctx context.Context, req *p.VideoCollectionGetReq) (*p.VideoCollectionGetRes, error) {
	var (
		filterRequest gworm.FilterRequest
		filters       *bson.D
		err           error
		item          *p.VideoCollectionItem
	)
	if req.Id == "" {
		return nil, errors.ErrorIdValueMissing("主键ID字段的值为空")
	}
	filterRequest = gworm.FilterRequest{
		PropertyFilters: []*gworm.PropertyFilter{
			{
				Property: "_id",
				Value:    req.Id,
			},
		},
	}
	filters, err = filterRequest.GetFilters()
	if err != nil {
		// todo parameter error
		return nil, err
	}
	singleResult := s.collection.FindOne(ctx, filters, nil)
	if singleResult.Err() != nil {
		return nil, err
	}
	item = (*p.VideoCollectionItem)(nil)
	err = singleResult.Decode(&item)
	if err != nil {
		return nil, err
	}
	return &p.VideoCollectionGetRes{
		Found: true,
		Item:  item,
	}, nil
}
