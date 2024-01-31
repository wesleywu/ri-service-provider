package logic

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/api/errors"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/gworm"
	"github.com/wesleywu/ri-service-provider/gworm/mongodb"
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
	m := &gworm.Model{
		Type:       gworm.MONGO,
		MongoModel: mongodb.NewModel(s.collection),
	}
	m, err = gworm.ApplyFilter(ctx, filterRequest, m)
	if err != nil {
		return nil, err
	}
	singleResult := s.collection.FindOne(ctx, m.MongoModel.Filter)
	if singleResult.Err() != nil {
		return nil, err
	}
	item = (*p.VideoCollectionItem)(nil)
	err = singleResult.Decode(item)
	if err != nil {
		return nil, err
	}
	return &p.VideoCollectionGetRes{
		Found: true,
		Item:  item,
	}, nil
}
