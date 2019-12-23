package information

// ErrorResponse 错误回应
type ErrorResponse struct {
	Response
	data error
}

// CreateErrorResponse 返回错误消息
func CreateErrorResponse(err error) ErrorResponse {
	return ErrorResponse{}
}
