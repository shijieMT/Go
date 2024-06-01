# 编写简单微服务demo
## go-zero训练：根据要求编写对应微服务demo编写（rpc和api）
> 服务端实现User服务（内含getUser方法），接受id，返回id，name，gender  
> 客户端实现Video服务（内含getVideo方法），get请求，接受id，返回id和name（根据id调用getUser获取对应name）  
> [参考资料](https://go-zero.dev/docs/tasks)
## 编写user模块
### 编写proto文件
/user/rpc/user.proto
```protobuf
syntax = "proto3";
package user;
option go_package = "./user";

message IdRequest{
  string id = 1;
}

message UserResponse{
   string id = 1;
   string name = 2;
   string gender = 3;
}

service User{
  rpc getUser(IdRequest) returns (UserResponse);
}
```
### 1. 生成文件
```shell
goctl rpc protoc user/rpc/user.proto  --go_out=user/rpc/types --go-grpc_out=user/rpc/types --zrpc_out=user/rpc
```
### 2. 编写 internal/logic/getuserlogic.go
```go
func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	return &user.UserResponse{
		Id:     string("1234"),
		Name:   string("玛奇朵"),
		Gender: string("女"),
	}, nil
}
```
### 3. 整理模块资源关系
```shell
go mod tidy
```
### 4. 运行user.go
```shell
go run .\user.go
```
### 5. 使用Postman或Apifox进行测试
略
## 编写video模块
### 1. 编写video.api
/video/api/video.api
```api
type (
	VideoReq {
		Id string `path:"id"`
	}
	VideoRes {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
)

service video {
	@handler getVideo
	get /api/videos/:id (VideoReq) returns (VideoRes)
}
```
### 2. 生成代码
```shell
goctl api go -api video/api/video.api -dir video/api/
```
### 3. 添加user rpc配置
> 因为要在video里面调用user的rpc服务

video/api/internal/config/config.go

```go
package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// 新增 UserRpc zrpc.RpcClientConf
	UserRpc zrpc.RpcClientConf
}
```
### 4.完善服务依赖
video/api/internal/svc/servicecontext.go
```go
package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero_study/user/rpc/userclient"
	"go-zero_study/video/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	// 新增 UserRpc userclient.User
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// 新增 UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
```
### 添加yaml配置
video/api/etc/video.yaml
```yaml
Name: video
Host: 0.0.0.0
Port: 8888
# 新增
UserRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: user.rpc
```
### 完善服务依赖
video/api/internal/logic/getvideologic.go
```go
...
func (l *GetVideoLogic) GetVideo(req *types.VideoReq) (resp *types.VideoRes, err error) {
	// todo: add your logic here and delete this line
	user1, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: "1",
	})
	if err != nil {
		return nil, err
	}
	return &types.VideoRes{
		Id:   req.Id,
		Name: user1.Name,
	}, nil
}
```
## 启动程序
运行user rpc
```shell
go run user\rpc\user.go -f user\rpc\etc\user.yaml
```
运行video api
```shell
go run video\api\video.go -f video\api\etc\video.yaml
```
测试：
```shell
curl 127.0.0.1:8888/api/videos/1
```
## 回顾操作
> 1. 编写用户微服务的rpc服务的proto文件  
> 2. 生成代码  
> 3. 添加自己的逻辑  
> 4. 编写视频微服务的api服务的api文件  
> 5. 生成代码  
> 6. 完善依赖，配置  
> 7. 添加自己的逻辑  
