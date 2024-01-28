package logic

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/gowing/util/gworm"
	"github.com/wesleywu/gowing/util/gworm/mongodb"
	"github.com/wesleywu/gowing/util/gwwrapper"
	p "github.com/wesleywu/ri-service-provider/provider/api/video_collection/v1"
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
	total, err = m.MongoModel.Count(ctx)
	if err != nil {
		return nil, err
	}
	list = []*p.VideoCollectionItem{}
	err = m.Fields(p.VideoCollectionItem{}).
		Page(int(pageRequest.Number), int(pageRequest.Size)).
		Order(pageRequest.OrderString()).
		Scan(ctx, &list)
	if err != nil {
		return nil, err
	}
	pageInfo = &gworm.PageInfo{}
	pageInfo.From(pageRequest, uint32(len(list)), uint64(total))

	return &p.VideoCollectionListRes{
		Total:   gwwrapper.WrapInt64(total),
		Current: gwwrapper.WrapUInt32(pageInfo.Number),
		Items:   list,
	}, nil
}
