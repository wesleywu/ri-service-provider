// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/castbox/go-guru/pkg/goguru/conf"
	"github.com/castbox/go-guru/pkg/infra/logger"
	"github.com/castbox/go-guru/pkg/infra/mongodb"
	"github.com/castbox/go-guru/pkg/infra/otlp"
	"github.com/castbox/go-guru/pkg/infra/redis"
	"github.com/castbox/go-guru/pkg/middleware/servicecache"
	"github.com/castbox/go-guru/pkg/server/grpc"
	"github.com/castbox/go-guru/pkg/server/http"
	"github.com/go-kratos/kratos/v2"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/service"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

// Injectors from wire.go:

func wireApp(contextContext context.Context, server_HTTP *conf.Server_HTTP, server_GRPC *conf.Server_GRPC, server_ServiceCache *conf.Server_ServiceCache, database *conf.Database, confRedis *conf.Redis, log *conf.Log, confOtlp *conf.Otlp) (*kratos.App, func(), error) {
	appMetadata := newAppMetadata()
	configs := logger.NewConfigsByGuru(appMetadata, log)
	logLogger, err := logger.NewLogger(appMetadata, configs)
	if err != nil {
		return nil, nil, err
	}
	httpConfigs, err := http.NewConfigsByGuru(server_HTTP)
	if err != nil {
		return nil, nil, err
	}
	otlpConfigs, err := otlp.NewConfigsByGuru(appMetadata, confOtlp)
	if err != nil {
		return nil, nil, err
	}
	helper := logger.NewLoggerHelper(contextContext, logLogger)
	tracerProvider, err := otlp.NewTracerProvider(contextContext, appMetadata, otlpConfigs, helper)
	if err != nil {
		return nil, nil, err
	}
	encodeResponseFunc := http.DefaultResponseEncoderFunc()
	redisConfigs := redis.NewConfigsByGuru(confRedis)
	universalClient, cleanup, err := redis.NewRedisCache(contextContext, redisConfigs, helper, tracerProvider)
	if err != nil {
		return nil, nil, err
	}
	cacheProvider, err := servicecache.NewCacheProvider(contextContext, server_ServiceCache, logLogger, tracerProvider, universalClient)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	v := useMiddlewares(cacheProvider, logLogger)
	server, err := http.NewHTTPServer(contextContext, httpConfigs, logLogger, tracerProvider, encodeResponseFunc, v...)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	grpcConfigs, err := grpc.NewConfigsByGuru(server_GRPC)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	grpcServer, err := grpc.NewGRPCServer(contextContext, grpcConfigs, logLogger, tracerProvider, v...)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mongodbConfigs := mongodb.NewConfigsByGuru(database, tracerProvider)
	client, cleanup2, err := mongodb.NewMongoClient(contextContext, helper, mongodbConfigs)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	videoCollectionRepo, err := proto.NewVideoCollectionRepo(client, helper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	videoCollection := service.NewVideoCollectionService(videoCollectionRepo, helper)
	httpRegisterInfo, err := service.RegisterToHTTPServer(server, videoCollection)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	grpcRegisterInfo, err := service.RegisterToGRPCServer(grpcServer, videoCollection)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app, err := newApp(contextContext, appMetadata, logLogger, server, grpcServer, httpRegisterInfo, grpcRegisterInfo)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
