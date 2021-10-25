package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(context *gin.Context) {
		name := context.DefaultQuery("name", "枯藤")
		context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "404 not found2222")
	})
	r.Run(":8000")
}
