package rpc_client

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/castbox/go-guru/pkg/guru/service/conf"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/stretchr/testify/assert"
	"github.com/wesleywu/gowing/protobuf/gwtypes"
	p "github.com/wesleywu/ri-service-provider/api/video_collection/v1"
	"github.com/wesleywu/ri-service-provider/gwwrapper"
)

var (
	ctx                   = context.Background()
	flagconf              string
	videoCollectionClient p.VideoCollectionClient
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
		Name:        gwwrapper.AnyString("每日推荐视频"),
		ContentType: gwwrapper.AnyInt32Slice([]int32{1, 2}),
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
	assert.Equal(t, int64(2), *res.Total)
}

func TestList(t *testing.T) {
	req := &p.VideoCollectionListReq{
		Id:          nil,
		Name:        gwwrapper.AnyCondition(gwtypes.OperatorType_Like, gwtypes.MultiType_Exact, gwtypes.WildcardType_Contains, gwwrapper.AnyString("每日")),
		ContentType: gwwrapper.AnyUInt32Slice([]uint32{1, 2}),
		FilterType:  nil,
		Count:       gwwrapper.AnyCondition(gwtypes.OperatorType_GT, gwtypes.MultiType_Exact, gwtypes.WildcardType_None, gwwrapper.AnyUInt32(0)),
		IsOnline:    nil,
		CreatedAt:   nil,
		UpdatedAt:   nil,
		Page:        1,
		PageSize:    1,
	}
	helper.Info(gjson.MustEncodeString(req))
	res, err := videoCollectionClient.List(ctx, req)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, int64(2), res.Total)
	assert.Equal(t, 1, len(res.Items))
	helper.Info(gjson.MustEncodeString(res))
}

func BenchmarkList(b *testing.B) {

	req := &p.VideoCollectionListReq{
		Id:          nil,
		Name:        gwwrapper.AnyCondition(gwtypes.OperatorType_Like, gwtypes.MultiType_Exact, gwtypes.WildcardType_Contains, gwwrapper.AnyString("每日")),
		ContentType: gwwrapper.AnyUInt32Slice([]uint32{1, 2}),
		FilterType:  nil,
		Count:       gwwrapper.AnyCondition(gwtypes.OperatorType_GT, gwtypes.MultiType_Exact, gwtypes.WildcardType_None, gwwrapper.AnyUInt32(0)),
		IsOnline:    nil,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}
	helper.Info(gjson.MustEncodeString(req))
	res, err := videoCollectionClient.List(ctx, req)
	if err != nil {
		panic(err)
	}
	assert.Equal(b, int64(2), res.Total)
	helper.Info(gjson.MustEncodeString(res))
}

func TestCreateDeleteOne(t *testing.T) {
	id := "87104859-5598"
	_, err := videoCollectionClient.Delete(ctx, &p.VideoCollectionDeleteReq{
		Id: id,
	})
	if err != nil {
		panic(err)
	}

	createRes, err := videoCollectionClient.Create(ctx, &p.VideoCollectionCreateReq{
		Id:          gwwrapper.WrapString(id),
		Name:        gwwrapper.WrapString("特别长的名称特别长的名称特别长的名称特别长的"),
		ContentType: gwwrapper.WrapInt32(3),
		FilterType:  gwwrapper.WrapInt32(4),
		Count:       gwwrapper.WrapUInt32(401),
		IsOnline:    gwwrapper.WrapBool(true),
	})
	if err != nil {
		panic(err)
	}
	helper.Info(gjson.MustEncodeString(createRes))

	oneRes, err := videoCollectionClient.One(ctx, &p.VideoCollectionOneReq{
		Id: gwwrapper.AnyString(id),
	})
	if err != nil {
		panic(err)
	}
	assert.Equal(t, oneRes.Found, true)

	_, err = videoCollectionClient.Delete(ctx, &p.VideoCollectionDeleteReq{
		Id: id,
	})
	if err != nil {
		panic(err)
	}
}
