package logic

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/gworm"
	"github.com/wesleywu/ri-service-provider/provider/internal/service/video_collection/mapping"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListLogic struct {
	metadata   *appinfo.AppMetadata
	helper     *log.Helper
	collection *mongo.Collection
}

func NewListLogic(metadata *appinfo.AppMetadata, helper *log.Helper, collection *mongo.Collection) *ListLogic {
	return &ListLogic{
		metadata:   metadata,
		helper:     helper,
		collection: collection,
	}
}

func (s *ListLogic) List(ctx context.Context, req *p.VideoCollectionListReq) (*p.VideoCollectionListRes, error) {
	var (
		filterRequest   gworm.FilterRequest
		pageRequest     *gworm.PageRequest
		list            []*p.VideoCollectionItem
		pageInfo        *gworm.PageInfo
		pageRecordCount int
		total           int64
		err             error
	)
	filterRequest, err = gworm.ExtractFilters(ctx, req, mapping.VideoCollectionColumnMap)
	if err != nil {
		return nil, err
	}
	pageRequest = gworm.NewPageRequest(req.Page, req.PageSize, req.OrderBy)

	opts := options.Find().SetLimit(pageRequest.Size).SetSkip(pageRequest.Offset)
	if pageRequest.HasSort() {
		opts.SetSort(pageRequest.MongoSortOption())
	}

	list = []*p.VideoCollectionItem{}
	cursor, err := s.collection.Find(ctx, filterRequest.Filters, opts)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	pageRecordCount = len(list)
	// 如果没有初始偏移量，并且返回记录条数 < pageSize，说明符合条件的记录已经全部返回
	if pageRequest.Offset == 0 && pageRecordCount < int(pageRequest.Size) {
		total = int64(len(list))
	} else { // 否则执行 CountDocument 重新计算所有记录条数
		total, err = s.collection.CountDocuments(ctx, filterRequest.Filters, nil)
		if err != nil {
			return nil, err
		}
	}
	pageInfo = &gworm.PageInfo{}
	pageInfo.From(pageRequest, int64(len(list)), total)

	return &p.VideoCollectionListRes{
		Total:   total,
		Current: pageInfo.Number,
		Items:   list,
	}, nil
}
