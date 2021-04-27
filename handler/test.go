package handler

import (
	"github.com/fullstacker-go/concurr_gin/model"
	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {
	id := c.Param("id")
	website := "https://www." + id
	size := make(chan model.WebStats)
	go model.ResponseSize(website, size)
	c.JSON(200, <-size)

}
