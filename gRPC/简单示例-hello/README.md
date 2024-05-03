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

