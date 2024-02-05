package mapping

// VideoCollectionColumnsDef defines and stores column names for table demo_video_collection.
type VideoCollectionColumnsDef struct {
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
	VideoCollectionColumnMap = map[string]string{
		"Id":          "_id",
		"Name":        "name",
		"ContentType": "contentType",
		"FilterType":  "filterType",
		"Count":       "count",
		"IsOnline":    "isOnline",
		"CreatedAt":   "createdAt",
		"UpdatedAt":   "updatedAt",
	}
	VideoCollectionColumns = VideoCollectionColumnsDef{
		Id:          "id",
		Name:        "name",
		ContentType: "contentType",
		FilterType:  "filterType",
		Count:       "count",
		IsOnline:    "is_online",
		CreatedAt:   "createdAt",
		UpdatedAt:   "updatedAt",
	}
)
