package logic

import (
	"context"
	"fmt"

	"github.com/castbox/go-guru/pkg/infra/mongodb/filters"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/data/mapping"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OneLogic struct {
	collection *mongo.Collection
	helper     *log.Helper
}

func NewOneLogic(collection *mongo.Collection, helper *log.Helper) *OneLogic {
	return &OneLogic{
		collection: collection,
		helper:     helper,
	}
}

func (s *OneLogic) One(ctx context.Context, req *p.VideoCollectionOneReq) (*p.VideoCollectionOneRes, error) {
	var (
		filter       *bson.D
		singleResult *mongo.SingleResult
		item         *p.VideoCollectionItem
		err          error
	)
	filter, err = filters.ExtractBsonFilter(req, mapping.VideoCollectionColumnMap, s.helper)
	if err != nil {
		return nil, err
	}
	singleResult = s.collection.FindOne(ctx, filter, nil)
	if err = singleResult.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &p.VideoCollectionOneRes{
				Found: false,
				Item:  nil,
			}, nil
		}
		s.helper.WithContext(ctx).Error(err)
		err = errors.Wrap(err, fmt.Sprintf("单条记录查询失败: %v", req))
		return nil, err
	}
	item = (*p.VideoCollectionItem)(nil)
	err = singleResult.Decode(&item)
	if err != nil {
		s.helper.WithContext(ctx).Error(err)
		err = errors.Wrap(err, fmt.Sprintf("获取单条数据记录失败: %v", req))
		return nil, err
	}
	return &p.VideoCollectionOneRes{
		Found: true,
		Item:  item,
	}, nil
}
