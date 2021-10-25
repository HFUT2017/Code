package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			context.String(http.StatusInternalServerError, fmt.Sprintf("%v", err))
		}
		files := form.File["files"]
		for _, file := range files {
			if err := context.SaveUploadedFile(file, file.Filename); err != nil {
				context.String(http.StatusInternalServerError, fmt.Sprintf("%s上传失败!", file.Filename))
				return
			}
		}
		context.String(http.StatusOK, "上传成功!共上传%d个文件!", len(files))
	})
	r.Run(":8000")
}
