package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `uri:"user"`
	Password string `uri:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/:user/:password", func(context *gin.Context) {
		var uri Login
		if err := context.ShouldBindUri(&uri); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if uri.User != "root" || uri.Password != "admin" {
			context.JSON(http.StatusBadRequest, gin.H{"status": 304})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": 200})
	})
	r.Run(":8000")
}
