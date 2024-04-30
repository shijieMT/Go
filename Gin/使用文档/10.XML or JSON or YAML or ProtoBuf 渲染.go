package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func func10() {
	r := gin.Default()

	// gin.H 是 map[string]interface{} 的一种快捷方式
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// 你也可以使用一个结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意 msg.Name 在 JSON 中变成了 "user"
		// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

/*
Protobuf（Protocol Buffers）是一种由Google开发的跨语言、跨平台的序列化框架。
它主要用于定义数据结构和协议，以便于数据的序列化和反序列化。Protobuf在很多方面都非常有用：
1.
数据交换格式：Protobuf提供了一种紧凑、高效的数据交换格式，特别适合于网络通信和数据存储。
由于其二进制格式，Protobuf序列化后的数据通常比JSON或XML等文本格式更小，传输速度更快。
2.
跨语言兼容性：Protobuf支持多种编程语言，包括但不限于C++, Java, Python, Go等。
这意味着可以用一种语言定义数据结构，然后在其他语言中使用这些结构，这对于多语言环境下的项目非常有用。
3.
向前和向后兼容性：Protobuf设计时考虑了版本兼容性问题。
即使数据结构发生变化，旧版本的代码仍然可以读取新版本的数据，反之亦然，只要遵循一定的规则。
4.
性能：由于其高效的序列化和反序列化机制，Protobuf在性能方面通常优于JSON或XML等其他数据交换格式。
*/
