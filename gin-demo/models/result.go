package models

type Result struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"请求信息"`
	Data    interface{} `json:"data" `
}
