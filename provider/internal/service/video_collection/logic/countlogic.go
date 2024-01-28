package logic

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/gowing/errors/gwerror"
	"github.com/wesleywu/gowing/util/gworm"
	"github.com/wesleywu/gowing/util/gworm/mongodb"
	"github.com/wesleywu/gowing/util/gwwrapper"
	p "github.com/wesleywu/ri-service-provider/provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/provider/internal/service/video_collection/mapping"
	"go.mongodb.org/mongo-driver/mongo"
)

type CountLogic struct {
	metadata   *appinfo.AppMetadata
	helper     *log.Helper
	collection *mongo.Collection
}

func NewCountLogic(metadata *appinfo.AppMetadata, helper *log.Helper, collection *mongo.Collection) *CountLogic {
	return &CountLogic{
		metadata:   metadata,
		helper:     helper,
		collection: collection,
	}
}

func (s *CountLogic) Count(ctx context.Context, req *p.VideoCollectionCountReq) (*p.VideoCollectionCountRes, error) {
	var (
		filterRequest gworm.FilterRequest
		count         int64
		err           error
	)
	filterRequest, err = gworm.ExtractFilters(ctx, req, mapping.VideoCollectionColumnMap, gworm.MONGO)
	m := &gworm.Model{
		Type:       gworm.MONGO,
		MongoModel: mongodb.NewModel(ctx, s.collection.Name()),
	}
	m, err = gworm.ApplyFilter(ctx, filterRequest, m)
	if err != nil {
		return nil, err
	}
	count, err = m.MongoModel.Count(ctx)
	if err != nil {
		s.helper.WithContext(ctx).Error(err)
		err = gwerror.WrapServiceErrorf(err, req, "获取数据记录总数失败")
		return nil, err
	}
	return &p.VideoCollectionCountRes{
		Total: gwwrapper.WrapInt64(count),
	}, nil
}
