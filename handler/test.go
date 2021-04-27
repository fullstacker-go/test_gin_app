package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fullstacker-go/concurr_gin/model"
	"github.com/gin-gonic/gin"
)

func responseSize(url string, channel chan model.WebStats) {
	fmt.Println("Getting", url) // Unchanged
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Returning...", err)
		return
	}
	defer response.Body.Close()              // Unchanged
	body, _ := ioutil.ReadAll(response.Body) // Unchanged
	// Send body length value via channel.
	channel <- model.WebStats{url, len(body)}
}
func TestHandler(c *gin.Context) {
	id := c.Param("id")
	website := "https://" + id
	size := make(chan model.WebStats)
	go responseSize(website, size)
	c.JSON(200, <-size)

}
