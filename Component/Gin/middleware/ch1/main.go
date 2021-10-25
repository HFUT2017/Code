package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		context.Set("request", "中间件")
		status := context.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	r := gin.Default()
	r.Use(MiddleWare())
	{
		r.GET("/ce", func(context *gin.Context) {
			req, _ := context.Get("request")
			fmt.Println("request:", req)
			context.JSON(200, gin.H{"request": req})
		})
	}
	r.Run(":8000")
}
