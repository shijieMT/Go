# 一、体验Kratos项目helloworld
### 1. 创建项目
```shell
# 安装依赖
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
go install github.com/google/wire/cmd/wire@latest
# 创建项目-helloworld
kratos new helloworld
```
### ~~2. 编写proto文件~~(跳过，转到第三步)
```shell
kratos proto add api/helloworld/v1/demo.proto
```
将文件内容修改为我们想要实现的内容-SayHello
```protobuf
syntax = "proto3";

package helloworld.v1;

import "google/api/annotations.proto";

option go_package = "github.com/go-kratos/service-layout/api/helloworld/v1;v1";
option java_multiple_files = true;
option java_package = "dev.kratos.api.helloworld.v1";
option java_outer_classname = "HelloWorldProtoV1";

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply)  {
    option (google.api.http) = {
        // 定义一个 GET 接口，并且把 name 映射到 HelloRequest
        get: "/helloworld/{name}",
        // 可以添加附加接口
        additional_bindings {
            // 定义一个 POST 接口，并且把 body 映射到 HelloRequest
            post: "/v1/greeter/say_hello",
            body: "*",
        }
    };
  }
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```
### 3. 运行程序
```shell
go mod tidy
# 运行项目
kratos run
```
构建项目时，程序自带了一个api/helloworld/v1/greeter.proto，刚好是我们需要的proto文件  
访问 http://127.0.0.1:8000/helloworld/ShiJie 成功响应

# 二、自己构建helloworld项目demo
> 在第一步中，已经可以成功创建项目，并成功运行，但使用的是预设好的代码，并不是我们自己编写的  
> 这次来自己构建项目，并完成SayHello功能

### 1. 创建项目
```shell
# 删除之前的项目
rm -r ./*
# cli-构建项目
kratos new helloworld
# 删除已有的greeter相关文件
rm ./helloworld/api/helloworld/v1/greeter*
rm ./helloworld/internal/service/greeter.go
```
### 2. 创建proto文件，生成并更改文件
```shell
# 进入项目目录
cd ./helloworld
# 创建proto文件
kratos proto add api/helloworld/v1/demo.proto
```
#### 2.1. 编写demo.proto
```protobuf
syntax = "proto3";

package helloworld.v1;

import "google/api/annotations.proto";

option go_package = "github.com/go-kratos/service-layout/api/helloworld/v1;v1";
option java_multiple_files = true;
option java_package = "dev.kratos.api.helloworld.v1";
option java_outer_classname = "HelloWorldProtoV1";

// The greeting service definition.
service Demo {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply)  {
    option (google.api.http) = {
        // 定义一个 GET 接口，并且把 name 映射到 HelloRequest
        get: "/helloworld/{name}",
        // 可以添加附加接口
        additional_bindings {
            // 定义一个 POST 接口，并且把 body 映射到 HelloRequest
            post: "/v1/greeter/say_hello",
            body: "*",
        }
    };
  }
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```
#### 2.2. 用这个proto文件生成我们需要的代码
```shell
# 生成 client 源码
kratos proto client api/helloworld/v1/demo.proto
# pb.go grpc.go http.go 文件已经生成完毕
```
#### 2.3. 修改internal/service
```shell
# 生成 server 源码
kratos proto server api/helloworld/v1/demo.proto -t internal/service
# internal/service/demo.go 文件已经生成完毕
```
根据greeter.go修改demo.go(biz部分暂时没有做处理)
```go
package service

import (
	"context"

	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
)

type DemoService struct {
	v1.UnimplementedDemoServer

	uc *biz.GreeterUsecase
}

// NewDemoService new a Demo service.
func NewDemoService(uc *biz.GreeterUsecase) *DemoService {
	return &DemoService{uc: uc}
}

func (s *DemoService) SayHello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "看到界面时，您已完成整个流程"}, nil
}
```
修改service.go
```go
package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewDemoService)
```
#### 2.4. 修改internal/server 的grpc和http文件（修改标红部分）  
   grpc代码修改
```go
package server

import (
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, demo *service.DemoService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterDemoServer(srv, demo)
	return srv
}
```
http代码修改
```go
package server

import (
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, demo *service.DemoService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterDemoHTTPServer(srv, demo)
	return srv
}
```
### 3. 使用wire命令，重新生成wire_gen，并启动项目
重新生成wire_gen
```shell
# Mac -> cmd/helloworld/
wire
# Win -> cmd/helloworld/
go run -mod=mod github.com/google/wire/cmd/wire
```

启动项目

```shell
# 运行项目
kratos run
```
浏览器访问  
http://127.0.0.1:8000/helloworld/Shijie  
返回结果：
```json
{
    "message": "看到界面时，您已完成整个流程"
}
```
