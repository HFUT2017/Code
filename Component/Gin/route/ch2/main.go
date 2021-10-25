package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(context *gin.Context) {
		name := context.DefaultQuery("name", "用户")
		context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.Run(":8000")
}
