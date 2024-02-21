package logic

import (
	"context"

	apiErrors "github.com/castbox/go-guru/pkg/goguru/error"
	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/gwerror"
	"github.com/wesleywu/ri-service-provider/gworm"
	"github.com/wesleywu/ri-service-provider/gwwrapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteLogic struct {
	metadata   *appinfo.AppMetadata
	helper     *log.Helper
	collection *mongo.Collection
}

func NewDeleteLogic(metadata *appinfo.AppMetadata, helper *log.Helper, collection *mongo.Collection) *DeleteLogic {
	return &DeleteLogic{
		metadata:   metadata,
		helper:     helper,
		collection: collection,
	}
}

func (s *DeleteLogic) Delete(ctx context.Context, req *p.VideoCollectionDeleteReq) (*p.VideoCollectionDeleteRes, error) {
	var (
		filterRequest gworm.FilterRequest
		filters       *bson.D
		result        *mongo.DeleteResult
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
	result, err = s.collection.DeleteOne(ctx, filters, nil)
	if err != nil {
		s.helper.WithContext(ctx).Error(err)
		err = gwerror.WrapServiceErrorf(err, req, "删除记录失败")
		return nil, err
	}
	deleteCount := result.DeletedCount
	message := "删除记录成功"
	if int(deleteCount) == 0 {
		message = "找不到要删除的记录"
	}
	return &p.VideoCollectionDeleteRes{
		Message:      gwwrapper.WrapString(message),
		RowsAffected: gwwrapper.WrapInt64(deleteCount),
	}, nil
}
