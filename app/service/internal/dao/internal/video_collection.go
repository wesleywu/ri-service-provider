package internal

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VideoCollectionDao is the manager for logic model data accessing and custom defined data operations functions management.
type VideoCollectionDao struct {
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
	videoCollectionDao = &VideoCollectionDao{
		Group:     "default",
		Table:     "demo_video_collection",
		Columns:   videoCollectionColumns,
		ColumnMap: videoCollectionColumnMap,
	}
)

// NewVideoCollectionDao creates and returns a new DAO object for table data access.
func NewVideoCollectionDao() *VideoCollectionDao {
	return videoCollectionDao
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VideoCollectionDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VideoCollectionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(videoCollectionDao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VideoCollectionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
