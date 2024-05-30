// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/castbox/go-guru/pkg/goguru/conf"
	"github.com/castbox/go-guru/pkg/middleware/servicecache"
	"github.com/castbox/go-guru/pkg/server"
	"github.com/castbox/go-guru/pkg/util/logger"
	"github.com/castbox/go-guru/pkg/util/mongodb"
	"github.com/castbox/go-guru/pkg/util/otlp"
	"github.com/castbox/go-guru/pkg/util/redis"
	"github.com/go-kratos/kratos/v2"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/data"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/service"
)

// Injectors from wire.go:

func wireApp(contextContext context.Context, confServer *conf.Server, server_ServiceCache *conf.Server_ServiceCache, database *conf.Database, confRedis *conf.Redis, log *conf.Log, confOtlp *conf.Otlp) (*kratos.App, func(), error) {
	appMetadata := newAppMetadata()
	logLogger, err := logger.NewLogger(appMetadata, log)
	if err != nil {
		return nil, nil, err
	}
	tracerProvider, err := otlp.NewTracerProvider(contextContext, appMetadata, confOtlp, logLogger)
	if err != nil {
		return nil, nil, err
	}
	encodeResponseFunc := server.DefaultResponseEncoderFunc()
	universalClient, cleanup, err := redis.NewRedisCache(contextContext, confRedis, logLogger, tracerProvider)
	if err != nil {
		return nil, nil, err
	}
	cacheProvider, err := servicecache.NewCacheProvider(contextContext, server_ServiceCache, logLogger, tracerProvider, universalClient)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	v := useMiddlewares(cacheProvider, logLogger)
	httpServer, err := server.NewHTTPServer(contextContext, confServer, logLogger, tracerProvider, encodeResponseFunc, v...)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mongoDatabase, cleanup2, err := mongodb.NewMongoDatabase(contextContext, database, logLogger, tracerProvider)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	helper := logger.NewLoggerHelper(contextContext, logLogger)
	videoCollectionRepo := data.NewVideoCollectionRepo(mongoDatabase, helper)
	videoCollection := service.NewVideoCollectionService(videoCollectionRepo, helper)
	registerInfo, err := service.RegisterToHTTPServer(httpServer, videoCollection)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app, err := newApp(contextContext, appMetadata, logLogger, httpServer, registerInfo)
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