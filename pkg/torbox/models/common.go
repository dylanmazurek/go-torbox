package models

type BaseResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Detail  string `json:"detail"`
	Data    any    `json:"data"`
}
