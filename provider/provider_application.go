/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "github.com/WesleyWu/gowing/boot"
	"github.com/WesleyWu/gowing/rpc/dubbogo"
	"github.com/WesleyWu/ri-service-provider/provider/video_collection/service"
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	ServiceName       = "VideoCollection"
	JaegerUdpEndpoint = "localhost:6831"
)

func main() {
	command := gcmd.Command{
		Name:  "VideoCollection Service Provider",
		Usage: "main --port=${PORT}",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			tp, err := jaeger.Init(ServiceName, JaegerUdpEndpoint)
			if err != nil {
				g.Log().Fatal(ctx, err)
			}
			defer tp.Shutdown(ctx)
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
				ApplicationName: "repo_service",
				Protocol:        "tri",
				Port:            port,
				Services: []dubbogo.ServiceInfo{
					{
						ServerImplStructName: "VideoCollectionImpl",
						Service:              service.VideoCollection,
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
	ctx := gctx.New()
	command.Run(ctx)
}
