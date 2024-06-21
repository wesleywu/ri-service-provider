package service

import (
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/wesleywu/ri-service-provider/api/episode/service/v1"
)

// HttpRegisterInfo 用于提供已注册的服务清单，虽然这个意义不大，但如果没有这个 struct 并且被 main.newApp 使用，则无法 wire RegisterToHTTPServer 函数在 App 启动时执行
type HttpRegisterInfo struct {
	registeredServices []string
}

func (r *HttpRegisterInfo) String() string {
	return fmt.Sprintf("Registered HTTP services: %s", strings.Join(r.registeredServices, ", "))
}

// GrpcRegisterInfo 用于提供已注册的服务清单，虽然这个意义不大，但如果没有这个 struct 并且被 main.newApp 使用，则无法 wire RegisterToGRPCServer 函数在 App 启动时执行
type GrpcRegisterInfo struct {
	registeredServices []string
}

func (r *GrpcRegisterInfo) String() string {
	return fmt.Sprintf("Registered HTTP services: %s", strings.Join(r.registeredServices, ", "))
}

// RegisterToHTTPServer 将所有的服务注册到 http server，实际是注册了 http 路由和 handler
func RegisterToHTTPServer(svr *http.Server, EpisodeSvc *Episode) (*HttpRegisterInfo, error) {
	var (
		info = &HttpRegisterInfo{
			registeredServices: []string{
				"templateV1",
			},
		}
	)
	v1.RegisterEpisodeHTTPServer(svr, EpisodeSvc)
	err := v1.RegisterEpisodeGuruServer(EpisodeSvc)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// RegisterToGRPCServer 将所有的服务注册到 grpc server
func RegisterToGRPCServer(svr *grpc.Server, EpisodeSvc *Episode) (*GrpcRegisterInfo, error) {
	var (
		info = &GrpcRegisterInfo{
			registeredServices: []string{
				"templateV1",
			},
		}
	)
	v1.RegisterEpisodeServer(svr, EpisodeSvc)
	err := v1.RegisterEpisodeGuruServer(EpisodeSvc)
	if err != nil {
		return nil, err
	}
	return info, nil
}
