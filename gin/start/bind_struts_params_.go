package main

import "github.com/gin-gonic/gin"

type Request struct {
	Name string	`form:"name"`
	Id int	`form:"id"`
}
func main() {
	g:=gin.Default()
	g.GET("test", handleRequet)
	g.POST("test", handleRequet)
	g.Run()
}

func handleRequet(context *gin.Context) {
	var request Request
	//  get取url参数中，post通过content-type 来判断 form-data urlencoded application/json
	error := context.ShouldBind(&request)
	if error!=nil{
		context.Abort()
		return
	}
	context.JSON(200, gin.H{
		"name":request.Name,
		"id":request.Id,
	})
}
