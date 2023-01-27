package rpc_client

import (
	"context"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"fmt"
	"github.com/WesleyWu/ri-service-provider/app/video_collection/model"
	"github.com/WesleyWu/ri-service-provider/gowing/dubbogo"
	"github.com/WesleyWu/ri-service-provider/gowing/gwtypes"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ctx                   = gctx.New()
	videoCollectionClient = new(model.VideoCollectionClientImpl)
)

func init() {
	err := startDubboConsumer(ctx)
	if err != nil {
		panic(err)
	}
}

func TestCount(t *testing.T) {
	req := &model.VideoCollectionCountReq{
		Id:          nil,
		Name:        gwtypes.AnyString("推荐视频集"),
		ContentType: gwtypes.AnyUInt32Slice([]uint32{1, 2}),
		FilterType:  nil,
		Count:       nil,
		IsOnline:    nil,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}
	fmt.Println(gjson.MustEncodeString(req))
	res, err := videoCollectionClient.Count(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Println(gjson.MustEncodeString(res))
	assert.Equal(t, int32(2), res.Total)
}

func TestList(t *testing.T) {
	req := &model.VideoCollectionListReq{
		Id:          nil,
		Name:        gwtypes.AnyCondition(gwtypes.OperatorType_Like, gwtypes.MultiType_Exact, gwtypes.WildcardType_Contains, gwtypes.AnyString("每日")),
		ContentType: gwtypes.AnyUInt32Slice([]uint32{1, 2}),
		FilterType:  nil,
		Count:       gwtypes.AnyCondition(gwtypes.OperatorType_GT, gwtypes.MultiType_Exact, gwtypes.WildcardType_None, gwtypes.AnyUInt32(1)),
		IsOnline:    nil,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}
	fmt.Println(gjson.MustEncodeString(req))
	res, err := videoCollectionClient.List(ctx, req)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, int32(1), res.Total)
	fmt.Println(gjson.MustEncodeString(res))
}

func TestCreate(t *testing.T) {
	res, err := videoCollectionClient.Create(ctx, &model.VideoCollectionCreateReq{
		Id:          gwtypes.WrapString("87104859-5598"),
		Name:        gwtypes.WrapString("特别长的名称特别长的名称特别长的名称特别长的"),
		ContentType: gwtypes.WrapInt32(3),
		FilterType:  gwtypes.WrapInt32(4),
		Count:       gwtypes.WrapUInt32(401),
		IsOnline:    gwtypes.WrapBool(true),
	})
	if err != nil {
		panic(err)
	}
	g.Log().Infof(ctx, gjson.MustEncodeString(res))
}

func startDubboConsumer(ctx context.Context) error {
	registryId := g.Cfg().MustGet(ctx, "rpc.registry.id", "nacosRegistry").String()
	registryProtocol := g.Cfg().MustGet(ctx, "rpc.registry.protocol", "nacos").String()
	registryAddress := g.Cfg().MustGet(ctx, "rpc.registry.address", "127.0.0.1:8848").String()
	registryNamespace := g.Cfg().MustGet(ctx, "rpc.registry.namespace", "public").String()
	development := g.Cfg().MustGet(ctx, "server.debug", "true").Bool()
	loggerStdout := g.Cfg().MustGet(ctx, "logger.stdout", "true").Bool()
	loggerPath := g.Cfg().MustGet(ctx, "rpc.consumer.logDir").String()
	if g.IsEmpty(loggerPath) {
		loggerPath = g.Cfg().MustGet(ctx, "logger.path", "./data/log/gf-app").String()
	}
	loggerFileName := g.Cfg().MustGet(ctx, "rpc.consumer.logFile", "consumer.log").String()
	loggerLevel := g.Cfg().MustGet(ctx, "rpc.consumer.logLevel", "warn").String()

	dubbogo.AddConsumerReference(
		&dubbogo.ConsumerReference{
			ClientImplStructName: "VideoCollectionClientImpl",
			Service:              videoCollectionClient,
			Protocol:             "tri",
		})
	return dubbogo.StartConsumers(ctx,
		&dubbogo.Registry{
			Id:        registryId,
			Type:      registryProtocol,
			Address:   registryAddress,
			Namespace: registryNamespace,
		},
		&dubbogo.ConsumerOption{
			CheckProviderExists: true,
			TimeoutSeconds:      180,
		},
		&dubbogo.LoggerOption{
			Development: development,
			Stdout:      loggerStdout,
			LogDir:      loggerPath,
			LogFileName: loggerFileName,
			Level:       loggerLevel,
		})
}