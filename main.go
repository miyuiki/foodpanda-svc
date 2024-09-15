package main

import (
	"foodpanda-svc/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.GET("/check", service.CheckNewOrder)
	r.GET("/getOrder/:code", service.GetOrderByCode)
	r.Run(":8080")
}
