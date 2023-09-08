package api

import (
	"context"

	"github.com/wesleywu/gowing/rpc/proxy"
	m "github.com/wesleywu/ri-service-provider/api/video_collection/model"
	p "github.com/wesleywu/ri-service-provider/proto/video_collection"
)

type VideoCollectionApi struct {
	client *p.VideoCollectionClientImpl
}

func NewVideoCollectionApi(client *p.VideoCollectionClientImpl) *VideoCollectionApi {
	return &VideoCollectionApi{
		client: client,
	}
}

// Count 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionApi) Count(ctx context.Context, req *m.VideoCollectionCountReq) (res *m.VideoCollectionCountRes, err error) {
	return proxy.CallServiceMethod[
		*p.VideoCollectionCountReq,
		*p.VideoCollectionCountRes,
		*m.VideoCollectionCountReq,
		*m.VideoCollectionCountRes,
	](ctx, req, true, s.client.Count)
}

// List 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionApi) List(ctx context.Context, req *m.VideoCollectionListReq) (res *m.VideoCollectionListRes, err error) {
	return proxy.CallServiceMethod[
		*p.VideoCollectionListReq,
		*p.VideoCollectionListRes,
		*m.VideoCollectionListReq,
		*m.VideoCollectionListRes,
	](ctx, req, true, s.client.List)
}

// One 根据req指定的查询条件获取单条数据
// 支持排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionApi) One(ctx context.Context, req *m.VideoCollectionOneReq) (res *m.VideoCollectionOneRes, err error) {
	return proxy.CallServiceMethod[
		*p.VideoCollectionOneReq,
		*p.VideoCollectionOneRes,
		*m.VideoCollectionOneReq,
		*m.VideoCollectionOneRes,
	](ctx, req, true, s.client.One)
}

// Create 插入记录
// 包括表中所有字段，支持字段类型自动转换，支持对非主键且可为空字段不赋值
// 未赋值或赋值为nil的字段将被更新为 NULL 或数据库表指定的DEFAULT
func (s *VideoCollectionApi) Create(ctx context.Context, req *m.VideoCollectionCreateReq) (res *m.VideoCollectionCreateRes, err error) {
	return proxy.CallServiceMethod[
		*p.VideoCollectionCreateReq,
		*p.VideoCollectionCreateRes,
		*m.VideoCollectionCreateReq,
		*m.VideoCollectionCreateRes,
	](ctx, req, false, s.client.Create)
}

// Update 根据主键更新对应记录
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新（即不会修改原记录的字段值）
func (s *VideoCollectionApi) Update(ctx context.Context, req *m.VideoCollectionUpdateReq) (res *m.VideoCollectionUpdateRes, err error) {
	return proxy.CallServiceMethod[
		*p.VideoCollectionUpdateReq,
		*p.VideoCollectionUpdateRes,
		*m.VideoCollectionUpdateReq,
		*m.VideoCollectionUpdateRes,
	](ctx, req, false, s.client.Update)
}

// Upsert 根据主键（或唯一索引）是否存在且已在req中赋值，更新或插入对应记录。
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新/插入（即更新时不会修改原记录的字段值）
func (s *VideoCollectionApi) Upsert(ctx context.Context, req *m.VideoCollectionUpsertReq) (res *m.VideoCollectionUpsertRes, err error) {
	return proxy.CallServiceMethod[
		*p.VideoCollectionUpsertReq,
		*p.VideoCollectionUpsertRes,
		*m.VideoCollectionUpsertReq,
		*m.VideoCollectionUpsertRes,
	](ctx, req, false, s.client.Upsert)
}

// Delete 根据req指定的条件删除表中记录
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionApi) Delete(ctx context.Context, req *m.VideoCollectionDeleteReq) (res *m.VideoCollectionDeleteRes, err error) {
	return proxy.CallServiceMethod[
		*p.VideoCollectionDeleteReq,
		*p.VideoCollectionDeleteRes,
		*m.VideoCollectionDeleteReq,
		*m.VideoCollectionDeleteRes,
	](ctx, req, false, s.client.Delete)
}
