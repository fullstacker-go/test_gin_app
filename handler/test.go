package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func responseSize(url string, channel chan int) {
	fmt.Println("Getting", url)              // Unchanged
	response, _ := http.Get(url)             // Unchanged
	defer response.Body.Close()              // Unchanged
	body, _ := ioutil.ReadAll(response.Body) // Unchanged
	// Send body length value via channel.
	channel <- len(body)
}
func TestHandler(c *gin.Context) {
	size := make(chan int)
	resp_godoc := make(chan int)
	resp_exm := make(chan int)
	go responseSize("https://golang.org/", size)
	go responseSize("https://golang.org/doc", resp_godoc)
	go responseSize("https://example.com", resp_exm)
	//fmt.Println(<-size)
	c.String(http.StatusOK, "ResponseTime of https://golang.org/ is %d \n", <-size)
	c.String(http.StatusOK, "ResponseTime of https://example.com is %d \n", <-resp_exm)
	c.String(http.StatusOK, "ResponseTime of https://golang.org/doc is %d ", <-resp_godoc)
}
