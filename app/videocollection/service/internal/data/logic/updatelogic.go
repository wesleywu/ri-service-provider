package logic

import (
	"context"
	"fmt"
	"time"

	guruErrors "github.com/castbox/go-guru/pkg/goguru/error"
	"github.com/castbox/go-guru/pkg/util/mongodb/filters"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UpdateLogic struct {
	collection *mongo.Collection
	helper     *log.Helper
}

func NewUpdateLogic(collection *mongo.Collection, helper *log.Helper) *UpdateLogic {
	return &UpdateLogic{
		collection: collection,
		helper:     helper,
	}
}

func (s *UpdateLogic) Update(ctx context.Context, req *p.VideoCollectionUpdateReq) (*p.VideoCollectionUpdateRes, error) {
	var (
		filter       *bson.D
		updateResult *mongo.UpdateResult
		err          error
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

	req.UpdatedAt = timestamppb.New(time.Now())
	update := bson.D{
		{
			"$set", req,
		},
	}
	updateResult, err = s.collection.UpdateOne(ctx, filter, update, nil)
	if err != nil {
		s.helper.WithContext(ctx).Error(err)
		err = errors.Wrap(err, fmt.Sprintf("更新记录失败: %v", req))
		return nil, err
	}
	return &p.VideoCollectionUpdateRes{
		Message:       "更新记录成功",
		MatchedCount:  updateResult.MatchedCount,
		ModifiedCount: updateResult.ModifiedCount,
	}, err
}
