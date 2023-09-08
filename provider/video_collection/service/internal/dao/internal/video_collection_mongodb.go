package internal

import (
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/wesleywu/gowing/errors/gwerror"
	"github.com/wesleywu/gowing/util/gworm"
	"github.com/wesleywu/gowing/util/gworm/mongodb"
	p "github.com/wesleywu/ri-service-provider/proto/video_collection"
)

// VideoCollectionDaoMongodb is the manager for logic model data accessing and custom defined data operations functions management.
type VideoCollectionDaoMongodb struct {
	Collection string // Collection is the underlying table name of the DAO.
	Group      string // Group is the database configuration group name of current DAO.
	ColumnMap  map[string]string
	Type       gworm.ModelType
}

var (
	videoCollectionMongoColumnMap = g.MapStrStr{
		"Id":          "_id",
		"Name":        "name",
		"ContentType": "contentType",
		"FilterType":  "filterType",
		"Count":       "count",
		"IsOnline":    "isOnline",
		"CreatedAt":   "createdAt",
		"UpdatedAt":   "updatedAt",
	}
	videoCollectionDaoMongodb = &VideoCollectionDaoMongodb{
		Group:      "default",
		Collection: "video_collection",
		ColumnMap:  videoCollectionMongoColumnMap,
		Type:       gworm.MONGO,
	}
)

// NewVideoCollectionDaoMongodb creates and returns a new DAO object for table data access.
func NewVideoCollectionDaoMongodb() *VideoCollectionDaoMongodb {
	return videoCollectionDaoMongodb
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VideoCollectionDaoMongodb) Ctx(ctx context.Context) *gworm.Model {
	return &gworm.Model{
		Type:       gworm.MONGO,
		MongoModel: mongodb.NewModel(ctx, dao.Collection),
	}
}

func (dao *VideoCollectionDaoMongodb) Create(ctx context.Context, data interface{}) (*gworm.Result, error) {
	result, err := dao.Ctx(ctx).InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (dao *VideoCollectionDaoMongodb) Update(ctx context.Context, keyValue interface{}, data interface{}) (*gworm.Result, error) {
	result, err := dao.Ctx(ctx).
		FieldsEx(videoCollectionColumns.Id, videoCollectionColumns.CreatedAt).
		WherePri(keyValue).
		Update(ctx, data)
	return result, err
}

func (dao *VideoCollectionDaoMongodb) Upsert(ctx context.Context, keyValue interface{}, data interface{}) (*gworm.Result, error) {
	result, err := dao.Ctx(ctx).
		FieldsEx(videoCollectionColumns.Id, videoCollectionColumns.CreatedAt).
		WherePri(keyValue).
		Upsert(ctx, data)
	return result, err
}

func (dao *VideoCollectionDaoMongodb) Delete(ctx context.Context, data interface{}) (*gworm.Result, error) {
	var (
		result *gworm.Result
		err    error
	)
	m := dao.Ctx(ctx).WithAll()
	m, err = gworm.ParseConditions(ctx, data, dao.ColumnMap, m)
	if err != nil {
		return nil, err
	}
	result, err = m.Delete(ctx)
	return result, err
}

func (dao *VideoCollectionDaoMongodb) Count(ctx context.Context, req gworm.FilterRequest) (int64, error) {
	var err error
	m := dao.Ctx(ctx).WithAll()
	m, err = gworm.ApplyFilter(ctx, req, m)
	if err != nil {
		return 0, err
	}
	count, err := m.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (dao *VideoCollectionDaoMongodb) One(ctx context.Context, req gworm.FilterRequest, pageRequest gworm.PageRequest) (item *p.VideoCollectionItem, err error) {
	m := dao.Ctx(ctx).WithAll()
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

func (dao *VideoCollectionDaoMongodb) List(ctx context.Context, req gworm.FilterRequest, pageRequest gworm.PageRequest) (
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
	pageInfo = &gworm.PageInfo{}
	pageInfo.From(pageRequest, uint32(len(list)), uint64(total))
	return list, pageInfo, nil
}
