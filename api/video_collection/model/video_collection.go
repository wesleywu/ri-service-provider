package model

import (
	"context"
	"github.com/WesleyWu/gowing/util/gwwrapper"
	proto "github.com/WesleyWu/ri-service-provider/proto/video_collection"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// VideoCollection 数据对象
type VideoCollection struct {
	Id          string      `json:"id"`          // 视频集ID，字符串格式
	Name        string      `json:"name"`        // 视频集名称
	ContentType int         `json:"contentType"` // 内容类型
	FilterType  int         `json:"filterType"`  // 筛选类型
	Count       uint32      `json:"count"`       // 集合内视频数量
	IsOnline    bool        `json:"isOnline"`    // 是否上线：0 未上线|1 已上线
	CreatedAt   *gtime.Time `json:"createdAt"`   // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"`   // 更新时间
}

// VideoCollectionInput 用于Insert、Update、Upsert的输入数据对象结构
type VideoCollectionInput struct {
	Id          interface{} `p:"id" v:"required#视频集ID，字符串格式不能为空" json:"id"` // 视频集ID，字符串格式
	Name        interface{} `p:"name" json:"name"`                          // 视频集名称
	ContentType interface{} `p:"contentType" json:"contentType"`            // 内容类型
	FilterType  interface{} `p:"filterType" json:"filterType"`              // 筛选类型
	Count       interface{} `p:"count" json:"count"`                        // 集合内视频数量
	IsOnline    interface{} `p:"isOnline" json:"isOnline"`                  // 是否上线：0 未上线|1 已上线
	CreatedAt   interface{} `p:"createdAt" json:"createdAt"`                // 创建时间
	UpdatedAt   interface{} `p:"updatedAt" json:"updatedAt"`                // 更新时间
}

// VideoCollectionQuery 用于 Query By Example 模式的查询条件数据结构
type VideoCollectionQuery struct {
	Id          interface{} `p:"id" json:"id"`                                         // 视频集ID，字符串格式
	Name        interface{} `p:"name" wildcard:"none" json:"name,omitempty"`           // 视频集名称
	ContentType interface{} `p:"contentType" json:"contentType,omitempty"`             // 内容类型
	FilterType  interface{} `p:"filterType" json:"filterType,omitempty"`               // 筛选类型
	Count       interface{} `p:"count" json:"count,omitempty"`                         // 集合内视频数量
	IsOnline    interface{} `p:"isOnline" json:"isOnline,omitempty"`                   // 是否上线：0 未上线|1 已上线
	CreatedAt   interface{} `p:"createdAt" multi:"between" json:"createdAt,omitempty"` // 创建时间
	UpdatedAt   interface{} `p:"updatedAt" multi:"between" json:"updatedAt,omitempty"` // 更新时间
}

// VideoCollectionCountReq 查询记录总条数的条件数据结构
type VideoCollectionCountReq struct {
	g.Meta `json:"-" path:"/count" method:"get"`
	VideoCollectionQuery
}

// VideoCollectionCountRes 查询记录总条数的返回结果
type VideoCollectionCountRes struct {
	Total int `json:"total"`
}

// VideoCollectionOneReq 查询单一记录的条件数据结构
type VideoCollectionOneReq struct {
	g.Meta `json:"-" path:"/one" method:"get"`
	VideoCollectionQuery
	OrderBy string `json:"orderBy,omitempty"` // 排序方式
}

// VideoCollectionOneRes 查询单一记录的返回结果
type VideoCollectionOneRes struct {
	VideoCollection
}

// VideoCollectionListReq 用于列表查询的查询条件数据结构，支持翻页和排序参数，支持查询条件参数类型自动转换
type VideoCollectionListReq struct {
	g.Meta `json:"-" path:"/list" method:"get"`
	VideoCollectionQuery
	Page     uint32 `d:"1" v:"min:0#分页号码错误" json:"page,omitempty"`          // 当前页码
	PageSize uint32 `d:"10" v:"max:50#分页数量最大50条" json:"pageSize,omitempty"` // 每页记录数
	OrderBy  string `json:"orderBy,omitempty"`                              // 排序方式
}

// VideoCollectionListRes 分页返回结果
type VideoCollectionListRes struct {
	Total   uint64             `json:"total"`   // 记录总数
	Current uint32             `json:"current"` // 当前页码
	Items   []*VideoCollection `json:"items"`   // 当前页记录列表
}

// VideoCollectionCreateReq 插入操作请求参数
type VideoCollectionCreateReq struct {
	g.Meta `orm:"do:true" json:"-" path:"/" method:"post"`
	VideoCollectionInput
}

// VideoCollectionCreateRes 插入操作返回结果
type VideoCollectionCreateRes struct {
	Message      string `json:"message"`      // 提示信息
	LastInsertId int64  `json:"lastInsertId"` // 上一条INSERT插入的记录主键，当主键为自增长时有效
	RowsAffected int64  `json:"rowsAffected"` // 影响的条数
}

// VideoCollectionUpdateReq 更新操作请求参数
type VideoCollectionUpdateReq struct {
	g.Meta `orm:"do:true" json:"-" path:"/:id" method:"patch"`
	VideoCollectionInput
}

// VideoCollectionUpdateRes 更新操作返回结果
type VideoCollectionUpdateRes struct {
	Message      string `json:"message"` // 提示信息
	RowsAffected int64  `json:"rowsAffected"`
}

// VideoCollectionUpsertReq 更新插入操作请求参数
type VideoCollectionUpsertReq struct {
	g.Meta `orm:"do:true" json:"-" path:"/" method:"put"`
	VideoCollectionInput
}

// VideoCollectionUpsertRes 更新插入操作返回结果
type VideoCollectionUpsertRes struct {
	Message      string `json:"message"`      // 提示信息
	LastInsertId int64  `json:"lastInsertId"` // 上一条INSERT插入的记录主键，当主键为自增长时有效
	RowsAffected int64  `json:"rowsAffected"` // 影响的条数
}

// VideoCollectionDeleteReq 删除操作请求参数
type VideoCollectionDeleteReq struct {
	g.Meta `json:"-" path:"/*id" method:"delete"`
	VideoCollectionQuery
}

// VideoCollectionDeleteRes 删除操作返回结果
type VideoCollectionDeleteRes struct {
	Message      string `json:"message"`      // 提示信息
	RowsAffected int64  `json:"rowsAffected"` // 影响的条数
}

func (s *VideoCollectionCountReq) ToMessage(ctx context.Context) *proto.VideoCollectionCountReq {
	// todo 实现
	return nil
}

func (s *VideoCollectionCountRes) FromMessage(ctx context.Context, res *proto.VideoCollectionCountRes) {
	// todo 实现
	return
}

func (s *VideoCollectionOneReq) ToMessage(ctx context.Context) *proto.VideoCollectionOneReq {
	// todo 支持 Condition Any 类型
	req := &proto.VideoCollectionOneReq{
		Meta:        0,
		Id:          nil,
		Name:        nil,
		ContentType: nil,
		FilterType:  nil,
		Count:       nil,
		IsOnline:    nil,
		CreatedAt:   nil,
		UpdatedAt:   nil,
		OrderBy:     "",
	}
	if !g.IsEmpty(s.Id) {
		req.Id = gwwrapper.AnyString(gconv.String(s.Id))
	}
	if !g.IsEmpty(s.Name) {
		req.Name = gwwrapper.AnyString(gconv.String(s.Name))
	}
	if !g.IsEmpty(s.ContentType) {
		req.ContentType = gwwrapper.AnyInt32(gconv.Int32(s.ContentType))
	}
	if !g.IsEmpty(s.FilterType) {
		req.FilterType = gwwrapper.AnyInt32(gconv.Int32(s.FilterType))
	}
	if !g.IsEmpty(s.Count) {
		req.Count = gwwrapper.AnyUInt32(gconv.Uint32(s.Count))
	}
	if !g.IsEmpty(s.IsOnline) {
		req.IsOnline = gwwrapper.AnyBool(gconv.Bool(s.IsOnline))
	}
	if !g.IsEmpty(s.CreatedAt) {
		req.CreatedAt = gwwrapper.AnyString(gconv.String(s.CreatedAt))
	}
	if !g.IsEmpty(s.Id) {
		req.UpdatedAt = gwwrapper.AnyString(gconv.String(s.UpdatedAt))
	}
	return req
}

func (s *VideoCollectionOneRes) FromMessage(ctx context.Context, res *proto.VideoCollectionOneRes) {
	// todo 支持日期类型
	err := gconv.Struct(res, &s)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
	}
	return
}

