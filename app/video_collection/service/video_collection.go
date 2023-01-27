package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/WesleyWu/ri-service-provider/app/video_collection/model"
	"github.com/WesleyWu/ri-service-provider/app/video_collection/service/internal/dao"
	"github.com/WesleyWu/ri-service-provider/gowing/util/errors"
	"github.com/WesleyWu/ri-service-provider/gowing/util/orm"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type IVideoCollection interface {
	Count(ctx context.Context, req *model.VideoCollectionCountReq) (*model.VideoCollectionCountRes, error)
	One(ctx context.Context, req *model.VideoCollectionOneReq) (*model.VideoCollectionOneRes, error)
	List(ctx context.Context, req *model.VideoCollectionListReq) (*model.VideoCollectionListRes, error)
	Create(ctx context.Context, req *model.VideoCollectionCreateReq) (*model.VideoCollectionCreateRes, error)
	Update(ctx context.Context, req *model.VideoCollectionUpdateReq) (*model.VideoCollectionUpdateRes, error)
	Upsert(ctx context.Context, req *model.VideoCollectionUpsertReq) (*model.VideoCollectionUpsertRes, error)
	Delete(ctx context.Context, req *model.VideoCollectionDeleteReq) (*model.VideoCollectionDeleteRes, error)
}

type VideoCollectionImpl struct {
	model.UnimplementedVideoCollectionServer
}

var (
	VideoCollection IVideoCollection = new(VideoCollectionImpl)
)

// Count 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionImpl) Count(ctx context.Context, req *model.VideoCollectionCountReq) (*model.VideoCollectionCountRes, error) {
	var err error
	m := dao.VideoCollection.Ctx(ctx).WithAll()
	m, err = orm.ParseConditions(ctx, req, dao.VideoCollection.ColumnMap, m)
	if err != nil {
		return nil, err
	}
	count, err := m.Count()
	if err != nil {
		return nil, err
	}
	return &model.VideoCollectionCountRes{Total: gconv.Int32(count)}, err
}

// List 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionImpl) List(ctx context.Context, req *model.VideoCollectionListReq) (*model.VideoCollectionListRes, error) {
	var (
		total int
		page  int
		order string
		list  []*model.VideoCollectionItem
		err   error
	)
	m := dao.VideoCollection.Ctx(ctx).WithAll()
	m, err = orm.ParseConditions(ctx, req, dao.VideoCollection.ColumnMap, m)
	if err != nil {
		return nil, err
	}
	total, err = m.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		err = errors.WrapServiceErrorf(err, req, "获取数据总记录数失败")
		return nil, err
	}
	if req.Page == 0 {
		req.Page = 1
	}
	page = int(req.Page)
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	if !g.IsEmpty(req.OrderBy) {
		order = req.OrderBy
	}
	list = []*model.VideoCollectionItem{}
	err = m.Fields(model.VideoCollectionItem{}).Page(page, int(req.PageSize)).Order(order).Scan(&list)
	if err != nil {
		g.Log().Error(ctx, err)
		err = errors.WrapServiceErrorf(err, req, "获取数据列表失败")
		return nil, err
	}
	return &model.VideoCollectionListRes{
		Total:   int32(total),
		Current: int32(page),
		Items:   list,
	}, nil
}

// One 根据req指定的查询条件获取单条数据
// 支持排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionImpl) One(ctx context.Context, req *model.VideoCollectionOneReq) (*model.VideoCollectionOneRes, error) {
	var (
		list  []*model.VideoCollectionItem
		order string
		err   error
	)
	m := dao.VideoCollection.Ctx(ctx).WithAll()
	m, err = orm.ParseConditions(ctx, req, dao.VideoCollection.ColumnMap, m)
	if err != nil {
		return nil, err
	}
	if !g.IsEmpty(req.OrderBy) {
		order = req.OrderBy
	}
	err = m.Fields(model.VideoCollectionItem{}).Order(order).Limit(1).Scan(&list)
	if err != nil {
		g.Log().Error(ctx, err)
		err = errors.WrapServiceErrorf(err, req, "获取单条数据记录失败")
		return nil, err
	}
	if g.IsEmpty(list) || len(list) == 0 {
		return nil, errors.NewNotFoundErrorf(req, "找不到要获取的数据")
	}
	v := list[0]
	return &model.VideoCollectionOneRes{
		Id:          v.Id,
		Name:        v.Name,
		ContentType: v.ContentType,
		FilterType:  v.FilterType,
		Count:       v.Count,
		IsOnline:    v.IsOnline,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}, nil
}

