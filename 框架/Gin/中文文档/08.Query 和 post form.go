package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=manu&message=this_is_great
*/
/*
	todo 使用curl和postman测试
*/
func func8() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		// id: 1234; page: 1; name: manu; message: this_is_great
	})
	router.Run(":8080")
}

/*	本地创建 html，点击按钮发送 post请求
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Test POST Request</title>
</head>
<body>
<form action="http://localhost:8080/post" method="post">
    <input type="hidden" name="id" value="123">
    <input type="hidden" name="name" value="John Doe">
    <input type="hidden" name="message" value="Hello World">
    <input type="submit" value="Submit">
</form>
</body>
</html>
*/
