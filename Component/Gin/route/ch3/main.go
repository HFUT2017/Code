package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		action = strings.Trim(action, "/")
		context.String(http.StatusOK, name+" is "+action)
	})
	r.Run(":8000")
}
