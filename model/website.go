package model

type WebStats struct {
	Domain string `json:"domain_name"`
	Size   int    `json:"response_time"`
}
