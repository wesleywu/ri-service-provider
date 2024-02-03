package logic

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/gwerror"
	"github.com/wesleywu/ri-service-provider/gworm"
	"github.com/wesleywu/ri-service-provider/provider/internal/service/video_collection/mapping"
	"go.mongodb.org/mongo-driver/mongo"
)

type OneLogic struct {
	metadata   *appinfo.AppMetadata
	helper     *log.Helper
	collection *mongo.Collection
}

func NewOneLogic(metadata *appinfo.AppMetadata, helper *log.Helper, collection *mongo.Collection) *OneLogic {
	return &OneLogic{
		metadata:   metadata,
		helper:     helper,
		collection: collection,
	}
}

func (s *OneLogic) One(ctx context.Context, req *p.VideoCollectionOneReq) (*p.VideoCollectionOneRes, error) {
	var (
		filterRequest gworm.FilterRequest
		pageRequest   gworm.PageRequest
		singleResult  *mongo.SingleResult
		item          *p.VideoCollectionItem
		err           error
	)
	filterRequest, err = gworm.ExtractFilters(ctx, req, mapping.VideoCollectionColumnMap)
	pageRequest.AddSortByString(req.OrderBy)
	singleResult = s.collection.FindOne(ctx, filterRequest.Filters, nil)
	if singleResult.Err() != nil {
		return nil, singleResult.Err()
	}
	item = (*p.VideoCollectionItem)(nil)
	err = singleResult.Decode(item)
	if err != nil {
		s.helper.WithContext(ctx).Error(err)
		err = gwerror.WrapServiceErrorf(err, req, "获取单条数据记录失败")
		return nil, err
	}
	return &p.VideoCollectionOneRes{
		Found: true,
		Item:  item,
	}, nil
}
