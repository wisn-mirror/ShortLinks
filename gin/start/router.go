package main

import (
	bytes2 "bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	//bindAllUrl()
	//getUrlParams()
	getFromData()
	//getBodyParamsGetUrlFirst()
}

func bindAllUrl() {
	g := gin.Default()
	//指定所有以 user 开头的url 都通过这个路由
	g.GET("/user/*action", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"FullPath": context.FullPath(),
			"params":   context.Params,
		})
	})
	g.Run()
}
//获取url中的参数
func getUrlParams()  {
	g:=gin.Default()
	g.GET("/test", func(context *gin.Context) {
		name := context.Query("name")
		id:=context.Query("id")
		context.JSON(200,gin.H{
			"name":name,
			"id":id,
		})
	})
	g.Run()
}
//获取表单中的数据  from-data , x-www-form-urlencoded 都兼容
func getFromData()  {
	g:=gin.Default()
	g.POST("/test", func(context *gin.Context) {
		name:=context.PostForm("name")
		ids:=context.PostFormArray("ids")
		context.JSON(200,gin.H{
			"name":name,
			"ids":ids,
		})
	})
	g.Run(":8088")
}

type testQuest struct {
	Name string `json:name`
}
//获取请求中的body ioutil.ReadAll读取(先获取) 和表单参数
func getBodyParamsGetUrlFirst()  {
	g:=gin.Default()
	g.POST("test", func(context *gin.Context) {
		name:=context.Query("name")
		id:=context.Query("id")

		bytes,error:=ioutil.ReadAll(context.Request.Body)
		if error != nil {
			fmt.Println(error)
			context.Abort()
			return
		}
		//如果Post请求  ioutil.ReadAll(先获取)表单参数读取无效
		//要重新给数据放到请求中
		context.Request.Body=ioutil.NopCloser(bytes2.NewBuffer(bytes))
		from1value:=context.PostForm("from1")
		context.JSON(200,gin.H{
			"context":string(bytes),
			"name":name,
			"id":id,
			"from1value":from1value,
		})
	})
	g.Run()
}
