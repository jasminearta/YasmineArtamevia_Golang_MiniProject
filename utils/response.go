package utils

type Response struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) *Response {
	return &Response{Message: message}
}

func NewSuccessResponse(message string) *Response {
	return &Response{Message: message}
}
