package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/form", func(context *gin.Context) {
		types := context.DefaultPostForm("type", "post")
		username := context.PostForm("username")
		password := context.PostForm("passowrd")
		context.String(http.StatusOK, fmt.Sprintf("username %s password %s type %s", username, password, types))
	})
	r.Run(":8000")
}
