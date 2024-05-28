package logic

import (
	"context"
	"fmt"

	"github.com/castbox/go-guru/pkg/goguru/query"
	"github.com/castbox/go-guru/pkg/util/mongodb/filters"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/data/mapping"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListLogic struct {
	collection *mongo.Collection
	helper     *log.Helper
}

func NewListLogic(collection *mongo.Collection, helper *log.Helper) *ListLogic {
	return &ListLogic{
		collection: collection,
		helper:     helper,
	}
}

func (s *ListLogic) List(ctx context.Context, req *p.VideoCollectionListReq) (*p.VideoCollectionListRes, error) {
	var (
		filter          *bson.D
		pageRequest     *query.PageRequest
		list            []*p.VideoCollectionItem
		pageInfo        *query.PageInfo
		pageRecordCount int
		total           int64
		err             error
	)
	filter, err = filters.ExtractBsonFilter(req, mapping.VideoCollectionColumnMap, s.helper)
	if err != nil {
		return nil, err
	}
	if req.PageRequest != nil {
		pageRequest = req.PageRequest.FillDefault()
	} else {
		pageRequest = (&query.PageRequest{}).FillDefault()
	}

	opts := options.Find().SetLimit(pageRequest.Size).SetSkip(pageRequest.StartOffset())
	if pageRequest.HasSort() {
		opts.SetSort(pageRequest.MongoSortOption())
	}

	list = []*p.VideoCollectionItem{}
	cursor, err := s.collection.Find(ctx, filter, opts)
	if err != nil {
		s.helper.WithContext(ctx).Error(err)
		err = errors.Wrap(err, fmt.Sprintf("列表查询失败: %v", req))
		return nil, err
	}
	err = cursor.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	pageRecordCount = len(list)
	// 如果没有初始偏移量，并且返回记录条数 < pageSize，说明符合条件的记录已经全部返回
	if pageRequest.StartOffset() == 0 && pageRecordCount < int(pageRequest.Size) {
		total = int64(len(list))
	} else { // 否则执行 CountDocument 重新计算所有记录条数
		total, err = s.collection.CountDocuments(ctx, filter, nil)
		if err != nil {
			return nil, err
		}
	}
	pageInfo = &query.PageInfo{}
	pageInfo.From(pageRequest, int64(len(list)), total)

	return &p.VideoCollectionListRes{
		PageInfo: pageInfo,
		Items:    list,
	}, nil
}