func (s *VideoCollectionListReq) ToMessage(ctx context.Context) *proto.VideoCollectionListReq {
	// todo 支持 Condition Any 类型
	res := &proto.VideoCollectionListReq{
		Meta:        0,
		Id:          nil,
		Name:        nil,
		ContentType: nil,
		FilterType:  nil,
		Count:       nil,
		IsOnline:    nil,
		CreatedAt:   nil,
		UpdatedAt:   nil,
		Page:        1,
		PageSize:    10,
		OrderBy:     "",
	}
	if !g.IsEmpty(s.Id) {
		res.Id = gwwrapper.AnyString(gconv.String(s.Id))
	}
	if !g.IsEmpty(s.Name) {
		res.Name = gwwrapper.AnyString(gconv.String(s.Name))
	}
	if !g.IsEmpty(s.ContentType) {
		res.ContentType = gwwrapper.AnyInt32(gconv.Int32(s.ContentType))
	}
	if !g.IsEmpty(s.FilterType) {
		res.FilterType = gwwrapper.AnyInt32(gconv.Int32(s.FilterType))
	}
	if !g.IsEmpty(s.Count) {
		res.Count = gwwrapper.AnyUInt32(gconv.Uint32(s.Count))
	}
	if !g.IsEmpty(s.IsOnline) {
		res.IsOnline = gwwrapper.AnyBool(gconv.Bool(s.IsOnline))
	}
	if !g.IsEmpty(s.CreatedAt) {
		res.CreatedAt = gwwrapper.AnyString(gconv.String(s.CreatedAt))
	}
	if !g.IsEmpty(s.Id) {
		res.UpdatedAt = gwwrapper.AnyString(gconv.String(s.UpdatedAt))
	}
	return res
}

