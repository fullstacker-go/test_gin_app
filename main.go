package main

import (
	"github.com/fullstacker-go/concurr_gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.GET("/test/:id", handler.TestHandler)
	r.POST("/stats", handler.PostHandler)
	r.Run(":3000")
}
