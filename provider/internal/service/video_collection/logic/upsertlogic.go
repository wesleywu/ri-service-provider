package logic

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/gworm"
	"github.com/wesleywu/ri-service-provider/gworm/mongodb"
	"github.com/wesleywu/ri-service-provider/gwwrapper"
	"github.com/wesleywu/ri-service-provider/provider/internal/service/video_collection/mapping"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpsertLogic struct {
	metadata   *appinfo.AppMetadata
	helper     *log.Helper
	collection *mongo.Collection
}

func NewUpsertLogic(metadata *appinfo.AppMetadata, helper *log.Helper, collection *mongo.Collection) *UpsertLogic {
	return &UpsertLogic{
		metadata:   metadata,
		helper:     helper,
		collection: collection,
	}
}

func (s *UpsertLogic) Upsert(ctx context.Context, req *p.VideoCollectionUpsertReq) (*p.VideoCollectionUpsertRes, error) {
	var (
		filterRequest gworm.FilterRequest
		result        *gworm.Result
		err           error
	)
	filterRequest, err = gworm.ExtractFilters(ctx, req, mapping.VideoCollectionColumnMap, gworm.MONGO)
	m := &gworm.Model{
		Type:       gworm.MONGO,
		MongoModel: mongodb.NewModel(s.collection),
	}
	m, err = gworm.ApplyFilter(ctx, filterRequest, m)
	if err != nil {
		return nil, err
	}
	result, err = m.Fields(p.VideoCollectionItem{}).
		FieldsEx(mapping.VideoCollectionColumns.Id, mapping.VideoCollectionColumns.CreatedAt).
		Upsert(ctx, req)
	if err != nil {
		return nil, err
	}
	return &p.VideoCollectionUpsertRes{
		Message:      gwwrapper.WrapString("创建记录成功"),
		RowsAffected: gwwrapper.WrapInt64(result.RowsAffected),
	}, err
}
