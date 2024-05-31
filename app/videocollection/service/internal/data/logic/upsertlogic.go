package logic

import (
	"context"
	"fmt"
	"time"

	guruErrors "github.com/castbox/go-guru/pkg/goguru/error"
	"github.com/castbox/go-guru/pkg/goguru/types"
	"github.com/castbox/go-guru/pkg/util/mongodb/codecs"
	"github.com/castbox/go-guru/pkg/util/mongodb/filters"
	"github.com/castbox/go-guru/pkg/util/sqids"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UpsertLogic struct {
	collection *mongo.Collection
	helper     *log.Helper
}

func NewUpsertLogic(collection *mongo.Collection, helper *log.Helper) *UpsertLogic {
	return &UpsertLogic{
		collection: collection,
		helper:     helper,
	}
}

func (s *UpsertLogic) Upsert(ctx context.Context, req *p.VideoCollectionUpsertReq) (*p.VideoCollectionUpsertRes, error) {
	var (
		filter       *bson.D
		updateResult *mongo.UpdateResult
		upsertedId   string
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
	opts := options.Update().SetUpsert(true)
	updateResult, err = s.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		s.helper.WithContext(ctx).Error(err)
		err = errors.Wrap(err, fmt.Sprintf("插入或更新记录失败: %v", req))
		return nil, err
	}

	message := "更新记录成功"
	if updateResult.UpsertedID != nil {
		upsertedId = updateResult.UpsertedID.(primitive.ObjectID).Hex()
		if codecs.UseObjectIDObfuscated {
			upsertedId = sqids.EncodeObjectID(upsertedId)
		}
		message = "插入记录成功"
	}
	return &p.VideoCollectionUpsertRes{
		Message:       message,
		UpsertedID:    types.Wrap(upsertedId),
		MatchedCount:  updateResult.MatchedCount,
		ModifiedCount: updateResult.ModifiedCount,
		UpsertedCount: updateResult.UpsertedCount,
	}, err
}
