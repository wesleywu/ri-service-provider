package logic

import (
	"context"
	"errors"
	"fmt"

	guruErrors "github.com/castbox/go-guru/pkg/goguru/error"
	"github.com/castbox/go-guru/pkg/infra/mongodb/filters"
	"github.com/castbox/go-guru/pkg/util/sqids"
	"github.com/go-kratos/kratos/v2/log"
	pkgErrors "github.com/pkg/errors"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetLogic struct {
	collection       *mongo.Collection
	helper           *log.Helper
	useIdObfuscating bool
}

func NewGetLogic(collection *mongo.Collection, useIdObfuscating bool, helper *log.Helper) *GetLogic {
	return &GetLogic{
		collection:       collection,
		helper:           helper,
		useIdObfuscating: useIdObfuscating,
	}
}

func (s *GetLogic) Get(ctx context.Context, req *p.VideoCollectionGetReq) (*p.VideoCollectionGetRes, error) {
	var (
		reqID  string
		filter *bson.D
		err    error
		item   *p.VideoCollectionGetRes
	)
	reqID = req.Id
	if reqID == "" {
		return nil, guruErrors.ErrorIdValueMissing("主键ID的值为空")
	}
	if s.useIdObfuscating {
		reqID, err = sqids.DecodeObjectID(reqID)
		if err != nil {
			return nil, err
		}
	}
	objectID, err := primitive.ObjectIDFromHex(reqID)
	if err != nil {
		return nil, guruErrors.ErrorIdValueInvalid("主键ID的值 %s 不是合法的 ObjectID Hex 字符串: ", req.Id)
	}
	filter, err = filters.NewObjectIdFilter(objectID)
	if err != nil {
		return nil, err
	}

	singleResult := s.collection.FindOne(ctx, filter, nil)
	if err = singleResult.Err(); err != nil {
		if errors.Is(singleResult.Err(), mongo.ErrNoDocuments) {
			return nil, guruErrors.ErrorRecordNotFound("找不到ID为 %s 的记录", req.Id)
		}
		s.helper.WithContext(ctx).Error(err)
		err = pkgErrors.Wrap(err, fmt.Sprintf("主键查询失败: %v", req))
		return nil, err
	}
	item = (*p.VideoCollectionGetRes)(nil)
	err = singleResult.Decode(&item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
