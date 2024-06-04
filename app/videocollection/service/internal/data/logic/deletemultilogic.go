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

type DeleteMultiLogic struct {
	collection *mongo.Collection
	helper     *log.Helper
}

func NewDeleteMultiLogic(collection *mongo.Collection, helper *log.Helper) *DeleteMultiLogic {
	return &DeleteMultiLogic{
		collection: collection,
		helper:     helper,
	}
}

func (s *DeleteMultiLogic) DeleteMulti(ctx context.Context, req *p.VideoCollectionDeleteMultiReq) (*p.VideoCollectionDeleteMultiRes, error) {
	var (
		filter       *bson.D
		count        int64
		deleteResult *mongo.DeleteResult
		err          error
	)
	filter, err = filters.ExtractBsonFilter(req, mapping.VideoCollectionColumnMap, s.helper)
	if err != nil {
		return nil, err
	}
	deleteResult, err = s.collection.DeleteMany(ctx, filter, nil)
	if err != nil {
		s.helper.WithContext(ctx).Error(err)
		err = errors.Wrap(err, fmt.Sprintf("删除记录失败: %v", req))
		return nil, err
	}
	count = deleteResult.DeletedCount
	message := "删除记录成功"
	if int(count) == 0 {
		message = "找不到要删除的记录"
	}
	return &p.VideoCollectionDeleteMultiRes{
		Message:      message,
		DeletedCount: count,
	}, nil
}
