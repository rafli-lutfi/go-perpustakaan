package model

type ErrorResponse struct {
	Status      int    `json:"status"`
	Error       string `json:"message"`
	Description string `json:"description"`
}
