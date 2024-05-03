# 示例编写步骤
## proto
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
## 代码生成
进入proto文件夹下，执行代码
>  **.**   代表目录  
>  **hello.proto**   为文件名
~~~dos
protoc --go_out=. hello.proto
protoc --go-grpc_out=. hello.proto
~~~
## 生成文件
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
### 服务端编写 todo
创建gRPC Server 对象，你可以理解为它是 Server 端的抽象对象
将 server (其包含需要被调用的服务端接口)注册到gRPC Server的内部注册中心。
这样可以在接受到请求时，通过内部的服务发现，发现该服务端接口并转接进行逻辑处理
创建Listen，监听TCP端口
gRPC Server开始lis.Accept，直到Stop
### 客户端编写 todo
客户端编写
创建与给定目标(服务端)的连接交互
创建server的客户端对象
发送RPC请求，等待同步响应，得到回调后返回响应结果
输出响应结果

