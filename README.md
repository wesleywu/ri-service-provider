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
- 【P1】支持链路跟踪
- 【P1】支持Stream形式的调用
- 【P2】支持go-zero
- 【P2】支持etcd注册中心
- 【P2】健康检测
- 【P2】支持国际化
