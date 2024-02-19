package logic

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/api/errors"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/gworm"
	"github.com/wesleywu/ri-service-provider/gwwrapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UpsertLogic struct {
	metadata   *appinfo.AppMetadata
	helper     *log.Helper
	collection *mongo.Collection
}

func NewUpsertLogic(metadata *appinfo.AppMetadata, helper *log.Helper, collection *mongo.Collection) *UpsertLogic {
	return &UpsertLogic{
		metadata:   metadata,
		helper:     helper,
		collection: collection,
	}
}

func (s *UpsertLogic) Upsert(ctx context.Context, req *p.VideoCollectionUpsertReq) (*p.VideoCollectionUpsertRes, error) {
	var (
		filterRequest gworm.FilterRequest
		filters       *bson.D
		updateResult  *mongo.UpdateResult
		err           error
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
	update := bson.D{
		{
			"$set", req,
		},
	}
	opts := options.Update().SetUpsert(true)
	updateResult, err = s.collection.UpdateOne(ctx, filters, update, opts)
	if err != nil {
		return nil, err
	}
	return &p.VideoCollectionUpsertRes{
		Message:      gwwrapper.WrapString("更新记录成功"),
		RowsAffected: gwwrapper.WrapInt64(updateResult.UpsertedCount),
	}, err
}
