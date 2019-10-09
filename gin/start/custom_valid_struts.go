package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type customValidRequest struct {
	//time_format:"2006-01-02"
	CheckIn  *time.Time `form:"check_in" binding:"required"`
	CheckOut *time.Time `form:"check_out" binding:"required"`
}

func main() {
	g := gin.Default()
	g.POST("/test", func(context *gin.Context) {
		var customValidRequest customValidRequest
		if error := context.ShouldBind(&customValidRequest); error != nil {
			//bytes,_:=ioutil.ReadAll(context.Request.Body)
			//fmt.Println(string(bytes))
			context.JSON(500, gin.H{
				"error": error.Error(),
			})
			context.Abort()
			return
		}

		context.JSON(200, gin.H{
			"message": "OK",
		})

	})
	g.Run()
}
