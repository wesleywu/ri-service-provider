package logic

import (
	"context"
	"fmt"

	guruErrors "github.com/castbox/go-guru/pkg/goguru/error"
	"github.com/castbox/go-guru/pkg/util/mongodb/filters"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteLogic struct {
	collection *mongo.Collection
	helper     *log.Helper
}

func NewDeleteLogic(collection *mongo.Collection, helper *log.Helper) *DeleteLogic {
	return &DeleteLogic{
		collection: collection,
		helper:     helper,
	}
}

func (s *DeleteLogic) Delete(ctx context.Context, req *p.VideoCollectionDeleteReq) (*p.VideoCollectionDeleteRes, error) {
	var (
		filter *bson.D
		result *mongo.DeleteResult
		err    error
	)
	if req.Id == "" {
		return nil, guruErrors.ErrorIdValueMissing("主键ID的值为空")
	}
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, guruErrors.ErrorIdValueInvalid("主键ID的值 %s 不是合法的 ObjectID Hex 字符串: ", req.Id)
	}
	filter, err = filters.NewObjectIdFilter(objectID)
	if err != nil {
		return nil, err
	}
	result, err = s.collection.DeleteOne(ctx, filter, nil)
	if err != nil {
		s.helper.WithContext(ctx).Error(err)
		err = errors.Wrap(err, fmt.Sprintf("删除记录失败: %v", req))
		return nil, err
	}
	deleteCount := result.DeletedCount
	message := "删除记录成功"
	if int(deleteCount) == 0 {
		message = "找不到要删除的记录"
	}
	return &p.VideoCollectionDeleteRes{
		Message:      message,
		DeletedCount: deleteCount,
	}, nil
}