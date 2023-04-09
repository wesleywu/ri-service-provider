package service

import (
	"context"
	"fmt"
	"github.com/WesleyWu/gowing/errors/gwerror"
	"github.com/WesleyWu/gowing/util/gworm"
	"github.com/WesleyWu/gowing/util/gwwrapper"
	p "github.com/WesleyWu/ri-service-provider/proto/video_collection"
	"github.com/WesleyWu/ri-service-provider/provider/video_collection_mongo/service/internal/dao"
	"github.com/gogf/gf/v2/frame/g"
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

type VideoCollectionImpl struct {
	p.UnimplementedVideoCollectionServer
}

var (
	VideoCollection IVideoCollection = new(VideoCollectionImpl)
)

// Count 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionImpl) Count(ctx context.Context, req *p.VideoCollectionCountReq) (*p.VideoCollectionCountRes, error) {
	var err error
	//m := dao.VideoCollection.Ctx(ctx)
	//m, err = gworm.ParseConditions(ctx, req, dao.VideoCollection.ColumnMap, m)
	//if err != nil {
	//	return nil, err
	//}
	//count, err := m.Count()
	//if err != nil {
	//	return nil, err
	//}
	return &p.VideoCollectionCountRes{Total: gwwrapper.WrapInt64(int64(0))}, err
}

// List 根据req指定的查询条件获取记录列表
// 支持翻页和排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionImpl) List(ctx context.Context, req *p.VideoCollectionListReq) (*p.VideoCollectionListRes, error) {
	var (
		total int64
		page  int
		//order string
		list []*p.VideoCollectionItem
		//err   error
	)
	//m := dao.VideoCollection.Ctx(ctx)
	//m, err = gworm.ParseConditions(ctx, req, dao.VideoCollection.ColumnMap, m)
	//if err != nil {
	//	return nil, err
	//}
	//total, err = m.Count()
	//if err != nil {
	//	g.Log().Error(ctx, err)
	//	err = gwerror.WrapServiceErrorf(err, req, "获取数据总记录数失败")
	//	return nil, err
	//}
	//if req.Page == 0 {
	//	req.Page = 1
	//}
	//page = int(req.Page)
	//if req.PageSize == 0 {
	//	req.PageSize = 10
	//}
	//if !g.IsEmpty(req.OrderBy) {
	//	order = req.OrderBy
	//}
	//list = []*p.VideoCollectionItem{}
	//err = m.Fields(p.VideoCollectionItem{}).Page(page, int(req.PageSize)).Order(order).Scan(&list)
	//if err != nil {
	//	g.Log().Error(ctx, err)
	//	err = gwerror.WrapServiceErrorf(err, req, "获取数据列表失败")
	//	return nil, err
	//}
	return &p.VideoCollectionListRes{
		Total:   gwwrapper.WrapInt64(total),
		Current: gwwrapper.WrapUInt32(uint32(page)),
		Items:   list,
	}, nil
}

