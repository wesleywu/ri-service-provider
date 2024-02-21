package logic

import (
	"context"

	apiErrors "github.com/castbox/go-guru/pkg/goguru/error"
	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/gworm"
	"github.com/wesleywu/ri-service-provider/gwwrapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateLogic struct {
	metadata   *appinfo.AppMetadata
	helper     *log.Helper
	collection *mongo.Collection
}

func NewUpdateLogic(metadata *appinfo.AppMetadata, helper *log.Helper, collection *mongo.Collection) *UpdateLogic {
	return &UpdateLogic{
		metadata:   metadata,
		helper:     helper,
		collection: collection,
	}
}

func (s *UpdateLogic) Update(ctx context.Context, req *p.VideoCollectionUpdateReq) (*p.VideoCollectionUpdateRes, error) {
	var (
		filterRequest gworm.FilterRequest
		filters       *bson.D
		updateResult  *mongo.UpdateResult
		err           error
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
	update := bson.D{
		{
			"$set", req,
		},
	}
	updateResult, err = s.collection.UpdateOne(ctx, filters, update, nil)
	if err != nil {
		return nil, err
	}
	return &p.VideoCollectionUpdateRes{
		Message:      gwwrapper.WrapString("更新记录成功"),
		RowsAffected: gwwrapper.WrapInt64(updateResult.ModifiedCount),
	}, err
}
