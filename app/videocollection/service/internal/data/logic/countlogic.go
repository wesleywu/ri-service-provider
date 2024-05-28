package logic

import (
	"context"
	"fmt"

	"github.com/castbox/go-guru/pkg/util/mongodb/filters"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/data/mapping"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CountLogic struct {
	collection *mongo.Collection
	helper     *log.Helper
}

func NewCountLogic(collection *mongo.Collection, helper *log.Helper) *CountLogic {
	return &CountLogic{
		collection: collection,
		helper:     helper,
	}
}

func (s *CountLogic) Count(ctx context.Context, req *p.VideoCollectionCountReq) (*p.VideoCollectionCountRes, error) {
	var (
		filter *bson.D
		count  int64
		err    error
	)
	filter, err = filters.ExtractBsonFilter(req, mapping.VideoCollectionColumnMap, s.helper)
	if err != nil {
		return nil, err
	}
	count, err = s.collection.CountDocuments(ctx, filter, nil)
	if err != nil {
		s.helper.WithContext(ctx).Error(err)
		err = errors.Wrap(err, fmt.Sprintf("获取数据记录总数失败: %v", req))
		return nil, err
	}
	return &p.VideoCollectionCountRes{
		TotalElements: count,
	}, nil
}
