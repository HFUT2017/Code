package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			context.String(http.StatusInternalServerError, "上传失败!")
		}
		context.SaveUploadedFile(file, file.Filename)
		context.String(http.StatusOK, fmt.Sprintf("上传成功!filename:%s", file.Filename))
	})
	r.Run(":8000")
}
