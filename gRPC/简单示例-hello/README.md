# 示例编写步骤
## proto文件编写
### 客户端
> hello-client/proto/hello.proto

~~~proto
syntax = "proto3";

package hello;

/*
option go_package = "import_path;package_name";
import_path 是生成的 Go 代码的导入路径，通常是一个相对于 GOPATH 的路径。
package_name 是生成的 Go 代码的包名。
 */
option go_package = ".;service";

// 定义一个服务，用于定义可以被调用的方法(可理解为 go的函数)
service HelloService {
  // 定义一个 RPC 方法，客户端可以通过该方法发送 Hello 请求
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// 定义 Hello 请求的消息结构（可理解为 go的结构体）
message HelloRequest {
  string name = 1; // 客户端发送的参数
}

// 定义 Hello 响应的消息结构
message HelloReply {
  string message = 1; // 服务端返回的消息
}
~~~

### 服务端
> hello-server/proto/hello.proto
> 代码同上
## pb.go代码生成
进入proto文件夹下，执行代码
>  **.**   代表目录  
>  **hello.proto**   为文件名
~~~dos
protoc --go_out=. hello.proto
protoc --go-grpc_out=. hello.proto
~~~
### 生成文件
> proto/hello.pb.go  （内含结构体定义）

~~~go
...
// 定义 Hello 请求的消息结构
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 客户端发送的参数
}
...
// 定义 Hello 响应的消息结构
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // 服务端返回的消息
}
...
~~~
> hello_grpc.pb.go  （内含方法实现）
~~~go hello.pb.go
...
type HelloServiceClient interface {
	// 定义一个 RPC 方法，客户端可以通过该方法发送 Hello 请求
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, HelloService_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
...
~~~
## 代码编写
### 1. 服务端编写
> 1. 开启TCP端口监听
> 2. 创建一个gRPC服务器实例
> 3. 注册gRPC服务到服务器实例中
> 4. 启动gRPC服务器，开始监听并处理来自客户端的请求
~~~go
package main

import (
	"context"
	service "gRPC_test/hello-client/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	service.UnimplementedHelloServiceServer
}

func (server) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloReply, error) {
	return &service.HelloReply{Message: string("我想发送的信息：" + req.Name + "你好！")}, nil
}

func main() {
	// 代码段：启动gRPC服务器并监听指定端口

	// 1. 开启TCP端口监听，指定监听本地的9090端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		// 如果监听端口时发生错误，则记录错误日志
		log.Printf("无法开启端口监听: %v\n", err)
		return // 发生错误时退出函数
	}

	// 2. 创建一个新的gRPC服务器实例
	grpcServer := grpc.NewServer()

	// 3. 注册gRPC服务到服务器实例中
	// 假设我们有一个名为service的包，其中包含RegisterHelloServiceServer函数
	// 该函数将HelloServiceServer的实现注册到gRPC服务器中
	service.RegisterHelloServiceServer(grpcServer, &server{})

	// 4. 启动gRPC服务器，开始监听并处理来自客户端的请求
	// 服务器将使用之前创建的监听器来接收连接
	if err := grpcServer.Serve(listen); err != nil {
		// 如果启动服务器时发生错误，则记录错误日志
		log.Printf("无法启动gRPC服务器: %v\n", err)
		return // 发生错误时退出函数
	}

	// 服务器启动后，将一直运行，直到程序被外部中断或调用Stop方法
}
~~~
### 2. 客户端编写
> 1. 创建与给定目标(服务端)的连接交互
> 2. 创建server的客户端对象
> 3. 执行RPC调用，等待同步响应，得到响应结果
> 4. 输出响应结果
~~~go
package main

import (
	"context"
	"fmt"
	service "gRPC_test/hello-client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// 代码段：创建gRPC客户端并连接到服务器，执行一个RPC调用

	// 1. 连接到gRPC服务器，这里使用了不安全的连接方式（禁用加密）
	// 使用grpc.WithTransportCredentials(insecure.NewCredentials())来指定不使用TLS加密
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// 如果连接服务器失败，则记录错误日志
		log.Printf("无法连接到服务器: %v\n", err)
		return // 发生错误时退出函数
	}
	// 使用defer语句确保在函数退出前关闭连接
	defer conn.Close()

	// 2. 创建客户端对象，这里假设service包中有一个NewHelloServiceClient函数
	// 该函数接受一个连接对象并返回一个客户端实例
	client := service.NewHelloServiceClient(conn)

	// 3. 执行RPC调用，这里调用的是客户端的SayHello方法
	// SayHello方法需要一个上下文对象和一个HelloRequest消息作为参数
	// 这里假设HelloRequest消息中有一个Name字段，我们将其设置为"焦糖玛奇朵"
	resp, err := client.SayHello(context.Background(), &service.HelloRequest{Name: "焦糖玛奇朵"})
	if err != nil {
		// 如果执行RPC调用时发生错误，则记录错误日志
		log.Printf("执行SayHello调用出错: %v\n", err)
		return // 发生错误时退出函数
	}

	// 4. 输出响应结果，这里假设HelloResponse消息中有一个GetMessage方法
	// 该方法返回一个字符串，我们将其打印到控制台
	fmt.Println(resp.GetMessage())
}

~~~
## 验证准备工作
### 1. 生成自签证书

> 生产环境可以购买证书或者使用一些平台发放的免费证书



* 安装openssl

  网站下载：http://slproweb.com/products/Win32OpenSSL.html

* 生成私钥文件

  ~~~shell
  ## 需要输入密码
  openssl genrsa -des3 -out ca.key 2048
  ~~~

* 创建证书请求

  ~~~shell
  openssl req -new -key ca.key -out ca.csr
  ~~~

* 生成ca.crt

  ~~~shell
  openssl x509 -req -days 365 -in ca.csr -signkey ca.key -out ca.crt
  ~~~

