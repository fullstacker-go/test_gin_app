package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type WebStats struct {
	Domain string `json:"domain_name"`
	Size   int    `json:"response_time"`
}

func ResponseSize(url string, channel chan WebStats) {
	fmt.Println("Getting", url) // Unchanged
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Returning...", err)
		return
	}
	defer response.Body.Close()              // Unchanged
	body, _ := ioutil.ReadAll(response.Body) // Unchanged
	// Send body length value via channel.
	channel <- WebStats{url, len(body)}
}
