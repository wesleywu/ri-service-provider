package rpc_client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/castbox/go-guru/pkg/goguru/orm"
	"github.com/castbox/go-guru/pkg/goguru/types"
	"github.com/stretchr/testify/assert"
	v1 "github.com/wesleywu/ri-service-provider/api/videocollection/service/v1"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/enum"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
	"go.opentelemetry.io/otel"
)

var (
	ctx    = context.Background()
	client v1.VideoCollectionClient
)

func setupSuite() func(tb *testing.M) {
	clients, cleanup, err := wireClient(ctx)
	if err != nil {
		panic(err)
	}
	client = clients.VideoCollection

	// Return a function to teardown the test
	return func(t *testing.M) {
		cleanup()
	}
}

func TestMain(m *testing.M) {
	// Setup code goes here
	f := setupSuite()
	defer f(m)
	code := m.Run()
	// Teardown code goes here
	os.Exit(code)
}

func TestVideoCollectionRepo_All(t *testing.T) {
	var (
		createReq      *p.VideoCollectionCreateReq
		upsertReq      *p.VideoCollectionUpsertReq
		updateReq      *p.VideoCollectionUpdateReq
		oneReq         *p.VideoCollectionOneReq
		countReq       *p.VideoCollectionCountReq
		getReq         *p.VideoCollectionGetReq
		listReq        *p.VideoCollectionListReq
		deleteReq      *p.VideoCollectionDeleteReq
		deleteMultiReq *p.VideoCollectionDeleteMultiReq
		createRes      *p.VideoCollectionCreateRes
		upsertRes      *p.VideoCollectionUpsertRes
		updateRes      *p.VideoCollectionUpdateRes
		oneRes         *p.VideoCollectionOneRes
		countRes       *p.VideoCollectionCountRes
		getRes         *p.VideoCollectionGetRes
		listRes        *p.VideoCollectionListRes
		deleteRes      *p.VideoCollectionDeleteRes
		deleteMultiRes *p.VideoCollectionDeleteMultiRes
		insertedID1    string
		insertedID2    = "qiihWlTCtVz72T9znB9"
		err            error
	)
	// 创建链路追踪
	tp := otel.GetTracerProvider()
	tracer := tp.Tracer("ri-service-provider-test")
	ctx, span := tracer.Start(ctx, "grpc_client_test.TestVideoCollectionRepo_All")
	defer span.End()
	// test Delete 删除1条可能之前存在的记录
	deleteReq = &p.VideoCollectionDeleteReq{
		Id: insertedID2,
	}
	deleteRes, err = client.Delete(ctx, deleteReq)
	assert.NoError(t, err)

	dateStarted := time.Now()

	// test Create 会插入一条记录
	createReq = &p.VideoCollectionCreateReq{
		Name:        types.Wrap("测试视频集01"),
		ContentType: types.Wrap(enum.ContentType_PortraitVideo),
		FilterType:  types.Wrap(enum.FilterType_Manual),
		Count:       types.WrapInt32(1234),
		IsOnline:    types.Wrap(false),
	}
	createRes, err = client.Create(ctx, createReq)
	assert.NoError(t, err)
	assert.NotNil(t, createRes)
	assert.NotNil(t, createRes.InsertedID)
	assert.Equal(t, int64(1), createRes.InsertedCount)
	insertedID1 = *createRes.InsertedID

	// test Upsert 会插入第二条记录
	upsertReq = &p.VideoCollectionUpsertReq{
		Id:          insertedID2,
		Name:        types.Wrap("测试视频集02"),
		ContentType: types.Wrap(enum.ContentType_LandscapeVideo),
		FilterType:  types.Wrap(enum.FilterType_Ruled),
		Count:       types.WrapInt32(2345),
		IsOnline:    types.Wrap(true),
	}
	upsertRes, err = client.Upsert(ctx, upsertReq)
	assert.NoError(t, err)
	assert.NotNil(t, upsertRes.UpsertedID)
	assert.Equal(t, insertedID2, *upsertRes.UpsertedID)
	assert.Equal(t, int64(1), upsertRes.UpsertedCount)

	// test One 第1次，命中1条记录
	oneReq = &p.VideoCollectionOneReq{
		Name:        types.AnyString("测试视频集01"),
		ContentType: types.AnyStringSlice([]string{"LandscapeVideo", "PortraitVideo"}),
		IsOnline:    types.AnyBoolSlice([]bool{true, false}),
		CreatedAt:   types.AnyCondition(orm.NewCondition(types.AnyTimestamp(dateStarted), orm.WithOperator(orm.OperatorType_GTE))),
	}
	oneRes, err = client.One(ctx, oneReq)
	assert.NoError(t, err)
	assert.Equal(t, true, oneRes.Found)
	assert.Equal(t, int32(1234), *oneRes.Item.Count)

	// test One 第2次，无命中记录
	oneReq = &p.VideoCollectionOneReq{
		Name:        types.AnyString("测试视频集01"),
		ContentType: types.AnyString("LandscapeVideo"),
		IsOnline:    types.AnyBoolSlice([]bool{true, false}),
		CreatedAt:   types.AnyCondition(orm.NewCondition(types.AnyTimestamp(dateStarted), orm.WithOperator(orm.OperatorType_GTE))),
	}
	oneRes, err = client.One(ctx, oneReq)
	assert.NoError(t, err)
	assert.Equal(t, false, oneRes.Found)

	// test One 第3次，命中1条记录
	oneReq = &p.VideoCollectionOneReq{
		Id: types.AnyObjectID(insertedID1),
	}
	oneRes, err = client.One(ctx, oneReq)
	assert.NoError(t, err)
	assert.Equal(t, true, oneRes.Found)
	assert.Equal(t, "测试视频集01", *oneRes.Item.Name)

	// test Count 第1次，共2条满足条件的记录
	countReq = &p.VideoCollectionCountReq{
		ContentType: types.AnyStringSlice([]string{"LandscapeVideo", "PortraitVideo"}),
		IsOnline:    types.AnyBoolSlice([]bool{true, false}),
		CreatedAt:   types.AnyCondition(orm.NewCondition(types.AnyTimestamp(dateStarted), orm.WithOperator(orm.OperatorType_GTE))),
	}
	countRes, err = client.Count(ctx, countReq)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), countRes.TotalElements)

	// test Count 第2次，共1条满足条件的记录
	countReq = &p.VideoCollectionCountReq{
		Name:      types.AnyString("测试视频集01"),
		CreatedAt: types.AnyCondition(orm.NewCondition(types.AnyTimestamp(dateStarted), orm.WithOperator(orm.OperatorType_GTE))),
	}
	countRes, err = client.Count(ctx, countReq)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), countRes.TotalElements)

	// test Count 第3次，命中2条记录
	nextDate := dateStarted.AddDate(0, 0, 1)
	countReq = &p.VideoCollectionCountReq{
		CreatedAt: types.AnyCondition(orm.NewCondition(types.AnyTimestampSlice([]time.Time{dateStarted, nextDate}), orm.WithMulti(orm.MultiType_Between))),
	}
	countRes, err = client.Count(ctx, countReq)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), countRes.TotalElements)

	// test List 第1次，返回第2页，每页1条记录，当页有1条记录，为满足条件的第2条记录，其 Name 为 "TemplateName456"
	listReq = &p.VideoCollectionListReq{
		ContentType: types.AnyStringSlice([]string{"LandscapeVideo", "PortraitVideo"}),
		IsOnline:    types.AnyBoolSlice([]bool{true, false}),
		CreatedAt:   types.AnyCondition(orm.NewCondition(types.AnyTimestamp(dateStarted), orm.WithOperator(orm.OperatorType_GTE))),
		PageRequest: &orm.PageRequest{
			Number: 2,
			Size:   1,
			Sorts: []*orm.SortParam{
				{
					Property:  "name",
					Direction: orm.SortDirection_Asc,
				},
			},
		},
	}
	listRes, err = client.List(ctx, listReq)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), listRes.PageInfo.Number)
	assert.Equal(t, int64(2), listRes.PageInfo.TotalElements)
	assert.Equal(t, int64(2), listRes.PageInfo.TotalPages)
	assert.Equal(t, int64(1), listRes.PageInfo.NumberOfElements)
	assert.Equal(t, false, listRes.PageInfo.First)
	assert.Equal(t, true, listRes.PageInfo.Last)

	// test Get 返回第一条记录
	getReq = &p.VideoCollectionGetReq{
		Id: insertedID1,
	}
	getRes, err = client.Get(ctx, getReq)
	assert.NoError(t, err)
	assert.NotNil(t, getRes.Name)
	assert.Equal(t, "测试视频集01", *getRes.Name)

	// test Update 修改第一条记录
	updateReq = &p.VideoCollectionUpdateReq{
		Id:       insertedID1,
		Name:     types.Wrap("测试视频集03"),
		Count:    types.WrapInt32(3456),
		IsOnline: types.Wrap(false),
	}
	updateRes, err = client.Update(ctx, updateReq)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), updateRes.ModifiedCount)

	// test Get 再次验证第一条记录
	getReq = &p.VideoCollectionGetReq{
		Id: insertedID1,
	}
	getRes, err = client.Get(ctx, getReq)
	assert.NoError(t, err)
	assert.NotNil(t, getRes.Name)
	assert.NotNil(t, getRes.Count)
	assert.NotNil(t, getRes.IsOnline)
	assert.Equal(t, "测试视频集03", *getRes.Name)
	assert.Equal(t, int32(3456), *getRes.Count)
	assert.Equal(t, false, *getRes.IsOnline)

	// test Upsert 修改第一条记录
	upsertReq = &p.VideoCollectionUpsertReq{
		Id:       insertedID1,
		Name:     types.Wrap("测试视频集04"),
		Count:    types.WrapInt32(4567),
		IsOnline: types.Wrap(true),
	}
	upsertRes, err = client.Upsert(ctx, upsertReq)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), updateRes.ModifiedCount)

	// test Get 再次验证第一条记录
	getReq = &p.VideoCollectionGetReq{
		Id: insertedID1,
	}
	getRes, err = client.Get(ctx, getReq)
	assert.NoError(t, err)
	assert.NotNil(t, getRes.Name)
	assert.NotNil(t, getRes.Count)
	assert.NotNil(t, getRes.IsOnline)
	assert.Equal(t, "测试视频集04", *getRes.Name)
	assert.Equal(t, int32(4567), *getRes.Count)
	assert.Equal(t, true, *getRes.IsOnline)

	// test DeleteMulti 删除2条记录
	deleteMultiReq = &p.VideoCollectionDeleteMultiReq{
		Id: types.AnyObjectIDSlice([]string{insertedID1, insertedID2}),
	}
	deleteMultiRes, err = client.DeleteMulti(ctx, deleteMultiReq)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), deleteMultiRes.DeletedCount)

	// test Delete 删除0条记录，因为之前的 deleteMulti 已经删除过了
	deleteReq = &p.VideoCollectionDeleteReq{
		Id: insertedID1,
	}
	deleteRes, err = client.Delete(ctx, deleteReq)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), deleteRes.DeletedCount)
}
