package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ginDemo1()
	//requestType()
	//staticFile()
	//urlParam()
}

//gin 入门
func ginDemo1() {
	g := gin.Default()
	g.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code":    200,
			"message": "pong",
			"data":    "hello",
		})
	})
	g.Run(":8080")
}

//各种请求 get post delete put ...
func requestType() {
	engine := gin.Default()
	engine.GET("/get", func(context *gin.Context) {
		context.String(200, "get")
	})
	engine.POST("/post", func(context *gin.Context) {
		context.String(200, "post")
	})
	engine.DELETE("/delete", func(context *gin.Context) {
		context.String(200, "delete")
	})
	engine.PUT("/put", func(context *gin.Context) {
		context.String(200, "put")
	})
	engine.HEAD("/head", func(context *gin.Context) {
		context.String(200, "head")
	})
	//any 支持所有的http请求方式 9种
	engine.Any("/any", func(context *gin.Context) {
		context.String(200, "any")
	})
	engine.Run()
}

//静态文件夹绑定
func staticFile()  {
	engine := gin.Default()
	//一种写法 这种必须使用全路径指定到具体文件
	engine.Static("/assets","./assetsdir")
	//第二种写法  这种相当于ftp，可以看到文件列表
	engine.StaticFS("/static",http.Dir("./staticdir"))
	//指定单个文件  这种可以指定单个文件，其他文件无效
	engine.StaticFile("/favicon.ico","./assetsdir/image.jpg")
	engine.Run()
}

//通过url 传递参数 严格的参数列表
func urlParam()  {
	g:=gin.Default()
	g.GET("/test1/:name/:id", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"name":context.Param("name"),
			"id":context.Param("id"),
		})
	})

	g.GET("/test2/:name/aa/:id", func(context *gin.Context) {

		context.JSON(200,gin.H{
			"test2Name":context.Param("name"),
			"id":context.Param("id"),
		})
	})
	g.Run()
}

