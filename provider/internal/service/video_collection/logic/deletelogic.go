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

type DeleteLogic struct {
	metadata   *appinfo.AppMetadata
	helper     *log.Helper
	collection *mongo.Collection
}

func NewDeleteLogic(metadata *appinfo.AppMetadata, helper *log.Helper, collection *mongo.Collection) *DeleteLogic {
	return &DeleteLogic{
		metadata:   metadata,
		helper:     helper,
		collection: collection,
	}
}

func (s *DeleteLogic) Delete(ctx context.Context, req *p.VideoCollectionDeleteReq) (*p.VideoCollectionDeleteRes, error) {
	var (
		filterRequest gworm.FilterRequest
		err           error
		result        *gworm.Result
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
	result, err = m.Delete(ctx)
	if err != nil {
		return nil, err
	}
	deleteCount := result.RowsAffected
	message := "删除记录成功"
	if int(deleteCount) == 0 {
		message = "找不到要删除的记录"
	}
	return &p.VideoCollectionDeleteRes{
		Message:      gwwrapper.WrapString(message),
		RowsAffected: gwwrapper.WrapInt64(deleteCount),
	}, nil
}
