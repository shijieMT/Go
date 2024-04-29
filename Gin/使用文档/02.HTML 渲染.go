package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func func2() {
	// 加载HTML模板文件()
	// 使用不同目录下名称相同的模板()
	// todo 自定义模板渲染器 自定义分隔符 自定义模板功能
}

func 使用不同目录下名称相同的模板() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	// 如果文件名字唯一，路径名只需要写index.tmpl即可访问，无论该文件在哪一个文件夹
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.Run(":8080")
	/*	templates/posts/index.tmpl
		{{ define "posts/index.tmpl" }}
		<html><h1>
			{{ .title }}
		</h1>
		<p>Using posts/index.tmpl</p>
		</html>
		{{ end }}
	*/
	/*	templates/users/index.tmpl
		{{ define "users/index.tmpl" }}
		<html><h1>
			{{ .title }}
		</h1>
		<p>Using users/index.tmpl</p>
		</html>
		{{ end }}
	*/
}
func 加载HTML模板文件() {
	router := gin.Default()
	// 加载HTML模板文件
	// 这个方法允许使用通配符（glob pattern）来指定模板文件的路径，这样可以一次性加载多个模板文件。
	router.LoadHTMLGlob("templates/*")
	// router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
	/*	templates/index.tmpl
		<html>
			<h1>
				{{ .title }}
			</h1>
		</html>
	*/
}
