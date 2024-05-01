package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*
博客链接：
https://gitee.com/moxi159753/LearningNotes/tree/master/Golang/Gin%E6%A1%86%E6%9E%B6/1_Gin%E5%86%85%E5%AE%B9%E4%BB%8B%E7%BB%8D#gin%E4%B8%AD%E9%97%B4%E4%BB%B6
*/

// todo 了解中间件
func func12() {
	// 创建一个新的Gin路由器实例
	r := gin.New()

	// 添加自定义中间件
	r.Use(myMiddleware)

	// 添加内置的Logger中间件
	r.Use(gin.Logger())

	// 添加内置的Recovery中间件
	r.Use(gin.Recovery())

	// 定义路由和处理函数
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	log.Fatal(r.Run(":8080"))
}

// 自定义中间件
func myMiddleware(c *gin.Context) {
	// 在请求处理之前执行的代码
	log.Println("请求到达自定义中间件")

	// 调用下一个中间件或处理函数
	c.Next()

	// 在请求处理之后执行的代码
	log.Println("请求处理完毕")
}

/*
	Gin框架内置了两个常用的中间件：Logger 和 Recovery。
1.
Logger 中间件：
Logger 中间件用于记录HTTP请求的详细信息，包括请求的路径、请求方法、响应状态码、请求耗时等。
这对于调试和监控应用程序非常有用。Logger 中间件默认会将日志输出到标准输出（stdout），但你也可以自定义日志的输出方式。
2.
Recovery 中间件：
Recovery 中间件用于捕获并处理运行时的恐慌（panic），并返回一个500内部服务器错误响应。
这可以防止服务器崩溃，并且可以提供一个友好的错误页面给用户。Recovery 中间件默认会返回一个包含错误堆栈的JSON响应。
*/
