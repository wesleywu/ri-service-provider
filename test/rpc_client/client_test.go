package rpc_client

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/castbox/go-guru/pkg/goguru/conf"
	"github.com/castbox/go-guru/pkg/goguru/query"
	goguruTypes "github.com/castbox/go-guru/pkg/goguru/types"
	"github.com/castbox/go-guru/pkg/util/gjson"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "github.com/wesleywu/ri-service-provider/api/videocollection/service/v1"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/enum"
	p "github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

var (
	ctx                   = context.Background()
	flagconf              string
	videoCollectionClient v1.VideoCollectionClient
	helper                *log.Helper
)

func setupSuite() func(tb *testing.M) {
	flag.StringVar(&flagconf, "conf", "configs", "config path, eg: -conf config.yaml")
	flag.Parse()
	c := config.New(
		config.WithSource(
			env.NewSource("GURU_"),
			file.NewSource(flagconf),
		),
	)
	defer func(c config.Config) {
		_ = c.Close()
	}(c)

	if err := c.Load(); err != nil {
		panic(err)
	}
	var cfg conf.Application
	if err := c.Scan(&cfg); err != nil {
		panic(err)
	}
	clients, cleanup, err := wireClient(ctx, cfg.Client, cfg.Log, cfg.Otlp)
	if err != nil {
		panic(err)
	}
	videoCollectionClient = clients.VideoCollection
	helper = log.NewHelper(clients.logger)

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

func TestCount(t *testing.T) {
	//teardownSuite := setupSuite(t)
	//defer teardownSuite(t)

	req := &p.VideoCollectionCountReq{
		Id:          nil,
		Name:        goguruTypes.AnyString("每日推荐视频"),
		ContentType: goguruTypes.AnyInt32Slice([]int32{1, 2}),
		FilterType:  nil,
		Count:       nil,
		IsOnline:    nil,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}
	helper.Info(gjson.MustEncodeString(req))
	res, err := videoCollectionClient.Count(ctx, req)
	if err != nil {
		panic(err)
	}
	helper.Info(gjson.MustEncodeString(res))
	assert.Equal(t, int64(2), res.TotalElements)
}

func TestList(t *testing.T) {
	req := &p.VideoCollectionListReq{
		Id: nil,
		Name: goguruTypes.AnyCondition(query.NewCondition(
			goguruTypes.AnyString("每日"), query.WithOperator(query.OperatorType_Like), query.WithWildcard(query.WildcardType_Contains))),
		ContentType: goguruTypes.AnyUInt32Slice([]uint32{1, 2}),
		FilterType:  nil,
		Count: goguruTypes.AnyCondition(query.NewCondition(
			goguruTypes.AnyUInt32(0), query.WithOperator(query.OperatorType_GT))),
		IsOnline:  nil,
		CreatedAt: nil,
		UpdatedAt: nil,
		PageRequest: &query.PageRequest{
			Number: 1,
			Size:   1,
			Sorts:  nil,
		},
	}
	helper.Info(gjson.MustEncodeString(req))
	res, err := videoCollectionClient.List(ctx, req)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, int64(2), res.PageInfo.TotalElements)
	assert.Equal(t, 1, len(res.Items))
	helper.Info(gjson.MustEncodeString(res))
}

func BenchmarkList(b *testing.B) {

	req := &p.VideoCollectionListReq{
		Id: nil,
		Name: goguruTypes.AnyCondition(query.NewCondition(
			goguruTypes.AnyString("每日"), query.WithOperator(query.OperatorType_Like), query.WithWildcard(query.WildcardType_Contains))),
		ContentType: goguruTypes.AnyUInt32Slice([]uint32{1, 2}),
		FilterType:  nil,
		Count: goguruTypes.AnyCondition(query.NewCondition(
			goguruTypes.AnyUInt32(0), query.WithOperator(query.OperatorType_GT))),
		IsOnline:  nil,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
	helper.Info(gjson.MustEncodeString(req))
	res, err := videoCollectionClient.List(ctx, req)
	if err != nil {
		panic(err)
	}
	assert.Equal(b, int64(2), res.PageInfo.TotalElements)
	helper.Info(gjson.MustEncodeString(res))
}

func TestCreateDeleteOne(t *testing.T) {
	createRes, err := videoCollectionClient.Create(ctx, &p.VideoCollectionCreateReq{
		Name:        goguruTypes.WrapString("特别长的名称特别长的名称特别长的名称特别长的"),
		ContentType: goguruTypes.Wrap(enum.ContentType_PortraitVideo),
		FilterType:  goguruTypes.Wrap(enum.FilterType_Manual),
		Count:       goguruTypes.WrapUInt32(401),
		IsOnline:    goguruTypes.WrapBool(true),
	})
	require.NoError(t, err)
	require.Equal(t, int64(1), createRes.InsertedCount)
	helper.Info(gjson.MustEncodeString(createRes))
	require.NotNil(t, createRes.InsertedID)
	id := *createRes.InsertedID
	helper.Info(id)

	oneRes, err := videoCollectionClient.One(ctx, &p.VideoCollectionOneReq{
		Id: goguruTypes.AnyObjectID(id),
	})
	if err != nil {
		panic(err)
	}
	require.Equal(t, true, oneRes.Found)
	require.NotNil(t, oneRes.Item)
	require.NotNil(t, oneRes.Item.Name)
	require.NotNil(t, oneRes.Item.ContentType)
	require.NotNil(t, oneRes.Item.FilterType)
	require.NotNil(t, oneRes.Item.Count)
	require.NotNil(t, oneRes.Item.IsOnline)
	require.NotNil(t, oneRes.Item.CreatedAt)
	require.NotNil(t, oneRes.Item.UpdatedAt)
	require.Equal(t, "特别长的名称特别长的名称特别长的名称特别长的", *oneRes.Item.Name)
	require.Equal(t, enum.ContentType_PortraitVideo, *oneRes.Item.ContentType)
	require.Equal(t, enum.FilterType_Manual, *oneRes.Item.FilterType)
	require.Equal(t, uint32(401), *oneRes.Item.Count)
	require.Equal(t, true, *oneRes.Item.IsOnline)

	deleteRes, err := videoCollectionClient.Delete(ctx, &p.VideoCollectionDeleteReq{
		Id: id,
	})
	require.NoError(t, err)
	require.Equal(t, int64(1), deleteRes.DeletedCount)
}
