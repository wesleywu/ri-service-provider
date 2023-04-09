package model

import (
	"github.com/WesleyWu/gowing/util/gwwrapper"
	proto "github.com/WesleyWu/ri-service-provider/proto/video_collection"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoCollectionCountReq 查询记录总条数的条件数据结构
type VideoCollectionCountReq struct {
	g.Meta      `json:"-" path:"/count" method:"get"`
	Id          interface{} `p:"id" json:"id"`                                         // 视频集ID，字符串格式
	Name        interface{} `p:"name" wildcard:"none" json:"name,omitempty"`           // 视频集名称
	ContentType interface{} `p:"contentType" json:"contentType,omitempty"`             // 内容类型
	FilterType  interface{} `p:"filterType" json:"filterType,omitempty"`               // 筛选类型
	Count       interface{} `p:"count" json:"count,omitempty"`                         // 集合内视频数量
	IsOnline    interface{} `p:"isOnline" json:"isOnline,omitempty"`                   // 是否上线：0 未上线|1 已上线
	CreatedAt   interface{} `p:"createdAt" multi:"between" json:"createdAt,omitempty"` // 创建时间
	UpdatedAt   interface{} `p:"updatedAt" multi:"between" json:"updatedAt,omitempty"` // 更新时间
}

// VideoCollectionCountRes 查询记录总条数的返回结果
type VideoCollectionCountRes struct {
	Total int `json:"total"`
}

// VideoCollectionOneReq 查询单一记录的条件数据结构
type VideoCollectionOneReq struct {
	g.Meta      `json:"-" path:"/one" method:"get"`
	Id          interface{} `p:"id" json:"id"`                                         // 视频集ID，字符串格式
	Name        interface{} `p:"name" wildcard:"none" json:"name,omitempty"`           // 视频集名称
	ContentType interface{} `p:"contentType" json:"contentType,omitempty"`             // 内容类型
	FilterType  interface{} `p:"filterType" json:"filterType,omitempty"`               // 筛选类型
	Count       interface{} `p:"count" json:"count,omitempty"`                         // 集合内视频数量
	IsOnline    interface{} `p:"isOnline" json:"isOnline,omitempty"`                   // 是否上线：0 未上线|1 已上线
	CreatedAt   interface{} `p:"createdAt" multi:"between" json:"createdAt,omitempty"` // 创建时间
	UpdatedAt   interface{} `p:"updatedAt" multi:"between" json:"updatedAt,omitempty"` // 更新时间
	OrderBy     string      `json:"orderBy,omitempty"`                                 // 排序方式
}

// VideoCollectionOneRes 查询单一记录的返回结果
type VideoCollectionOneRes struct {
	Id          string      `json:"id"`          // 视频集ID，字符串格式
	Name        string      `json:"name"`        // 视频集名称
	ContentType int         `json:"contentType"` // 内容类型
	FilterType  int         `json:"filterType"`  // 筛选类型
	Count       uint32      `json:"count"`       // 集合内视频数量
	IsOnline    bool        `json:"isOnline"`    // 是否上线：0 未上线|1 已上线
	CreatedAt   *gtime.Time `json:"createdAt"`   // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"`   // 更新时间
}

// VideoCollectionListReq 用于列表查询的查询条件数据结构，支持翻页和排序参数，支持查询条件参数类型自动转换
type VideoCollectionListReq struct {
	g.Meta      `json:"-" path:"/list" method:"get"`
	Id          interface{} `p:"id" json:"id"`                                         // 视频集ID，字符串格式
	Name        interface{} `p:"name" wildcard:"none" json:"name,omitempty"`           // 视频集名称
	ContentType interface{} `p:"contentType" json:"contentType,omitempty"`             // 内容类型
	FilterType  interface{} `p:"filterType" json:"filterType,omitempty"`               // 筛选类型
	Count       interface{} `p:"count" json:"count,omitempty"`                         // 集合内视频数量
	IsOnline    interface{} `p:"isOnline" json:"isOnline,omitempty"`                   // 是否上线：0 未上线|1 已上线
	CreatedAt   interface{} `p:"createdAt" multi:"between" json:"createdAt,omitempty"` // 创建时间
	UpdatedAt   interface{} `p:"updatedAt" multi:"between" json:"updatedAt,omitempty"` // 更新时间
	Page        uint32      `d:"1" v:"min:0#分页号码错误" json:"page,omitempty"`             // 当前页码
	PageSize    uint32      `d:"10" v:"max:50#分页数量最大50条" json:"pageSize,omitempty"`    // 每页记录数
	OrderBy     string      `json:"orderBy,omitempty"`                                 // 排序方式
}

// VideoCollectionItem 数据对象
type VideoCollectionItem struct {
	Id          string      `json:"id"`          // 视频集ID，字符串格式
	Name        string      `json:"name"`        // 视频集名称
	ContentType int         `json:"contentType"` // 内容类型
	FilterType  int         `json:"filterType"`  // 筛选类型
	Count       uint32      `json:"count"`       // 集合内视频数量
	IsOnline    bool        `json:"isOnline"`    // 是否上线：0 未上线|1 已上线
	CreatedAt   *gtime.Time `json:"createdAt"`   // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"`   // 更新时间
}

// VideoCollectionListRes 分页返回结果
type VideoCollectionListRes struct {
	Total   int64                  `json:"total"`   // 记录总数
	Current uint32                 `json:"current"` // 当前页码
	Items   []*VideoCollectionItem `json:"items"`   // 当前页记录列表
}

// VideoCollectionCreateReq 插入操作请求参数
type VideoCollectionCreateReq struct {
	g.Meta      `orm:"do:true" json:"-" path:"/" method:"post"`
	Id          interface{} `p:"id" v:"required#视频集ID，字符串格式不能为空" json:"id"` // 视频集ID，字符串格式
	Name        interface{} `p:"name" json:"name"`                          // 视频集名称
	ContentType interface{} `p:"contentType" json:"contentType"`            // 内容类型
	FilterType  interface{} `p:"filterType" json:"filterType"`              // 筛选类型
	Count       interface{} `p:"count" json:"count"`                        // 集合内视频数量
	IsOnline    interface{} `p:"isOnline" json:"isOnline"`                  // 是否上线：0 未上线|1 已上线
	CreatedAt   interface{} `p:"createdAt" json:"createdAt"`                // 创建时间
	UpdatedAt   interface{} `p:"updatedAt" json:"updatedAt"`                // 更新时间
}

// VideoCollectionCreateRes 插入操作返回结果
type VideoCollectionCreateRes struct {
	Message      string `json:"message"`      // 提示信息
	LastInsertId int64  `json:"lastInsertId"` // 上一条INSERT插入的记录主键，当主键为自增长时有效
	RowsAffected int64  `json:"rowsAffected"` // 影响的条数
}

// VideoCollectionUpdateReq 更新操作请求参数
type VideoCollectionUpdateReq struct {
	g.Meta      `orm:"do:true" json:"-" path:"/:id" method:"patch"`
	Id          interface{} `p:"id" v:"required#视频集ID，字符串格式不能为空" json:"id"` // 视频集ID，字符串格式
	Name        interface{} `p:"name" json:"name"`                          // 视频集名称
	ContentType interface{} `p:"contentType" json:"contentType"`            // 内容类型
	FilterType  interface{} `p:"filterType" json:"filterType"`              // 筛选类型
	Count       interface{} `p:"count" json:"count"`                        // 集合内视频数量
	IsOnline    interface{} `p:"isOnline" json:"isOnline"`                  // 是否上线：0 未上线|1 已上线
	CreatedAt   interface{} `p:"createdAt" json:"createdAt"`                // 创建时间
	UpdatedAt   interface{} `p:"updatedAt" json:"updatedAt"`                // 更新时间
}

// VideoCollectionUpdateRes 更新操作返回结果
type VideoCollectionUpdateRes struct {
	Message      string `json:"message"` // 提示信息
	RowsAffected int64  `json:"rowsAffected"`
}

// VideoCollectionUpsertReq 更新插入操作请求参数
type VideoCollectionUpsertReq struct {
	g.Meta      `orm:"do:true" json:"-" path:"/" method:"put"`
	Id          interface{} `p:"id" v:"required#视频集ID，字符串格式不能为空" json:"id"` // 视频集ID，字符串格式
	Name        interface{} `p:"name" json:"name"`                          // 视频集名称
	ContentType interface{} `p:"contentType" json:"contentType"`            // 内容类型
	FilterType  interface{} `p:"filterType" json:"filterType"`              // 筛选类型
	Count       interface{} `p:"count" json:"count"`                        // 集合内视频数量
	IsOnline    interface{} `p:"isOnline" json:"isOnline"`                  // 是否上线：0 未上线|1 已上线
	CreatedAt   interface{} `p:"createdAt" json:"createdAt"`                // 创建时间
	UpdatedAt   interface{} `p:"updatedAt" json:"updatedAt"`                // 更新时间
}

// VideoCollectionUpsertRes 更新插入操作返回结果
type VideoCollectionUpsertRes struct {
	Message      string `json:"message"`      // 提示信息
	LastInsertId int64  `json:"lastInsertId"` // 上一条INSERT插入的记录主键，当主键为自增长时有效
	RowsAffected int64  `json:"rowsAffected"` // 影响的条数
}

// VideoCollectionDeleteReq 删除操作请求参数
type VideoCollectionDeleteReq struct {
	g.Meta      `json:"-" path:"/*id" method:"delete"`
	Id          interface{} `p:"id" json:"id"`                                         // 视频集ID，字符串格式
	Name        interface{} `p:"name" wildcard:"none" json:"name,omitempty"`           // 视频集名称
	ContentType interface{} `p:"contentType" json:"contentType,omitempty"`             // 内容类型
	FilterType  interface{} `p:"filterType" json:"filterType,omitempty"`               // 筛选类型
	Count       interface{} `p:"count" json:"count,omitempty"`                         // 集合内视频数量
	IsOnline    interface{} `p:"isOnline" json:"isOnline,omitempty"`                   // 是否上线：0 未上线|1 已上线
	CreatedAt   interface{} `p:"createdAt" multi:"between" json:"createdAt,omitempty"` // 创建时间
	UpdatedAt   interface{} `p:"updatedAt" multi:"between" json:"updatedAt,omitempty"` // 更新时间
}

// VideoCollectionDeleteRes 删除操作返回结果
type VideoCollectionDeleteRes struct {
	Message      string `json:"message"`      // 提示信息
	RowsAffected int64  `json:"rowsAffected"` // 影响的条数
}

func (x *VideoCollectionCreateReq) MarshalMessage() (*proto.VideoCollectionCreateReq, error) {
	return &proto.VideoCollectionCreateReq{
		Id:          gwwrapper.WrapString(x.Id),
		Name:        gwwrapper.WrapString(x.Name),
		ContentType: gwwrapper.WrapInt32(x.ContentType),
		FilterType:  gwwrapper.WrapInt32(x.FilterType),
		Count:       gwwrapper.WrapUInt32(x.Count),
		IsOnline:    gwwrapper.WrapBool(x.IsOnline),
		CreatedAt:   gwwrapper.WrapTimestamp(x.CreatedAt),
		UpdatedAt:   gwwrapper.WrapTimestamp(x.UpdatedAt),
	}, nil
}

func (x *VideoCollectionUpdateReq) MarshalMessage() (*proto.VideoCollectionUpdateReq, error) {
	return &proto.VideoCollectionUpdateReq{
		Id:          gwwrapper.WrapString(x.Id),
		Name:        gwwrapper.WrapString(x.Name),
		ContentType: gwwrapper.WrapInt32(x.ContentType),
		FilterType:  gwwrapper.WrapInt32(x.FilterType),
		Count:       gwwrapper.WrapUInt32(x.Count),
		IsOnline:    gwwrapper.WrapBool(x.IsOnline),
		CreatedAt:   gwwrapper.WrapTimestamp(x.CreatedAt),
		UpdatedAt:   gwwrapper.WrapTimestamp(x.UpdatedAt),
	}, nil
}

func (x *VideoCollectionUpsertReq) MarshalMessage() (*proto.VideoCollectionUpsertReq, error) {
	return &proto.VideoCollectionUpsertReq{
		Id:          gwwrapper.WrapString(x.Id),
		Name:        gwwrapper.WrapString(x.Name),
		ContentType: gwwrapper.WrapInt32(x.ContentType),
		FilterType:  gwwrapper.WrapInt32(x.FilterType),
		Count:       gwwrapper.WrapUInt32(x.Count),
		IsOnline:    gwwrapper.WrapBool(x.IsOnline),
		CreatedAt:   gwwrapper.WrapTimestamp(x.CreatedAt),
		UpdatedAt:   gwwrapper.WrapTimestamp(x.UpdatedAt),
	}, nil
}

func (x *VideoCollectionItem) UnmarshalMessage(res *proto.VideoCollectionItem) error {
	var (
		id          string
		name        string
		contentType int32
		filterType  int32
		count       uint32
		isOnline    bool
		createdAt   *gtime.Time
		updatedAt   *gtime.Time
	)
	if idPtr := res.Id; idPtr != nil {
		id = *idPtr
	}
	if namePtr := res.Name; namePtr != nil {
		name = *namePtr
	}
	if contentTypePtr := res.ContentType; contentTypePtr != nil {
		contentType = *contentTypePtr
	}
	if filterTypePtr := res.FilterType; filterTypePtr != nil {
		filterType = *filterTypePtr
	}
	if countPtr := res.Count; countPtr != nil {
		count = *countPtr
	}
	if isOnlinePtr := res.IsOnline; isOnlinePtr != nil {
		isOnline = *isOnlinePtr
	}
	if createdAtPtr := res.CreatedAt; createdAtPtr != nil {
		createdAt = gtime.New(createdAtPtr.AsTime())
	}
	if updatedAtPtr := res.UpdatedAt; updatedAtPtr != nil {
		updatedAt = gtime.New(updatedAtPtr.AsTime())
	}
	*x = VideoCollectionItem{
		Id:          id,
		Name:        name,
		ContentType: int(contentType),
		FilterType:  int(filterType),
		Count:       count,
		IsOnline:    isOnline,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
	return nil
}

func (x *VideoCollectionOneRes) UnmarshalMessage(res *proto.VideoCollectionOneRes) error {
	var (
		id          string
		name        string
		contentType int32
		filterType  int32
		count       uint32
		isOnline    bool
		createdAt   *gtime.Time
		updatedAt   *gtime.Time
	)

	if idPtr := res.Id; idPtr != nil {
		id = *idPtr
	}
	if namePtr := res.Name; namePtr != nil {
		name = *namePtr
	}
	if contentTypePtr := res.ContentType; contentTypePtr != nil {
		contentType = *contentTypePtr
	}
	if filterTypePtr := res.FilterType; filterTypePtr != nil {
		filterType = *filterTypePtr
	}
	if countPtr := res.Count; countPtr != nil {
		count = *countPtr
	}
	if isOnlinePtr := res.IsOnline; isOnlinePtr != nil {
		isOnline = *isOnlinePtr
	}
	if createdAtPtr := res.CreatedAt; createdAtPtr != nil {
		createdAt = gtime.New(createdAtPtr.AsTime())
	}
	if updatedAtPtr := res.UpdatedAt; updatedAtPtr != nil {
		updatedAt = gtime.New(updatedAtPtr.AsTime())
	}
	*x = VideoCollectionOneRes{
		Id:          id,
		Name:        name,
		ContentType: int(contentType),
		FilterType:  int(filterType),
		Count:       count,
		IsOnline:    isOnline,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
	return nil
}

func (x *VideoCollectionListRes) UnmarshalMessage(res *proto.VideoCollectionListRes) error {
	items := make([]*VideoCollectionItem, len(res.Items))
	for i, item := range res.Items {
		items[i] = &VideoCollectionItem{}
		_ = items[i].UnmarshalMessage(item)
	}
	*x = VideoCollectionListRes{
		Total:   *res.Total,
		Current: *res.Current,
		Items:   items,
	}
	return nil
}
