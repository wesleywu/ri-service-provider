package main

import (
	"context"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/WesleyWu/ri-service-provider/app/service"
	_ "github.com/WesleyWu/ri-service-provider/boot"
	"github.com/WesleyWu/ri-service-provider/gowing/dubbogo"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	command := gcmd.Command{
		Name:  "VideoCollection Service Provider",
		Usage: "main --port=${PORT}",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			port := parser.GetOpt("port").Int()
			if port <= 0 {
				port = g.Cfg().MustGet(ctx, "rpc.provider.port").Int()
			}
			if port < 100 {
				return gerror.New("port must greater than 100")
			}
			registryId := g.Cfg().MustGet(ctx, "rpc.registry.id", "nacosRegistry").String()
			registryProtocol := g.Cfg().MustGet(ctx, "rpc.registry.protocol", "nacos").String()
			registryAddress := g.Cfg().MustGet(ctx, "rpc.registry.address", "127.0.0.1:8848").String()
			registryNamespace := g.Cfg().MustGet(ctx, "rpc.registry.namespace", "public").String()
			development := g.Cfg().MustGet(ctx, "server.debug", "true").Bool()
			loggerStdout := g.Cfg().MustGet(ctx, "logger.stdout", "true").Bool()
			loggerPath := g.Cfg().MustGet(ctx, "rpc.provider.logDir").String()
			if g.IsEmpty(loggerPath) {
				loggerPath = g.Cfg().MustGet(ctx, "logger.path", "./data/log/gf-app").String()
			}
			loggerFileName := g.Cfg().MustGet(ctx, "rpc.provider.logFile", "provider.log").String()
			loggerLevel := g.Cfg().MustGet(ctx, "rpc.provider.logLevel", "info").String()
			return dubbogo.StartProvider(ctx, &dubbogo.Registry{
				Id:        registryId,
				Type:      registryProtocol,
				Address:   registryAddress,
				Namespace: registryNamespace,
			}, &dubbogo.ProviderInfo{
				Protocol: "tri",
				Port:     port,
				Services: []dubbogo.ServiceInfo{
					{
						ServerImplStructName: "VideoCollectionImpl",
						Service:              service.VideoCollection,
					},
					{
						ServerImplStructName: "VideoCollectionRepoImpl",
						Service:              service.VideoCollectionRepo,
					},
				},
			}, &dubbogo.LoggerOption{
				Development: development,
				Stdout:      loggerStdout,
				LogDir:      loggerPath,
				LogFileName: loggerFileName,
				Level:       loggerLevel,
			})
		},
	}
	command.Run(gctx.New())
}
