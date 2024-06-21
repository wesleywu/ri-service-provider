# Reference Implementation of a RPC service provider

基于 GoFrame 和 Dubbo-Go Rpc框架开发的 RPC 服务，实现一个标准的 Domain Repository 服务，包括如下接口
- Count : 数据库表中符合条件的记录条数
- One   : 数据库表中符合条件的首条记录
- List  : 数据库表中符合条件的多条记录（支持翻页）
- Create: 在数据库表中创建记录
- Update: 在数据库表中更新特定记录
- Upsert: 在数据库表中创建或更新特定记录
- Delete: 从数据库表中删除符合条件的记录（可能为多条）

本项目的目标是，为 [gowing](https://github.com/wesleywu/gowing) 代码生成工具提供一个 Domain Repository 微服务的具体参考实现

## 功能说明
- 【完成】支持多个Service注册到同一个Provider服务
- 【完成】字段支持google.protobuf.Any类型
    - 提供Any字段的wrap和unwrap工具方法
- 【完成】字段支持Condition复合条件类型
- 【完成】支持dubbo-go协议
- 【完成】支持nacos/zookeeper注册中心
- 【完成】支持对protobuf定义的message添加自定义tag，例如v（数据校验）, json（通常用于去掉缺省存在的omitempty）
- 【完成】支持对输入结构体的validation
- 【完成】支持为传入参数 Struct 设定 gmeta.Meta （可用于在orm操作时忽略 nil 值的字段）
- 【完成】支持与Pixiu API网关的集成
- 【完成】支持 rpc call 的声明式透明缓存（对业务代码无侵入）
- 【完成】支持链路跟踪
- 【完成】当服务实现代码被引入时，支持进程内 (inproc) 服务调用，方便实现单体服务
- 【P1】更完善的错误处理
- 【P1】支持Stream形式的调用
- 【P2】支持etcd注册中心
- 【P2】健康检测
- 【P2】支持国际化

## 一、微服务方式运行

### 1. 运行 provider
#### - 创建 MySQL 数据库和表
`sql 脚本待提供` 
#### - 启动 Nacos
#### - 启动 PRC service provider
`go run ./provider --port 22000`

### 2. 运行 consumer 
#### 启动 restful-api，通过 RPC consumer 访问微服务 provider 
`go run ./api/cmd/rpc`

测试

`curl --location --request GET 'http://localhost:8200/app/episode/one?name=视频推荐'`

### 3. 测试dubbo-go provider

#### - RPC consumer
运行 `test/rpc_client/client_test.go` 里的 `TestCase`

#### - gRPC client
`dubbo-go`协议底层采用`gRPC`协议，可以采用`gRPC`客户端工具`evans`来测试。详细介绍请跳转 https://github.com/ktr0731/evans

```
# evans -r repl -p 22000
# 输出如下
  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client
proto.video_collection.Episode@127.0.0.1:22000>

# 查看grpc server提供了哪些服务和rpc接口
proto.video_collection.Episode@127.0.0.1:22000> show service

+-----------------+--------+--------------------------+--------------------------+
|     SERVICE     |  RPC   |       REQUEST TYPE       |      RESPONSE TYPE       |
+-----------------+--------+--------------------------+--------------------------+
| Episode | Count  | EpisodeCountReq  | EpisodeCountRes  |
| Episode | One    | EpisodeOneReq    | EpisodeOneRes    |
| Episode | List   | EpisodeListReq   | EpisodeListRes   |
| Episode | Create | EpisodeCreateReq | EpisodeCreateRes |
| Episode | Update | EpisodeUpdateReq | EpisodeUpdateRes |
| Episode | Upsert | EpisodeUpsertReq | EpisodeUpsertRes |
| Episode | Delete | EpisodeDeleteReq | EpisodeDeleteRes |
+-----------------+--------+--------------------------+--------------------------+

# 查看grpc server提供了哪些服务和rpc接口
proto.video_collection.Episode@127.0.0.1:22000> service Episode
# 已切换到 Episode 服务下，之后的 call 命令可以直接调用该服务下的 rpc method 了

proto.video_collection.Episode@127.0.0.1:22000> call One
✔ _
meta (TYPE_ENUM) => 
id::type_url (TYPE_STRING) => type.googleapis.com/google.protobuf.StringValue
id::value (TYPE_BYTES) => \n\r32559711-7365
name::type_url (TYPE_STRING) =>
name::value (TYPE_BYTES) =>
contentType::type_url (TYPE_STRING) =>
contentType::value (TYPE_BYTES) =>
filterType::type_url (TYPE_STRING) =>
filterType::value (TYPE_BYTES) =>
count::type_url (TYPE_STRING) =>
count::value (TYPE_BYTES) =>
isOnline::type_url (TYPE_STRING) =>
isOnline::value (TYPE_BYTES) =>
createdAt::type_url (TYPE_STRING) =>
createdAt::value (TYPE_BYTES) =>
updatedAt::type_url (TYPE_STRING) =>
updatedAt::value (TYPE_BYTES) =>
orderBy (TYPE_STRING) =>

# 结果如下
{
  "contentType": 1,
  "count": 0,
  "createdAt": "2023-01-10T13:36:58.111869Z",
  "filterType": 0,
  "id": "32559711-7365",
  "isOnline": false,
  "name": "",
  "updatedAt": "2023-01-10T15:10:26.002333Z"
}
```

#### - HTTP client
首先运行 pixiu API 网关

**需修改 conf.yaml 中的 address 为 nacos 的容器IP:Port （需要能够从docker容器中访问到）**
**pixiu API http(s) -> dubbo-go 协议转换无法支持otel链路跟踪，因为 pixiu 底层使用的开源框架 grpc-http-proxy 不支持 **

```
cd test/pixiu
docker run -p 8888:8888 \
    -v $(pwd)/conf.yaml:/etc/pixiu/conf.yaml \
    -v $(pwd)/log.yml:/etc/pixiu/log.yml \
    thehackercat/dubbo-go-pixiu-gateway:dubbo-go-pixiu-gateway-0.5.1-rc
```
然后可运行 `test/http_client/http_test.go` 里的 `TestCase`，也可以用 PostMan/cURL 等工具来测试

例如
```shell
curl --location --request POST 'http://localhost:8888/repo_service/Episode/List' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": {
        "@type":"type.googleapis.com/gowing.protobuf.Condition",
        "operator": "Like",
        "wildcard": "StartsWith",
        "value": {
            "@type":"type.googleapis.com/google.protobuf.StringValue",
            "value":"推荐"
        }
    },
    "contentType": {
        "@type":"type.googleapis.com/gowing.protobuf.Condition",
        "multi": "In",
        "value": {
            "@type":"type.googleapis.com/gowing.protobuf.Int32Slice",
            "value":[-1, 0, 1]
        }
    },
    "count": {
        "@type":"type.googleapis.com/gowing.protobuf.Condition",
        "operator": "GT",
        "value": {
            "@type":"type.googleapis.com/google.protobuf.Int32Value",
            "value":-1
        }
    }
}'
```

## 二、单体方式运行
### 1. 创建 MySQL 数据库和表
sql 脚本待提供

### 2. 启动单体 restful-api
`go run ./api/cmd/local`

测试

`curl --location --request GET 'http://localhost:8200/app/episode/one?name=视频推荐'`


## 三、附加说明
### Jaeger 链路跟踪
- 运行 Jaeger Docker 镜像

```
docker run \
  -p 14268:14268 \
  -p 16686:16686 \
  -p 5775:5775/udp \
  -p 5778:5778 \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 9411:9411 \
  -d jaegertracing/all-in-one:1.14
```

- 浏览器访问 Jaeger

`http://localhost:16686/`

### JSON 反序列化
**对于通过 pixiu API 网关代理的 HTTP -> RPC 请求，输入参数为json格式，需要特定的编码格式才能被 RPC 服务正确反序列化**
- 在 Create/Update/Upsert 请求中，*bool 类型字段的编码方式为： "true" | "false" 小写字符串
- 在 List/Count/One/Delete 请求中，请求结构体中的字段均为 google.protobuf.Any 类型，编码方式如下
  - 单值：
  ```
  "name": {
    "@type":"type.googleapis.com/google.protobuf.StringValue",
    "value":"视频推荐"
  },
  "isOnline": {
    "@type":"type.googleapis.com/google.protobuf.BoolValue",
    "value":false
  }
  ```
  - 数组：
  ```
  "name": {
    "@type":"type.googleapis.com/gowing.protobuf.StringSlice",
    "value":["视频推荐","腾讯视频推荐"]
  },
  "isOnline": {
    "@type":"type.googleapis.com/gowing.protobuf.BoolSlice",
    "value":[true, false]
  }
  ```
  - Condition复杂条件类型：
  ```
  "name": {
    "@type":"type.googleapis.com/gowing.protobuf.Condition",
    "operator": "Like",
    "wildcard": "StartsWith",
    "value": {
      "@type":"type.googleapis.com/google.protobuf.StringValue",
      "value":"每日"
    }
  }
  ```
  参见 `http_test.go`

### 为什么service方法返回值的结构体中，所有字段都是指针？
参考文档 
- https://protobuf.dev/programming-guides/field_presence/#how-to-enable-explicit-presence-in-proto3
- https://github.com/protocolbuffers/protobuf/blob/main/docs/field_presence.md

我们的需求是显式输出字段的值，无论其是否 zero value （例如 bool 类型的 false，int 类型的 0），这个需求在 protobuf 中称为
explicit presence (显式展现)。

一个使用 protobuf 序列化协议的 rpc 调用发生并返回结果时，protobuf 会对结果流式序列化 (stream serialization) 为字节流，
以方便进行网络传输。在proto3版本中，字段的缺省展现约束 (Presence disciplines) 是 no presence，那么当字段值为其类型的缺省值时，字段将不会被展现。
- 数字类型的缺省值为 0
- 枚举类型的缺省值为 0 值对应的枚举
- 字符串、字节数组、以及 repeated （即数组）类型的缺省值为 0 长度的值
- 自定义Message类型的缺省值为 语言相关的 null 值 （对于Go则是nil）

要想在 proto3 版本中设定字段的展现约束为 explicit presence ，需要在字段前加上 optional 限定词。
值得注意，proto2 版本中字段的缺省展现约束就是显式展现。

在加上 optional 之后，protoc 生成出来的 struct 源码中，字段就会被表述为类型的指针了。**还好，实际上并不影响对结果的使用。**

这个机制与 json 序列化的 omitempty 是类似的，no presence 相当于 omitempty。
为 json tag 添加 omitempty 也是 protoc 生成代码时，对字段tag的缺省行为，无法通过设置选项来修改这个缺省行为。
要绕过这个缺省行为，必须使用后处理方式，在 protoc 代码生成完毕后，重新修改一遍生成的代码，
本项目里使用是 `github.com/srikrsna/protoc-gen-gotag` 提供的方案。