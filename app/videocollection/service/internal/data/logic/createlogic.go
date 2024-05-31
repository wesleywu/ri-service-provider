package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/castbox/go-guru/pkg/goguru/types"
	"github.com/castbox/go-guru/pkg/util/mongodb/codecs"
	"github.com/castbox/go-guru/pkg/util/sqids"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateLogic struct {
	collection *mongo.Collection
	helper     *log.Helper
}

func NewCreateLogic(collection *mongo.Collection, helper *log.Helper) *CreateLogic {
	return &CreateLogic{
		collection: collection,
		helper:     helper,
	}
}

func (s *CreateLogic) Create(ctx context.Context, req *p.VideoCollectionCreateReq) (*p.VideoCollectionCreateRes, error) {
	req.CreatedAt = timestamppb.New(time.Now())
	req.UpdatedAt = timestamppb.New(time.Now())
	res, err := s.collection.InsertOne(ctx, req)
	if err != nil {
		s.helper.WithContext(ctx).Error(err)
		err = errors.Wrap(err, fmt.Sprintf("创建记录失败: %v", req))
		return nil, err
	}
	var insertedID string
	if res.InsertedID != nil {
		insertedID = res.InsertedID.(primitive.ObjectID).Hex()
		if codecs.UseObjectIDObfuscated {
			insertedID = sqids.EncodeObjectID(insertedID)
		}
	}
	return &p.VideoCollectionCreateRes{
		Message:       "创建记录成功",
		InsertedID:    types.Wrap(insertedID),
		InsertedCount: 1,
	}, err
}
