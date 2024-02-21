package logic

import (
	"context"
	"errors"

	apiErrors "github.com/castbox/go-guru/pkg/goguru/error"
	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
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
		item          *p.VideoCollectionGetRes
	)
	if req.Id == "" {
		return nil, apiErrors.ErrorIdValueMissing("主键ID字段的值为空")
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
		if errors.Is(singleResult.Err(), mongo.ErrNoDocuments) {
			return nil, apiErrors.ErrorRecordNotFound("找不到ID为 %s 的记录", req.Id)
		}
		return nil, err
	}
	item = (*p.VideoCollectionGetRes)(nil)
	err = singleResult.Decode(&item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