// Create 插入记录
// 包括表中所有字段，支持字段类型自动转换，支持对非主键且可为空字段不赋值
// 未赋值或赋值为nil的字段将被更新为 NULL 或数据库表指定的DEFAULT
func (s *VideoCollectionImpl) Create(ctx context.Context, req *model.VideoCollectionCreateReq) (*model.VideoCollectionCreateRes, error) {
	var (
		result       sql.Result
		lastInsertId int64
		rowsAffected int64
		err          error
	)
	result, err = dao.VideoCollection.Ctx(ctx).Insert(req)
	if err != nil {
		if reqErr, ok := errors.DbErrorToRequestError(req, err, dao.VideoCollectionDbType); ok {
			return nil, reqErr
		}
		err = errors.WrapServiceErrorf(err, req, "插入失败")
		g.Log().Error(ctx, err)
		return nil, err
	}
	lastInsertId, err = result.LastInsertId()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "获取插入记录主键键值出错")
		g.Log().Error(ctx, err)
		return nil, err
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "获取插入记录条数出错")
		g.Log().Error(ctx, err)
		return nil, err
	}
	message := "插入成功"
	if rowsAffected == 0 {
		message = "未插入任何记录" // should not happen
	}
	return &model.VideoCollectionCreateRes{
		Message:      message,
		LastInsertId: lastInsertId,
		RowsAffected: rowsAffected,
	}, nil
}

// Update 根据主键更新对应记录
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新（即不会修改原记录的字段值）
func (s *VideoCollectionImpl) Update(ctx context.Context, req *model.VideoCollectionUpdateReq) (*model.VideoCollectionUpdateRes, error) {
	var (
		result       sql.Result
		rowsAffected int64
		err          error
	)
	result, err = dao.VideoCollection.Ctx(ctx).
		FieldsEx(dao.VideoCollection.Columns.Id,
			dao.VideoCollection.Columns.CreatedAt).
		WherePri(req.Id).
		Update(req)
	if err != nil {
		if reqErr, ok := errors.DbErrorToRequestError(req, err, dao.VideoCollectionDbType); ok {
			return nil, reqErr
		}
		err = errors.WrapServiceErrorf(err, req, "更新失败")
		g.Log().Error(ctx, err)
		return nil, err
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "获取插入记录条数出错")
		return nil, err
	}
	message := "更新成功"
	if rowsAffected == 0 {
		return nil, errors.NewNotFoundErrorf(req, "不存在要更新的记录")
	}
	return &model.VideoCollectionUpdateRes{
		Message:      message,
		RowsAffected: rowsAffected,
	}, nil
}

// Upsert 根据主键（或唯一索引）是否存在且已在req中赋值，更新或插入对应记录。
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新/插入（即更新时不会修改原记录的字段值）
func (s *VideoCollectionImpl) Upsert(ctx context.Context, req *model.VideoCollectionUpsertReq) (*model.VideoCollectionUpsertRes, error) {
	var (
		result       sql.Result
		lastInsertId int64
		rowsAffected int64
		err          error
	)
	result, err = dao.VideoCollection.Ctx(ctx).FieldsEx(dao.VideoCollection.Columns.CreatedAt).Data(req).Save()
	if err != nil {
		if reqErr, ok := errors.DbErrorToRequestError(req, err, dao.VideoCollectionDbType); ok {
			return nil, reqErr
		}
		err = errors.WrapServiceErrorf(err, req, "插入/更新失败")
		g.Log().Error(ctx, err)
		return nil, err
	}
	lastInsertId, err = result.LastInsertId()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "获取插入记录主键键值出错")
		g.Log().Error(ctx, err)
		return nil, err
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "获取插入/更新记录条数出错")
		g.Log().Error(ctx, err)
		return nil, err
	}
	message := "更新成功"
	if rowsAffected == 0 {
		message = "未插入/更新任何记录" // should not happen
	} else if rowsAffected == 1 {
		message = "插入成功"
	}
	return &model.VideoCollectionUpsertRes{
		Message:      message,
		LastInsertId: lastInsertId,
		RowsAffected: rowsAffected,
	}, nil
}

// Delete 根据req指定的条件删除表中记录
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionImpl) Delete(ctx context.Context, req *model.VideoCollectionDeleteReq) (*model.VideoCollectionDeleteRes, error) {
	var (
		result       sql.Result
		rowsAffected int64
		err          error
	)
	m := dao.VideoCollection.Ctx(ctx).WithAll()
	//if req.IsEmpty() {
	//	return nil, errors.NewBadRequestErrorf(req, "必须指定删除条件")
	//}
	m, err = orm.ParseConditions(ctx, req, dao.VideoCollection.ColumnMap, m)
	if err != nil {
		return nil, err
	}
	result, err = m.Delete()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "删除失败")
		g.Log().Error(ctx, err)
		return nil, err
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		err = errors.WrapServiceErrorf(err, req, "获取删除记录条数出错")
		return nil, err
	}
	message := fmt.Sprintf("已删除%d条记录", rowsAffected)
	if rowsAffected == 0 {
		message = "不存在要删除的记录"
	}

	return &model.VideoCollectionDeleteRes{
		Message:      message,
		RowsAffected: rowsAffected,
	}, nil
}
