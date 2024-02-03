package logic

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/gworm"
	"github.com/wesleywu/ri-service-provider/gwwrapper"
	"github.com/wesleywu/ri-service-provider/provider/internal/service/video_collection/mapping"
	"go.mongodb.org/mongo-driver/mongo"
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
		filterRequest gworm.FilterRequest
		pageRequest   gworm.PageRequest
		list          []*p.VideoCollectionItem
		pageInfo      *gworm.PageInfo
		total         int64
		err           error
	)
	filterRequest, err = gworm.ExtractFilters(ctx, req, mapping.VideoCollectionColumnMap)
	// todo: page size not set to default
	if err != nil {
		return nil, err
	}
	pageRequest.AddSortByString(req.OrderBy)
	list = []*p.VideoCollectionItem{}
	// todo: 实现翻页功能
	cursor, err := s.collection.Find(ctx, filterRequest.Filters, nil)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	if len(list) == int(pageRequest.Size) {
		total, err = s.collection.CountDocuments(ctx, filterRequest.Filters, nil)
		if err != nil {
			return nil, err
		}
	} else {
		total = int64(len(list))
	}
	pageInfo = &gworm.PageInfo{}
	pageInfo.From(pageRequest, uint32(len(list)), uint64(total))

	return &p.VideoCollectionListRes{
		Total:   gwwrapper.WrapInt64(total),
		Current: gwwrapper.WrapUInt32(pageInfo.Number),
		Items:   list,
	}, nil
}
