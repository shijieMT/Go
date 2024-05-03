# 安装开发环境
## Protocol Buffers安装
- 第一步：下载通用编译器

  地址：https://github.com/protocolbuffers/protobuf/releases

  根据不同的操作系统，下载不同的包，我是windows电脑，解压出来是`protoc.exe`

- 第二步：配置环境变量
  
  ![image-20220423002031614](学习文档/img/image-20220423002031614.png)

- 第三步：检查protoc命令是否可以使用
  
~~~bash
protoc
~~~

> 如何使用protobuf呢？

1. 定义了一种源文件，扩展名为 `.proto`，使用这种源文件，可以定义存储类的内容(消息类型)
2. protobuf有自己的编译器 `protoc`，可以将 `.proto` 编译成对应语言的文件，就可以进行使用了
## 安装gRPC核心库（新版）

~~~go
go instal1 google.go7ang.org/protobuf/cmd/protoc-gen-go@latest
go instal1 google.go7ang.org/grpc/cmd/protoc-gen-go-grpc@latest
~~~
# 学习资料
## todo
