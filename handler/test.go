package handler

import (
	"github.com/fullstacker-go/concurr_gin/model"
	"github.com/gin-gonic/gin"
)

// func TestHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	website := "https://www." + id
// 	size := make(chan model.WebStats)
// 	go model.ResponseSize(website, size)
// 	c.JSON(200, <-size)

// }
func PostHandler(c *gin.Context) {
	var websites []*model.WebStats
	c.Bind(&websites)
	size := make(chan int)
	for _, website := range websites {
		domain := "https://www." + website.Domain
		go model.ResponseSize(domain, size)
		website.Domain = domain
		website.Size = <-size
	}

	c.JSON(200, &websites)

}
