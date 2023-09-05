package internal

import (
	"context"
	"database/sql"

	"github.com/WesleyWu/gowing/errors/gwerror"
	"github.com/WesleyWu/gowing/util/gworm"
	p "github.com/WesleyWu/ri-service-provider/proto/video_collection"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// VideoCollectionDaoGdb is the manager for logic model data accessing and custom defined data operations functions management.
type VideoCollectionDaoGdb struct {
	Table     string                 // Table is the underlying table name of the DAO.
	Group     string                 // Group is the database configuration group name of current DAO.
	Columns   VideoCollectionColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
	ColumnMap map[string]string
}

// VideoCollectionColumns defines and stores column names for table demo_video_collection.
type VideoCollectionColumns struct {
	Id          string // 视频集ID，字符串格式
	Name        string // 视频集名称
	ContentType string // 内容类型
	FilterType  string // 筛选类型
	Count       string // 集合内视频数量
	IsOnline    string // 是否上线：0 未上线|1 已上线
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

var (
	videoCollectionColumns = VideoCollectionColumns{
		Id:          "id",
		Name:        "name",
		ContentType: "content_type",
		FilterType:  "filter_type",
		Count:       "count",
		IsOnline:    "is_online",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
	}
	videoCollectionColumnMap = g.MapStrStr{
		"Id":          "id",
		"Name":        "name",
		"ContentType": "content_type",
		"FilterType":  "filter_type",
		"Count":       "count",
		"IsOnline":    "is_online",
		"CreatedAt":   "created_at",
		"UpdatedAt":   "updated_at",
	}
	videoCollectionDaoGdb = &VideoCollectionDaoGdb{
		Group:     "default",
		Table:     "demo_video_collection",
		Columns:   videoCollectionColumns,
		ColumnMap: videoCollectionColumnMap,
	}
)

// NewVideoCollectionDaoGdb creates and returns a new DAO object for table data access.
func NewVideoCollectionDaoGdb() *VideoCollectionDaoGdb {
	return videoCollectionDaoGdb
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VideoCollectionDaoGdb) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VideoCollectionDaoGdb) Ctx(ctx context.Context) *gworm.Model {
	return &gworm.Model{
		Type:    gworm.GF_ORM,
		GfModel: dao.DB().Model(videoCollectionDaoGdb.Table).Safe().Ctx(ctx),
	}
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VideoCollectionDaoGdb) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).GfModel.Transaction(ctx, f)
}

func (dao *VideoCollectionDaoGdb) Create(ctx context.Context, data interface{}) (*gworm.Result, error) {
	result, err := dao.Ctx(ctx).GfModel.Insert(data)
	if err != nil {
		if reqErr, ok := gwerror.DbErrorToRequestError(data, err, dao.DB().GetConfig().Type); ok {
			return nil, reqErr
		}
		return nil, err
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &gworm.Result{
		Type:           gworm.GF_ORM,
		SqlResult:      result,
		LastInsertedId: gconv.String(insertedId),
		RowsAffected:   rowsAffected,
	}, err
}

func (dao *VideoCollectionDaoGdb) Update(ctx context.Context, keyValue interface{}, data interface{}) (*gworm.Result, error) {
	result, err := dao.Ctx(ctx).GfModel.
		FieldsEx(videoCollectionColumns.Id, videoCollectionColumns.CreatedAt).
		WherePri(keyValue).
		Update(data)
	return &gworm.Result{
		Type:      gworm.GF_ORM,
		SqlResult: result,
	}, err
}

func (dao *VideoCollectionDaoGdb) Upsert(ctx context.Context, keyValue interface{}, data interface{}) (*gworm.Result, error) {
	result, err := dao.Ctx(ctx).GfModel.
		FieldsEx(videoCollectionColumns.Id, videoCollectionColumns.CreatedAt).
		Data(data).
		Save()
	return &gworm.Result{
		Type:      gworm.GF_ORM,
		SqlResult: result,
	}, err
}

func (dao *VideoCollectionDaoGdb) Delete(ctx context.Context, data interface{}) (*gworm.Result, error) {
	var (
		result sql.Result
		err    error
	)
	m := dao.Ctx(ctx).WithAll()
	m, err = gworm.ParseConditions(ctx, data, videoCollectionColumnMap, m)
	if err != nil {
		return nil, err
	}
	result, err = m.GfModel.Delete()
	return &gworm.Result{
		Type:      gworm.GF_ORM,
		SqlResult: result,
	}, err
}

func (dao *VideoCollectionDaoGdb) Count(ctx context.Context, data *p.VideoCollectionCountReq) (int64, error) {
	var err error
	m := dao.Ctx(ctx).WithAll()
	m, err = gworm.ParseConditions(ctx, data, videoCollectionColumnMap, m)
	if err != nil {
		return 0, err
	}
	count, err := m.GfModel.Count()
	if err != nil {
		return 0, err
	}
	return int64(count), nil
}

func (dao *VideoCollectionDaoGdb) One(ctx context.Context, req gworm.FilterRequest, pageRequest gworm.PageRequest) (item *p.VideoCollectionItem, err error) {
	m := dao.Ctx(ctx).WithAll()
	// todo replace parseCondition
	m, err = gworm.ApplyFilter(ctx, req, m)
	if err != nil {
		return nil, err
	}
	item = (*p.VideoCollectionItem)(nil)
	err = m.Fields(p.VideoCollectionItem{}).
		Order(pageRequest.OrderString()).
		Limit(1).
		Scan(ctx, &item)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, gwerror.NewNotFoundErrorf(req, "找不到要获取的数据")
		}
		return nil, err
	}
	if item == nil {
		return nil, gwerror.NewNotFoundErrorf(req, "找不到要获取的数据")
	}
	return item, nil
}

func (dao *VideoCollectionDaoGdb) List(ctx context.Context, req gworm.FilterRequest, pageRequest gworm.PageRequest) (
	list []*p.VideoCollectionItem, pageInfo *gworm.PageInfo, err error) {
	var (
		total int64
	)
	m := dao.Ctx(ctx).WithAll()
	m, err = gworm.ApplyFilter(ctx, req, m)
	if err != nil {
		return nil, nil, err
	}
	total, err = m.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list = []*p.VideoCollectionItem{}
	err = m.Fields(p.VideoCollectionItem{}).
		Page(int(pageRequest.Number), int(pageRequest.Size)).
		Order(pageRequest.OrderString()).
		Scan(ctx, &list)
	if err != nil {
		return nil, nil, err
	}
	pageInfo.From(pageRequest, uint32(len(list)), uint64(total))
	return list, pageInfo, nil
}
