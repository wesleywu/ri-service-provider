package main

//
//import (
//	"github.com/gogf/gf/contrib/trace/jaeger/v2"
//	"github.com/gogf/gf/v2/frame/g"
//	"github.com/gogf/gf/v2/os/gctx"
//	"github.com/wesleywu/go-lifespan/lifespan"
//	_ "github.com/wesleywu/gowing/boot"
//	_ "github.com/wesleywu/gowing/web/router"
//)
//
//var (
//	ctx               = gctx.New()
//	ServiceName       = "VideoCollectionApi"
//	JaegerUdpEndpoint = "localhost:6831"
//)
//
//func main() {
//	lifespan.OnBootstrap(ctx)
//	tp, err := jaeger.Init(ServiceName, JaegerUdpEndpoint)
//	if err != nil {
//		g.Log().Fatalf(ctx, "%+v", err)
//	}
//	defer tp.Shutdown(ctx)
//	s := g.Server()
//	s.Run()
//	lifespan.OnShutdown(ctx)
//}
