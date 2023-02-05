package api

import (
	"context"
	"github.com/WesleyWu/ri-service-provider/api/video_collection/model"
	proto "github.com/WesleyWu/ri-service-provider/proto/video_collection"
)

type VideoCollectionApi struct {
	client *proto.VideoCollectionClientImpl
}

func NewVideoCollectionApi(client *proto.VideoCollectionClientImpl) *VideoCollectionApi {
	return &VideoCollectionApi{
		client: client,
	}
}

// Count 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionApi) Count(ctx context.Context, req *model.VideoCollectionCountReq) (res *model.VideoCollectionCountRes, err error) {
	var (
		in  *proto.VideoCollectionCountReq
		out *proto.VideoCollectionCountRes
	)
	in = req.ToMessage(ctx)
	out, err = s.client.Count(ctx, in)
	if err != nil {
		return nil, err
	}
	res = &model.VideoCollectionCountRes{}
	res.FromMessage(ctx, out)
	return
}

// List 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionApi) List(ctx context.Context, req *model.VideoCollectionListReq) (res *model.VideoCollectionListRes, err error) {
	var (
		in  *proto.VideoCollectionListReq
		out *proto.VideoCollectionListRes
	)
	in = req.ToMessage(ctx)
	if err != nil {
		return nil, err
	}
	out, err = s.client.List(ctx, in)
	if err != nil {
		return nil, err
	}
	res = &model.VideoCollectionListRes{}
	res.FromMessage(ctx, out)
	return
}

// One 根据req指定的查询条件获取单条数据
// 支持排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionApi) One(ctx context.Context, req *model.VideoCollectionOneReq) (res *model.VideoCollectionOneRes, err error) {
	var (
		in  *proto.VideoCollectionOneReq
		out *proto.VideoCollectionOneRes
	)
	in = req.ToMessage(ctx)
	out, err = s.client.One(ctx, in)
	if err != nil {
		return nil, err
	}
	res = &model.VideoCollectionOneRes{}
	res.FromMessage(ctx, out)
	return
}

// Create 插入记录
// 包括表中所有字段，支持字段类型自动转换，支持对非主键且可为空字段不赋值
// 未赋值或赋值为nil的字段将被更新为 NULL 或数据库表指定的DEFAULT
func (s *VideoCollectionApi) Create(ctx context.Context, req *model.VideoCollectionCreateReq) (res *model.VideoCollectionCreateRes, err error) {
	var (
		in  *proto.VideoCollectionCreateReq
		out *proto.VideoCollectionCreateRes
	)
	in = req.ToMessage(ctx)
	out, err = s.client.Create(ctx, in)
	if err != nil {
		return nil, err
	}
	res = &model.VideoCollectionCreateRes{}
	res.FromMessage(ctx, out)
	return
}

// Update 根据主键更新对应记录
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新（即不会修改原记录的字段值）
func (s *VideoCollectionApi) Update(ctx context.Context, req *model.VideoCollectionUpdateReq) (res *model.VideoCollectionUpdateRes, err error) {
	var (
		in  *proto.VideoCollectionUpdateReq
		out *proto.VideoCollectionUpdateRes
	)
	in = req.ToMessage(ctx)
	out, err = s.client.Update(ctx, in)
	if err != nil {
		return nil, err
	}
	res = &model.VideoCollectionUpdateRes{}
	res.FromMessage(ctx, out)
	return
}

// Upsert 根据主键（或唯一索引）是否存在且已在req中赋值，更新或插入对应记录。
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新/插入（即更新时不会修改原记录的字段值）
func (s *VideoCollectionApi) Upsert(ctx context.Context, req *model.VideoCollectionUpsertReq) (res *model.VideoCollectionUpsertRes, err error) {
	var (
		in  *proto.VideoCollectionUpsertReq
		out *proto.VideoCollectionUpsertRes
	)
	in = req.ToMessage(ctx)
	out, err = s.client.Upsert(ctx, in)
	if err != nil {
		return nil, err
	}
	res = &model.VideoCollectionUpsertRes{}
	res.FromMessage(ctx, out)
	return
}

// Delete 根据req指定的条件删除表中记录
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionApi) Delete(ctx context.Context, req *model.VideoCollectionDeleteReq) (res *model.VideoCollectionDeleteRes, err error) {
	var (
		in  *proto.VideoCollectionDeleteReq
		out *proto.VideoCollectionDeleteRes
	)
	in = req.ToMessage(ctx)
	out, err = s.client.Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	res = &model.VideoCollectionDeleteRes{}
	res.FromMessage(ctx, out)
	return
}
