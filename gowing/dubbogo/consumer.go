package dubbogo

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/natefinch/lumberjack"
	"path"
)

var consumerConfigBuilder = config.NewConsumerConfigBuilder()

// AddConsumerReference 增加一个 consumer 的依赖配置
func AddConsumerReference(consumer *ConsumerReference) {
	consumerConfigBuilder.AddReference(consumer.ClientImplStructName,
		config.NewReferenceConfigBuilder().
			SetProtocol(consumer.Protocol).
			Build())
	config.SetConsumerService(consumer.Service)
}

func buildConsumerConfig(builder *config.ConsumerConfigBuilder, consumerOption *ConsumerOption) *config.ConsumerConfig {
	if consumerOption.TimeoutSeconds > 0 {
		builder.SetRequestTimeout(gconv.String(consumerOption.TimeoutSeconds) + "s")
	}
	return builder.SetCheck(consumerOption.CheckProviderExists).Build()
}

// StartConsumers 启动通过 AddConsumerReference 添加的所有的 Consumer
func StartConsumers(_ context.Context, registry *Registry, consumerOption *ConsumerOption, loggerOption *LoggerOption) error {
	consumerConfig := buildConsumerConfig(consumerConfigBuilder, consumerOption)
	if len(consumerConfig.References) == 0 {
		// return when there are no consumer references
		return nil
	}
	registryConfigBuilder := config.NewRegistryConfigBuilder().
		SetProtocol(registry.Type).
		SetAddress(registry.Address)
	if registry.Type == "nacos" && !g.IsEmpty(registry.Namespace) {
		registryConfigBuilder = registryConfigBuilder.SetNamespace(registry.Namespace)
	}

	var (
		loggerOutputPaths      []string
		loggerErrorOutputPaths []string
	)
	if loggerOption.Stdout {
		loggerOutputPaths = []string{"stdout", loggerOption.LogDir}
		loggerErrorOutputPaths = []string{"stderr", loggerOption.LogDir}
	} else {
		loggerOutputPaths = []string{loggerOption.LogDir}
		loggerErrorOutputPaths = []string{loggerOption.LogDir}
	}

	registryConfigBuilder.SetParams(map[string]string{
		constant.NacosLogDirKey:   loggerOption.LogDir,
		constant.NacosCacheDirKey: loggerOption.LogDir,
		constant.NacosLogLevelKey: loggerOption.Level,
	})

	rootConfig := config.NewRootConfigBuilder().
		AddRegistry(registry.Id, registryConfigBuilder.Build()).
		SetLogger(config.NewLoggerConfigBuilder().
			SetZapConfig(config.ZapConfig{
				Level:            loggerOption.Level,
				Development:      loggerOption.Development,
				OutputPaths:      loggerOutputPaths,
				ErrorOutputPaths: loggerErrorOutputPaths,
			}).
			SetLumberjackConfig(&lumberjack.Logger{
				Filename: path.Join(loggerOption.LogDir, loggerOption.LogFileName),
			}).Build()).
		SetConsumer(consumerConfig).
		Build()
	if err := config.Load(config.WithRootConfig(rootConfig)); err != nil {
		return err
	}
	return nil
}
