package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/wesleywu/ri-service-provider/api/episode/service/v1"
	p "github.com/wesleywu/ri-service-provider/app/episode/service/proto"
)

type Episode struct {
	v1.UnimplementedEpisodeServer
	repo   *p.EpisodeRepo
	helper *log.Helper
}

func NewEpisodeService(repo *p.EpisodeRepo, helper *log.Helper) *Episode {
	return &Episode{
		repo:   repo,
		helper: helper,
	}
}

// Count 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *Episode) Count(ctx context.Context, req *p.EpisodeCountReq) (*p.EpisodeCountRes, error) {
	return s.repo.Count(ctx, req)
}

// List 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *Episode) List(ctx context.Context, req *p.EpisodeListReq) (*p.EpisodeListRes, error) {
	return s.repo.List(ctx, req)
}

// One 根据req指定的查询条件获取单条数据
// 支持排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *Episode) One(ctx context.Context, req *p.EpisodeOneReq) (*p.EpisodeOneRes, error) {
	return s.repo.One(ctx, req)
}

// Get 根据主键/ID查询特定记录
func (s *Episode) Get(ctx context.Context, req *p.EpisodeGetReq) (*p.EpisodeGetRes, error) {
	return s.repo.Get(ctx, req)
}

// Create 插入记录
// 包括表中所有字段，支持字段类型自动转换，支持对非主键且可为空字段不赋值
// 未赋值或赋值为nil的字段将被更新为 NULL 或数据库表指定的DEFAULT
func (s *Episode) Create(ctx context.Context, req *p.EpisodeCreateReq) (*p.EpisodeCreateRes, error) {
	return s.repo.Create(ctx, req)
}

// Update 根据主键更新对应记录
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新（即不会修改原记录的字段值）
func (s *Episode) Update(ctx context.Context, req *p.EpisodeUpdateReq) (*p.EpisodeUpdateRes, error) {
	return s.repo.Update(ctx, req)
}

// Upsert 根据主键（或唯一索引）是否存在且已在req中赋值，更新或插入对应记录。
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新/插入（即更新时不会修改原记录的字段值）
func (s *Episode) Upsert(ctx context.Context, req *p.EpisodeUpsertReq) (*p.EpisodeUpsertRes, error) {
	return s.repo.Upsert(ctx, req)
}

// Delete 根据主键删除对应记录
func (s *Episode) Delete(ctx context.Context, req *p.EpisodeDeleteReq) (*p.EpisodeDeleteRes, error) {
	return s.repo.Delete(ctx, req)
}

// DeleteMulti 根据req指定的条件删除表中记录（可能多条）
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *Episode) DeleteMulti(ctx context.Context, req *p.EpisodeDeleteMultiReq) (*p.EpisodeDeleteMultiRes, error) {
	return s.repo.DeleteMulti(ctx, req)
}