func (s *VideoCollectionListRes) FromMessage(ctx context.Context, res *proto.VideoCollectionListRes) {
	// todo 支持日期类型
	s.Current = *res.Current
	s.Total = *res.Total
	if res.Items == nil {
		s.Items = []*VideoCollection{}
	} else {
		s.Items = make([]*VideoCollection, len(res.Items))
		for i, item := range res.Items {
			err := gconv.Struct(item, &s.Items[i])
			if err != nil {
				g.Log().Errorf(ctx, "%+v", err)
			}
		}
	}
	return
}

func (s *VideoCollectionCreateReq) ToMessage(ctx context.Context) *proto.VideoCollectionCreateReq {
	// todo 实现
	return nil
}

func (s *VideoCollectionCreateRes) FromMessage(ctx context.Context, res *proto.VideoCollectionCreateRes) {
	// todo 实现
	return
}

func (s *VideoCollectionUpdateReq) ToMessage(ctx context.Context) *proto.VideoCollectionUpdateReq {
	// todo 实现
	return nil
}

func (s *VideoCollectionUpdateRes) FromMessage(ctx context.Context, res *proto.VideoCollectionUpdateRes) {
	// todo 实现
	return
}

func (s *VideoCollectionUpsertReq) ToMessage(ctx context.Context) *proto.VideoCollectionUpsertReq {
	// todo 实现
	return nil
}

func (s *VideoCollectionUpsertRes) FromMessage(ctx context.Context, res *proto.VideoCollectionUpsertRes) {
	// todo 实现
	return
}

func (s *VideoCollectionDeleteReq) ToMessage(ctx context.Context) *proto.VideoCollectionDeleteReq {
	// todo 实现
	return nil
}

func (s *VideoCollectionDeleteRes) FromMessage(ctx context.Context, res *proto.VideoCollectionDeleteRes) {
	// todo 实现
	return
}

// IsEmpty 判断删除请求参数是否为空
func (s *VideoCollectionDeleteReq) IsEmpty() bool {
	return g.IsEmpty(s.Id) &&
		s.Name == nil &&
		s.ContentType == nil &&
		s.FilterType == nil &&
		s.Count == nil &&
		s.IsOnline == nil &&
		s.CreatedAt == nil &&
		s.UpdatedAt == nil
}
