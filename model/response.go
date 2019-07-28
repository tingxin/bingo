package model

// Response used to api response entity
type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// ResponseCount used to api ResponseCount entity
type ResponseCount struct {
	Response
	Total int `json:"total"`
}

// NewResponse used to create new response entity
func NewResponse(status int, msg string, data interface{}) *Response {
	return &Response{
		Status: status,
		Msg:    msg,
		Data:   data,
	}
}
