package logic

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/gowing/errors/gwerror"
	"github.com/wesleywu/gowing/util/gworm"
	"github.com/wesleywu/gowing/util/gworm/mongodb"
	p "github.com/wesleywu/ri-service-provider/provider/api/video_collection/v1"
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
		item          *p.VideoCollectionItem
		err           error
	)
	filterRequest, err = gworm.ExtractFilters(ctx, req, mapping.VideoCollectionColumnMap, gworm.MONGO)
	pageRequest.AddSortByString(req.OrderBy)
	m := &gworm.Model{
		Type:       gworm.MONGO,
		MongoModel: mongodb.NewModel(ctx, s.collection.Name()),
	}
	m, err = gworm.ApplyFilter(ctx, filterRequest, m)
	if err != nil {
		return nil, err
	}
	item = (*p.VideoCollectionItem)(nil)
	err = m.Fields(p.VideoCollectionItem{}).
		Order(pageRequest.OrderString()).
		Limit(1).
		Scan(ctx, &item)
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
