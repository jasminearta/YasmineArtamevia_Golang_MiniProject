// utils/response.go
package utils

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewErrorResponse(message string) *Response {
	return &Response{Message: message}
}

func NewSuccessResponse(message string, data interface{}) *Response {
	return &Response{
		Message: message,
		Data:    data,
	}
}
