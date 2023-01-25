# Reference Implementation of a RPC service provider

# Features
- 【完成】支持多个Service注册到同一个Provider服务
- 【完成】字段支持google.protobuf.Any类型
    - 支持Any字段的wrap和unwrap工具方法
- 【完成】字段支持Condition复合条件类型
- 【完成】支持dubbo-go协议
- 【完成】支持nacos注册中心
- 【完成】支持对protobuf定义的message添加自定义tag，例如v（数据校验）, json（通常用于去掉缺省存在的omitempty）等
- 【完成】支持输入结构体的validation
- 【完成】支持pixiu API网关的集成
- 【完成】支持为传入参数 Struct 设定 gmeta.Meta （可用于在orm操作时忽略 nil 值的字段）
- 【P1】支持链路跟踪
- 【P1】支持Stream形式的调用
- 【P2】支持go-zero
- 【P2】支持etcd注册中心
- 【P2】健康检测
- 【P2】支持国际化

# 运行 PRC service provider
`go run main.go`

# 测试
## RPC consumer
运行 `test/rpc_client` 下的 `TestCase`

## HTTP client
首先运行 pixiu API 网关
```
cd test/pixiu
sh run_pixiu.sh
```
然后可运行 `test/http_client` 下的 `TestCase`，也可以用 PostMan/cURL 等工具来测试

Note:
- bool 类型的序列化方式为： "true" | "false" 小写字符串