# 一、体验helloworld项目
### 1. 创建项目
```shell
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
kratos new helloworld
```
运行时，调整工作目录为main.go所在目录 helloworld/cmd/helloworld/main.go
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
# 安装依赖
go get github.com/google/wire/cmd/wire@latest
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
修改service
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
# 三、连接数据库与配置修改
### 1. 准备数据库实例
创建helloworld/deploy/mysql/docker-compose.yaml文件
```yaml
version: '3'
services:
  demo_db:
    image: mysql:8
    environment:
      MYSQL_ROOT_PASSWORD: dangerous
      MYSQL_DATABASE: demo # for database creation automatically
    ports:
      - 3306:3306
    volumes:
      - "./data:/var/lib/mysql"
```
启动数据库，成功
### 2. 引入gorm
使用gorm(此处默认您已经有过gorm的使用经验了)  
[以前写的简易gorm使用指南](https://github.com/shijieMT/Go/tree/main/Gorm)
```shell
go get gorm.io/driver/mysql
go get gorm.io/gorm
```
### 3. 修改 helloworld/internal/data/data.go
3.1. 新增NewDB方法  
3.2. 修改NewData方法  
3.3. 将NewDB添加到ProviderSet  
```go
package data

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"helloworld/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewGreeterRepo)

// Data .
type Data struct {
	DB *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{DB: db}, cleanup, nil
}

// NewDB .
func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	username := "root"      //账号
	password := "dangerous" //密码
	host := "127.0.0.1"     //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "demo"        //数据库名
	timeout := "10s"        //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}
	return db
}
```
### 4. 更新wire_gen.go
```shell
go run -mod=mod github.com/google/wire/cmd/wire
```
成功生成新的wire_gen.go
### 5. 将dsn添加到配置文件
#### 5.1. 更改 helloworld/configs/config.yaml
```yaml
server:
  http:
    addr: 0.0.0.0:8000
    timeout: 1s
  grpc:
    addr: 0.0.0.0:9000
    timeout: 1s
data:
  database:
    username : "root"
    password : "dangerous"
    host : "127.0.0.1"
    port : 3306
    Dbname : "demo"
    timeout : "10s"
  redis:
    addr: 127.0.0.1:6379
    read_timeout: 0.2s
    write_timeout: 0.2s
```
#### 5.2. 更改 helloworld/internal/conf/conf.proto 中的 Data部分（使其与配置文件结构一致）
```protobuf
syntax = "proto3";
package kratos.api;

option go_package = "helloworld/internal/conf;conf";

import "google/protobuf/duration.proto";

message Bootstrap {
  Server server = 1;
  Data data = 2;
}

message Server {
  message HTTP {
    string network = 1;
    string addr = 2;
    google.protobuf.Duration timeout = 3;
  }
  message GRPC {
    string network = 1;
    string addr = 2;
    google.protobuf.Duration timeout = 3;
  }
  HTTP http = 1;
  GRPC grpc = 2;
}

message Data {
  message Database {
    string username = 1;
    string password = 2;
    string host = 3;
    int64 port = 4;
    string Dbname = 5;
    string timeout = 6;
  }
  message Redis {
    string network = 1;
    string addr = 2;
    google.protobuf.Duration read_timeout = 3;
    google.protobuf.Duration write_timeout = 4;
  }
  Database database = 1;
  Redis redis = 2;
}
```
#### 5.3. 根据 conf.proto生成 conf.pb.go
项目路径下执行：
```shell
# Win
protoc --proto_path=./internal --proto_path=./third_party --go_out=paths=source_relative:./internal ./internal/conf/conf.proto
# Mac
make config
```
#### 5.4. 更改 helloworld/internal/data/data.go（使用配置文件中的信息构建dsn）
此处只放了 func NewDB，其他部分不变
```go
// NewDB .
func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Dbname,
		c.Database.Timeout,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}
	return db
}
```
#### 5.5. 尝试运行，看是否通过编译
```shell
kratos run
```
没问题，进入下一步
# 三、 
### 1. 更改biz文件夹下的内容
#### 1.1. 新建 helloworld/internal/biz/demo.go
仿照greeter.go编写repo和usercase
```go
package biz

import (
	"context"

	v1 "helloworld/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Demo is a Demo model.
type Demo struct {
	Hello string
}

// DemoRepo is a Greater repo.
type DemoRepo interface {
	Save(context.Context, *Demo) (*Demo, error)
	ListAll(context.Context) ([]*Demo, error)
}

// DemoUsecase is a Demo usecase.
type DemoUsecase struct {
	repo DemoRepo
	log  *log.Helper
}

// NewDemoUsecase new a Demo usecase.
func NewDemoUsecase(repo DemoRepo, logger log.Logger) *DemoUsecase {
	return &DemoUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateDemo creates a Demo, and returns the new Demo.
func (uc *DemoUsecase) CreateDemo(ctx context.Context, g *Demo) (*Demo, error) {
	uc.log.WithContext(ctx).Infof("CreateDemo: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
```
#### 1.2. 更改biz.go
```go
package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewDemoUsecase)
```
### 2. 更改data文件夹下的内容
#### 2.1. 更新data.go的ProviderSet
```go
// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewDemoRepo)
```
#### 2.2. 更改demo.go(方法暂时保持空实现)
```go
package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"

	"context"
)

type DemoRepo struct {
	data *Data
	log  *log.Helper
}

// NewDemoRepo .
func NewDemoRepo(data *Data, logger log.Logger) biz.DemoRepo {
	return &DemoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *DemoRepo) Save(ctx context.Context, g *biz.Demo) (*biz.Demo, error) {
	return g, nil
}

func (r *DemoRepo) ListAll(context.Context) ([]*biz.Demo, error) {
	return nil, nil
}
```
### 3. 更改service中使用的UserCase（填之前没改的坑）
```go
package service

import (
	"context"

	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
)

type DemoService struct {
	v1.UnimplementedDemoServer

	uc *biz.DemoUsecase
}

// NewDemoService new a Demo service.
func NewDemoService(uc *biz.DemoUsecase) *DemoService {
	return &DemoService{uc: uc}
}

func (s *DemoService) SayHello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "看到界面时，您已完成整个流程"}, nil
}
```
### 4. 更改完毕，重新生成wire_gen.go
```shell
go run -mod=mod github.com/google/wire/cmd/wire
```
### 5. 运行程序，看是否通过编译
```shell
kratos run
```
成功运行，继续（如果MySQL没启动，会报MySQL连接失败）
### 6. 分析数据流向，按照自己的思路编写代码
此时，我们可以发现：  
> data文件夹中编写了操作数据库的方法（如func Save），  
> biz（业务逻辑的组装层）文件夹中的 func CreateDemo调用了 func Save，把数据库的数据封装成结构体对象（实例），  
> service（api 定义的服务层）文件夹中的 func SayHello可以调用 func CreateDemo，拿到实例对象，进行后续操作

所以，我们先实现data，然后实现biz，再实现service，生成wire_gen后运行，  
此时调用api，就可以完成我们想要的逻辑了
### 7. 分析wire_gen 的注入逻辑  
注入顺序：
```text
1. db           ->  NewDB(confData, logger)  
2. dataData     ->  data.NewData(confData, logger, db)  
3. demoRepo     ->  data.NewDemoRepo(dataData, logger)  
4. demoUsecase  ->  biz.NewDemoUsecase(demoRepo, logger)  
5. demoService  ->  service.NewDemoService(demoUsecase)  
6. grpcServer   ->  server.NewGRPCServer(confServer, demoService, logger)  
7. httpServer   ->  server.NewHTTPServer(confServer, demoService, logger)  
8. app          ->  newApp(logger, grpcServer, httpServer)  
```
文件与对应实体：
```text
data        ->      repo
biz         ->      Usecase
service     ->      Service
server      ->      http & grpc Server
main        ->      app
```
