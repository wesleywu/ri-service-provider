package video_collection

import (
	"context"

	"github.com/castbox/go-guru/pkg/util/appinfo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wesleywu/ri-service-provider/api/internal/client"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
)

type IVideoCollection interface {
	Count(ctx context.Context, req *p.VideoCollectionCountReq) (*p.VideoCollectionCountRes, error)
	One(ctx context.Context, req *p.VideoCollectionOneReq) (*p.VideoCollectionOneRes, error)
	List(ctx context.Context, req *p.VideoCollectionListReq) (*p.VideoCollectionListRes, error)
	Create(ctx context.Context, req *p.VideoCollectionCreateReq) (*p.VideoCollectionCreateRes, error)
	Update(ctx context.Context, req *p.VideoCollectionUpdateReq) (*p.VideoCollectionUpdateRes, error)
	Upsert(ctx context.Context, req *p.VideoCollectionUpsertReq) (*p.VideoCollectionUpsertRes, error)
	Delete(ctx context.Context, req *p.VideoCollectionDeleteReq) (*p.VideoCollectionDeleteRes, error)
}

type ServiceImpl struct {
	p.UnimplementedVideoCollectionServer
	metadata *appinfo.AppMetadata
	helper   *log.Helper
	clients  *client.Clients
}

func NewVideoCollectionService(metadata *appinfo.AppMetadata, helper *log.Helper, clients *client.Clients) *ServiceImpl {
	return &ServiceImpl{
		metadata: metadata,
		helper:   helper,
		clients:  clients,
	}
}

// Count 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *ServiceImpl) Count(ctx context.Context, req *p.VideoCollectionCountReq) (*p.VideoCollectionCountRes, error) {
	return s.clients.VideoCollection.Count(ctx, req)
}

// List 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *ServiceImpl) List(ctx context.Context, req *p.VideoCollectionListReq) (*p.VideoCollectionListRes, error) {
	return s.clients.VideoCollection.List(ctx, req)
}

// One 根据req指定的查询条件获取单条数据
// 支持排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *ServiceImpl) One(ctx context.Context, req *p.VideoCollectionOneReq) (*p.VideoCollectionOneRes, error) {
	return s.clients.VideoCollection.One(ctx, req)
}

// Create 插入记录
// 包括表中所有字段，支持字段类型自动转换，支持对非主键且可为空字段不赋值
// 未赋值或赋值为nil的字段将被更新为 NULL 或数据库表指定的DEFAULT
func (s *ServiceImpl) Create(ctx context.Context, req *p.VideoCollectionCreateReq) (*p.VideoCollectionCreateRes, error) {
	return s.clients.VideoCollection.Create(ctx, req)
}

// Update 根据主键更新对应记录
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新（即不会修改原记录的字段值）
func (s *ServiceImpl) Update(ctx context.Context, req *p.VideoCollectionUpdateReq) (*p.VideoCollectionUpdateRes, error) {
	return s.clients.VideoCollection.Update(ctx, req)
}

// Upsert 根据主键（或唯一索引）是否存在且已在req中赋值，更新或插入对应记录。
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新/插入（即更新时不会修改原记录的字段值）
func (s *ServiceImpl) Upsert(ctx context.Context, req *p.VideoCollectionUpsertReq) (*p.VideoCollectionUpsertRes, error) {
	return s.clients.VideoCollection.Upsert(ctx, req)
}

// Delete 根据req指定的条件删除表中记录
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *ServiceImpl) Delete(ctx context.Context, req *p.VideoCollectionDeleteReq) (*p.VideoCollectionDeleteRes, error) {
	return s.clients.VideoCollection.Delete(ctx, req)
}
