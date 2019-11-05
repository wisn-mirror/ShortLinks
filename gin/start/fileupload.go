package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	//singleFileUpload()
	moreFileUpload()
}
//单文件上传
func singleFileUpload() {
	g := gin.Default()
	g.POST("/upload", func(context *gin.Context) {
		if file, error := context.FormFile("file"); error != nil {
			context.JSON(500, gin.H{
				"code":    500,
				"message": error.Error(),
			})
			context.Abort()
			return
		} else {
			saveFileError := context.SaveUploadedFile(file, "./upload/"+file.Filename)
			if saveFileError!=nil{
				context.JSON(500, gin.H{
					"code":    500,
					"message": saveFileError.Error(),
				})
				context.Abort()
				return
			}
			context.JSON(200, gin.H{
				"code":    200,
				"message": file.Filename,
			})
		}
	})
	g.Run()
}

type fileloadStatus struct {
	FileName string `json:"filename"`
	IsSuccess bool `json:"issuccess"`
}

//多文件上传
func moreFileUpload() {
	g := gin.Default()
	g.POST("/upload", func(context *gin.Context) {
		form, multipartError := context.MultipartForm()
		if multipartError!=nil{
			context.JSON(500, gin.H{
				"code":    500,
				"message": multipartError.Error(),
			})
			context.Abort()
			return
		}
		files:=form.File["files"]
		if files==nil||len(files)==0{
			context.JSON(200, gin.H{
				"code":    500,
				"message": "缺少必要参数",
			})
			return
		}
		var filestatus []fileloadStatus
		for index,file:=range files{
			saveFileError :=context.SaveUploadedFile(file,"./upload/"+strconv.Itoa(index)+"_"+file.Filename)
			if saveFileError!=nil{
				filestatus=append(filestatus,fileloadStatus{
					FileName:file.Filename,
					IsSuccess:false,
				})
			}else{
				filestatus=append(filestatus,fileloadStatus{
					FileName:file.Filename,
					IsSuccess:true,
				})
			}
		}

		context.JSON(200, gin.H{
			"code":    200,
			"message": "成功",
			"data":filestatus,
		})

	})
	g.Run()
}