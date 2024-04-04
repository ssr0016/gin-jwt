package main

import (
	"gin-jwt/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	server := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