// One 根据req指定的查询条件获取单条数据
// 支持排序参数，支持查询条件参数类型自动转换
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionImpl) One(ctx context.Context, req *p.VideoCollectionOneReq) (*p.VideoCollectionOneRes, error) {
	var (
		list []*p.VideoCollectionItem
		//order string
		//err   error
	)
	//m := dao.VideoCollection.Ctx(ctx)
	//m, err = gworm.ParseConditions(ctx, req, dao.VideoCollection.ColumnMap, m)
	//if err != nil {
	//	return nil, err
	//}
	//if !g.IsEmpty(req.OrderBy) {
	//	order = req.OrderBy
	//}
	//err = m.Fields(p.VideoCollectionItem{}).Order(order).Limit(1).Scan(&list)
	//if err != nil {
	//	g.Log().Error(ctx, err)
	//	err = gwerror.WrapServiceErrorf(err, req, "获取单条数据记录失败")
	//	return nil, err
	//}
	//if g.IsEmpty(list) || len(list) == 0 {
	//	return nil, gwerror.NewNotFoundErrorf(req, "找不到要获取的数据")
	//}
	v := list[0]
	return &p.VideoCollectionOneRes{
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
func (s *VideoCollectionImpl) Create(ctx context.Context, req *p.VideoCollectionCreateReq) (*p.VideoCollectionCreateRes, error) {
	var (
		result       *gworm.Result
		insertedId   string
		rowsAffected int64
		err          error
	)
	result, err = dao.VideoCollection.Ctx(ctx).InsertOne(ctx, req)
	if err != nil {
		if reqErr, ok := gwerror.DbErrorToRequestError(req, err, dao.VideoCollectionDbType); ok {
			return nil, reqErr
		}
		err = gwerror.WrapServiceErrorf(err, req, "插入失败")
		g.Log().Error(ctx, err)
		return nil, err
	}
	insertedId, err = result.InsertedId()
	if err != nil {
		err = gwerror.WrapServiceErrorf(err, req, "获取插入记录条数出错")
		g.Log().Error(ctx, err)
		return nil, err
	}
	message := "插入成功"
	if insertedId == "" {
		message = "未插入任何记录" // should not happen
	} else {
		rowsAffected = 1
	}
	return &p.VideoCollectionCreateRes{
		Message:      gwwrapper.WrapString(message),
		InsertedId:   gwwrapper.WrapString(insertedId),
		RowsAffected: gwwrapper.WrapInt64(rowsAffected),
	}, nil
}

// Update 根据主键更新对应记录
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新（即不会修改原记录的字段值）
func (s *VideoCollectionImpl) Update(ctx context.Context, req *p.VideoCollectionUpdateReq) (*p.VideoCollectionUpdateRes, error) {
	var (
		result       *gworm.Result
		rowsAffected int64
		err          error
	)
	result, err = dao.VideoCollection.Ctx(ctx).
		FieldsEx(dao.VideoCollection.Columns.Id,
			dao.VideoCollection.Columns.CreatedAt).
		WherePri(req.Id).
		Update(ctx, req)
	if err != nil {
		if reqErr, ok := gwerror.DbErrorToRequestError(req, err, dao.VideoCollectionDbType); ok {
			return nil, reqErr
		}
		err = gwerror.WrapServiceErrorf(err, req, "更新失败")
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		err = gwerror.WrapServiceErrorf(err, req, "获取插入记录条数出错")
		return nil, err
	}
	message := "更新成功"
	//if rowsAffected == 0 {
	//	return nil, gwerror.NewNotFoundErrorf(req, "不存在要更新的记录")
	//}
	return &p.VideoCollectionUpdateRes{
		Message:      gwwrapper.WrapString(message),
		RowsAffected: gwwrapper.WrapInt64(rowsAffected),
	}, nil
}

// Upsert 根据主键（或唯一索引）是否存在且已在req中赋值，更新或插入对应记录。
// 支持字段类型自动转换，支持对非主键字段赋值/不赋值
// 未赋值或赋值为nil的字段不参与更新/插入（即更新时不会修改原记录的字段值）
func (s *VideoCollectionImpl) Upsert(ctx context.Context, req *p.VideoCollectionUpsertReq) (*p.VideoCollectionUpsertRes, error) {
	var (
		//result       sql.Result
		insertedId   int64
		rowsAffected int64
		//err          error
	)
	//result, err = dao.VideoCollection.Ctx(ctx).FieldsEx(dao.VideoCollection.Columns.CreatedAt).Data(req).Save()
	//if err != nil {
	//	if reqErr, ok := gwerror.DbErrorToRequestError(req, err, dao.VideoCollectionDbType); ok {
	//		return nil, reqErr
	//	}
	//	err = gwerror.WrapServiceErrorf(err, req, "插入/更新失败")
	//	g.Log().Error(ctx, err)
	//	return nil, err
	//}
	//insertedId, _ = result.LastInsertId()
	//rowsAffected, err = result.RowsAffected()
	//if err != nil {
	//	err = gwerror.WrapServiceErrorf(err, req, "获取插入/更新记录条数出错")
	//	g.Log().Error(ctx, err)
	//	return nil, err
	//}
	message := "更新成功"
	if rowsAffected == 0 {
		message = "未插入/更新任何记录" // should not happen
	} else if rowsAffected == 1 {
		message = "插入成功"
	}
	return &p.VideoCollectionUpsertRes{
		Message:      gwwrapper.WrapString(message),
		InsertedId:   gwwrapper.WrapString(insertedId),
		RowsAffected: gwwrapper.WrapInt64(rowsAffected),
	}, nil
}

// Delete 根据req指定的条件删除表中记录
// 未赋值或或赋值为nil的字段不参与条件查询
func (s *VideoCollectionImpl) Delete(ctx context.Context, req *p.VideoCollectionDeleteReq) (*p.VideoCollectionDeleteRes, error) {
	var (
		//result       sql.Result
		rowsAffected int64
		//err          error
	)
	//m := dao.VideoCollection.Ctx(ctx)
	//m, err = gworm.ParseConditions(ctx, req, dao.VideoCollection.ColumnMap, m)
	//if err != nil {
	//	return nil, err
	//}
	//result, err = m.Delete()
	//if err != nil {
	//	err = gwerror.WrapServiceErrorf(err, req, "删除失败")
	//	g.Log().Error(ctx, err)
	//	return nil, err
	//}
	//rowsAffected, err = result.RowsAffected()
	//if err != nil {
	//	err = gwerror.WrapServiceErrorf(err, req, "获取删除记录条数出错")
	//	return nil, err
	//}
	message := fmt.Sprintf("已删除%d条记录", rowsAffected)
	if rowsAffected == 0 {
		message = "不存在要删除的记录"
	}

	return &p.VideoCollectionDeleteRes{
		Message:      gwwrapper.WrapString(message),
		RowsAffected: gwwrapper.WrapInt64(rowsAffected),
	}, nil
}
