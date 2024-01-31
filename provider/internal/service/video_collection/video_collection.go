package video_collection

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/provider/internal/data"
	"github.com/wesleywu/ri-service-provider/provider/internal/service/video_collection/logic"
	"go.mongodb.org/mongo-driver/mongo"
)

type IVideoCollection interface {
	Count(ctx context.Context, req *p.VideoCollectionCountReq) (*p.VideoCollectionCountRes, error)
	One(ctx context.Context, req *p.VideoCollectionOneReq) (*p.VideoCollectionOneRes, error)
	List(ctx context.Context, req *p.VideoCollectionListReq) (*p.VideoCollectionListRes, error)
	Get(ctx context.Context, req *p.VideoCollectionGetReq) (*p.VideoCollectionGetRes, error)
	Create(ctx context.Context, req *p.VideoCollectionCreateReq) (*p.VideoCollectionCreateRes, error)
	Update(ctx context.Context, req *p.VideoCollectionUpdateReq) (*p.VideoCollectionUpdateRes, error)
	Upsert(ctx context.Context, req *p.VideoCollectionUpsertReq) (*p.VideoCollectionUpsertRes, error)
	Delete(ctx context.Context, req *p.VideoCollectionDeleteReq) (*p.VideoCollectionDeleteRes, error)
	DeleteMulti(ctx context.Context, req *p.VideoCollectionDeleteMultiReq) (*p.VideoCollectionDeleteMultiRes, error)
}

type ServiceImpl struct {
	p.UnimplementedVideoCollectionServer
	metadata   *appinfo.AppMetadata
	helper     *log.Helper
	collection *mongo.Collection
}

func NewVideoCollectionService(metadata *appinfo.AppMetadata, helper *log.Helper, data *data.Data) *ServiceImpl {
	collection := data.MongoDatabase.Collection("video_collection")
	return &ServiceImpl{
		metadata:   metadata,
		helper:     helper,
		collection: collection,
	}
}

// Count 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *ServiceImpl) Count(ctx context.Context, req *p.VideoCollectionCountReq) (*p.VideoCollectionCountRes, error) {
	l := logic.NewCountLogic(s.metadata, s.helper, s.collection)
	return l.Count(ctx, req)
}

// List 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *ServiceImpl) List(ctx context.Context, req *p.VideoCollectionListReq) (*p.VideoCollectionListRes, error) {
	l := logic.NewListLogic(s.metadata, s.helper, s.collection)
	return l.List(ctx, req)
}

// One 根据req指定的查询条件获取单条数据
// 支持排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *ServiceImpl) One(ctx context.Context, req *p.VideoCollectionOneReq) (*p.VideoCollectionOneRes, error) {
	l := logic.NewOneLogic(s.metadata, s.helper, s.collection)
	return l.One(ctx, req)
}

// Get 根据主键/ID查询特定记录
func (s *ServiceImpl) Get(ctx context.Context, req *p.VideoCollectionGetReq) (*p.VideoCollectionGetRes, error) {
	l := logic.NewGetLogic(s.metadata, s.helper, s.collection)
	return l.Get(ctx, req)
}

// Create 插入记录
// 包括表中所有字段，支持字段类型自动转换，支持对非主键且可为空字段不赋值
// 未赋值或赋值为nil的字段将被更新为 NULL 或数据库表指定的DEFAULT
func (s *ServiceImpl) Create(ctx context.Context, req *p.VideoCollectionCreateReq) (*p.VideoCollectionCreateRes, error) {
	l := logic.NewCreateLogic(s.metadata, s.helper, s.collection)
	return l.Create(ctx, req)
}

// Update 根据主键更新对应记录
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新（即不会修改原记录的字段值）
func (s *ServiceImpl) Update(ctx context.Context, req *p.VideoCollectionUpdateReq) (*p.VideoCollectionUpdateRes, error) {
	l := logic.NewUpdateLogic(s.metadata, s.helper, s.collection)
	return l.Update(ctx, req)
}

// Upsert 根据主键（或唯一索引）是否存在且已在req中赋值，更新或插入对应记录。
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新/插入（即更新时不会修改原记录的字段值）
func (s *ServiceImpl) Upsert(ctx context.Context, req *p.VideoCollectionUpsertReq) (*p.VideoCollectionUpsertRes, error) {
	l := logic.NewUpsertLogic(s.metadata, s.helper, s.collection)
	return l.Upsert(ctx, req)
}

// Delete 根据主键删除对应记录
func (s *ServiceImpl) Delete(ctx context.Context, req *p.VideoCollectionDeleteReq) (*p.VideoCollectionDeleteRes, error) {
	l := logic.NewDeleteLogic(s.metadata, s.helper, s.collection)
	return l.Delete(ctx, req)
}

// DeleteMulti 根据req指定的条件删除表中记录（可能多条）
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *ServiceImpl) DeleteMulti(ctx context.Context, req *p.VideoCollectionDeleteMultiReq) (*p.VideoCollectionDeleteMultiRes, error) {
	l := logic.NewDeleteMultiLogic(s.metadata, s.helper, s.collection)
	return l.DeleteMulti(ctx, req)
}
