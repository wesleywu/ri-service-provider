package logic

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/gwerror"
	"github.com/wesleywu/ri-service-provider/gworm"
	"github.com/wesleywu/ri-service-provider/gwwrapper"
	"github.com/wesleywu/ri-service-provider/provider/internal/service/video_collection/mapping"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteMultiLogic struct {
	metadata   *appinfo.AppMetadata
	helper     *log.Helper
	collection *mongo.Collection
}

func NewDeleteMultiLogic(metadata *appinfo.AppMetadata, helper *log.Helper, collection *mongo.Collection) *DeleteMultiLogic {
	return &DeleteMultiLogic{
		metadata:   metadata,
		helper:     helper,
		collection: collection,
	}
}

func (s *DeleteMultiLogic) DeleteMulti(ctx context.Context, req *p.VideoCollectionDeleteMultiReq) (*p.VideoCollectionDeleteMultiRes, error) {
	var (
		filterRequest gworm.FilterRequest
		count         int64
		deleteResult  *mongo.DeleteResult
		err           error
	)
	filterRequest, err = gworm.ExtractFilters(ctx, req, mapping.VideoCollectionColumnMap)
	if err != nil {
		// todo parameter error
		return nil, err
	}
	deleteResult, err = s.collection.DeleteMany(ctx, filterRequest.Filters, nil)
	if err != nil {
		s.helper.WithContext(ctx).Error(err)
		err = gwerror.WrapServiceErrorf(err, req, "删除记录失败")
		return nil, err
	}
	count = deleteResult.DeletedCount
	message := "删除记录成功"
	if int(count) == 0 {
		message = "找不到要删除的记录"
	}
	return &p.VideoCollectionDeleteMultiRes{
		Message:      gwwrapper.WrapString(message),
		RowsAffected: gwwrapper.WrapInt64(count),
	}, nil
}
