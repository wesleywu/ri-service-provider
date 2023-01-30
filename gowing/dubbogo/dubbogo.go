package dubbogo

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/natefinch/lumberjack"
	"path"
)

const serverFilters = "echo, metrics, token, accesslog, tps, pshutdown, tracing, otelServerTrace, validation, cache"

func StartProvider(_ context.Context, registry *Registry, provider *ProviderInfo, logger *LoggerOption) error {
	if provider.Port <= 100 {
		return gerror.New("需要指定大于100的 port 参数，建议20000以上，不能和其他服务重复")
	}

	var (
		loggerOutputPaths      []string
		loggerErrorOutputPaths []string
	)

	if logger.Stdout {
		loggerOutputPaths = []string{"stdout", logger.LogDir}
		loggerErrorOutputPaths = []string{"stderr", logger.LogDir}
	} else {
		loggerOutputPaths = []string{logger.LogDir}
		loggerErrorOutputPaths = []string{logger.LogDir}
	}

	// application config
	applicationBuilder := config.NewApplicationConfigBuilder()
	applicationBuilder.SetName(provider.ApplicationName)

	// registry config
	registryConfigBuilder := config.NewRegistryConfigBuilder().
		SetProtocol(registry.Type).
		SetAddress(registry.Address)
	if registry.Type == "nacos" && !g.IsEmpty(registry.Namespace) {
		registryConfigBuilder = registryConfigBuilder.SetNamespace(registry.Namespace)
	}
	registryConfigBuilder.SetParams(map[string]string{
		constant.NacosLogDirKey:   logger.LogDir,
		constant.NacosCacheDirKey: logger.LogDir,
		constant.NacosLogLevelKey: logger.Level,
	})

	// metadata report config
	metadataReportConfigBuilder := config.NewMetadataReportConfigBuilder().
		SetProtocol(registry.Type).
		SetAddress(registry.Address)

	// protocol config
	protocolConfigBuilder := config.NewProtocolConfigBuilder().
		SetName(provider.Protocol).
		SetPort(gconv.String(provider.Port))
	if !g.IsEmpty(provider.IP) {
		protocolConfigBuilder = protocolConfigBuilder.SetIp(provider.IP)
	}

	// provider config
	providerConfigBuilder := config.NewProviderConfigBuilder()
	for _, service := range provider.Services {
		config.SetProviderService(service.Service)
		providerConfigBuilder.AddService(service.ServerImplStructName,
			config.NewServiceConfigBuilder().SetInterface("").Build())
	}
	providerConfigBuilder.SetFilter(serverFilters)

	// shutdown callbacks
	if provider.ShutdownCallbacks != nil {
		extension.AddCustomShutdownCallback(func() {
			for _, callback := range provider.ShutdownCallbacks {
				callback()
			}
		})
	}

	// logger config
	loggerConfigBuilder := config.NewLoggerConfigBuilder()
	loggerConfigBuilder.SetZapConfig(config.ZapConfig{
		Level:            logger.Level,
		Development:      logger.Development,
		OutputPaths:      loggerOutputPaths,
		ErrorOutputPaths: loggerErrorOutputPaths,
	})
	loggerConfigBuilder.SetLumberjackConfig(&lumberjack.Logger{
		Filename: path.Join(logger.LogDir, logger.LogFileName),
	})

	rootConfig := config.NewRootConfigBuilder().
		SetApplication(applicationBuilder.Build()).
		AddRegistry(registry.Id, registryConfigBuilder.Build()).
		SetMetadataReport(metadataReportConfigBuilder.Build()).
		AddProtocol("tripleKey", protocolConfigBuilder.Build()).
		SetProvider(providerConfigBuilder.Build()).
		SetLogger(loggerConfigBuilder.Build()).
		Build()
	if err := config.Load(config.WithRootConfig(rootConfig)); err != nil {
		return err
	}
	select {}
}
