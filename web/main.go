package main

import (
	"RTB-go/internal/pkg/env"
	"RTB-go/web/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(env.Mode())
	router := gin.New()
	router.GET("/index", handler.Index)
	err := router.Run(":8080")
	panic(err)
}
