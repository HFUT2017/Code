package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.POST("loginForm", func(context *gin.Context) {
		var form Login
		if err := context.Bind(&form); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "root" || form.Password != "admin" {
			context.JSON(http.StatusBadRequest, gin.H{"status": 304})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": 200})
	})
	r.Run(":8000")
}
