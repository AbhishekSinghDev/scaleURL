package main

import (
	"github.com/AbhishekSinghDev/scaleURL/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(middleware.RequestIdMiddleware(), middleware.LoggerMiddleware())

	router.Run()
}
