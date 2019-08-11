package model

// Response used to api response entity
type Response struct {
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Detail string      `json:"detail"`
	Level  int8        `json:"level"`
}

// ResponseCount used to api ResponseCount entity
type ResponseCount struct {
	Response
	Total int `json:"total"`
}

// NewResponse used to create new response entity
func NewResponse(msg string, data interface{}) *Response {
	return &Response{
		Msg:   msg,
		Data:  data,
		Level: 0,
	}
}
