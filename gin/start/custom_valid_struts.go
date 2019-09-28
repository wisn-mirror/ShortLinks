package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type customValidRequest struct {
	CheckIn  time.Time `form:"check_in" binding:"required" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required" time_format:"2006-01-02"`
}

func main() {
	g := gin.Default()
	g.GET("/test", func(context *gin.Context) {
		var customValidRequest customValidRequest
		if error := context.ShouldBind(&customValidRequest); error != nil {
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
