# awesomeProject

## 整体架构
  整体使用grpc ➕ protobuf 的微服务架构

### 特定的功能。

## protoc 使用说明

### 编写 *.proto文件,示例如下：

```
// 协议版本
syntax = "proto3";
// 包名
package hello;
option go_package="../hello";

// 服务
service Hello {
rpc sayHello(HelloRequest) returns (HelloResponse) {}
}

// 消息格式
message HelloRequest {
string name = 1;
}

// 消息格式
message HelloResponse {
string message = 1;
}
```

## 各个服务介绍
### hello服务
一个简单的返回特定信息的服务，用来供参考
### user服务
### crawl服务
