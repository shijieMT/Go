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
## 码神之路
[【码神之路】gRPC系列完整教程，go语言集成，十年大厂程序员讲解，通俗易懂](https://www.bilibili.com/video/BV16Z4y117yz?p=13&vd_source=d201ab3f18e3d32fee3a3605987bea6c)
> 前面使用的是老版本的gRPC，P13更换新版gRPC
## 狂神说
[【狂神说】gRPC最新超详细版教程通俗易懂 | Go语言全栈教程](https://www.bilibili.com/video/BV1S24y1U7Xp/?p=3&spm_id_from=333.1007.top_right_bar_window_history.content.click&vd_source=d201ab3f18e3d32fee3a3605987bea6c)

