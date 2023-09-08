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
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/wesleywu/go-lifespan/lifespan"
	_ "github.com/wesleywu/gowing/boot"
	_ "github.com/wesleywu/gowing/web/router"
)

var (
	ctx               = gctx.New()
	ServiceName       = "VideoCollectionApi"
	JaegerUdpEndpoint = "localhost:6831"
)

func main() {
	lifespan.OnBootstrap(ctx)
	tp, err := jaeger.Init(ServiceName, JaegerUdpEndpoint)
	if err != nil {
		g.Log().Fatalf(ctx, "%+v", err)
	}
	defer tp.Shutdown(ctx)
	s := g.Server()
	s.Run()
	lifespan.OnShutdown(ctx)
}
