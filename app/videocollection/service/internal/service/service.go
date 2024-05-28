package service

import (
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/wesleywu/ri-service-provider/api/videocollection/service/v1"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/internal/biz"
)

// RegisterInfo 用于提供已注册的服务清单，虽然这个意义不大，但如果没有这个 struct 并且被 main.newApp 使用，则无法 wire RegisterToGRPCServer 函数在 App 启动时执行
type RegisterInfo struct {
	registeredServices []string
}

func (r *RegisterInfo) String() string {
	return fmt.Sprintf("Registered HTTP services: %s", strings.Join(r.registeredServices, ", "))
}

type VideoCollection struct {
	v1.UnimplementedVideoCollectionServer
	repo   biz.VideoCollectionRepo
	helper *log.Helper
}

// RegisterToHTTPServer 将所有的服务注册到 http server，实际是注册了 http 路由和 handler
func RegisterToHTTPServer(svr *http.Server, videoCollectionSvc *VideoCollection) (*RegisterInfo, error) {
	var (
		info = &RegisterInfo{
			registeredServices: []string{
				"templateV1",
			},
		}
	)
	v1.RegisterVideoCollectionHTTPServer(svr, videoCollectionSvc)
	return info, nil
}
