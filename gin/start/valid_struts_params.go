package main

import "github.com/gin-gonic/gin"
//官方文档 validator.v8 支持多种验证
type validRequest struct {
	//验证参数 ，age 最小10
	Age int `form:"age" binding:"required,gt=10"`
	Name string `form:"name" binding:"required"`
}

func main() {
	g:=gin.Default()
	g.POST("/test", func(context *gin.Context) {
		var validRequest validRequest
		if error := context.ShouldBind(&validRequest);error!=nil{
			context.JSON(200,gin.H{
				"code":-1,
				"message":error.Error(),
			})
			context.Abort()
			return
		}
		context.JSON(200,gin.H{
			"code":200,
			"data":validRequest,
		})
	})
	g.Run()

}