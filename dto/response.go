package dto

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
