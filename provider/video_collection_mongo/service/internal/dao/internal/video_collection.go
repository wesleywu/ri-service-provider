package internal

import (
	"context"
	"github.com/WesleyWu/gowing/util/gworm"
	"github.com/WesleyWu/gowing/util/gworm/mongodb"
	"github.com/WesleyWu/gowing/util/gworm/mongodb/codecs"
	"github.com/gogf/gf/v2/frame/g"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VideoCollectionDao) Ctx(ctx context.Context) *gworm.Model {
	// Register custom codecs for protobuf Timestamp and wrapper types
	reg := codecs.Register(bson.NewRegistryBuilder()).Build()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017").SetRegistry(reg))
	if err != nil {
		panic(err)
	}
	collection := client.Database("gowing").Collection("video_collection")
	return &gworm.Model{
		Type:       gworm.MONGO,
		MongoModel: mongodb.NewModel(collection),
	}
}
