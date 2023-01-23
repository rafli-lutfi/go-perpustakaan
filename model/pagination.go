package model

type Pagination struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Result   interface{} `json:"result"`
}
