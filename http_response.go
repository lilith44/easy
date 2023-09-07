package easy

// HttpResponse defines a basic response payload.
type HttpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Succeed sends a successful response.
func Succeed(data any) HttpResponse {
	return HttpResponse{
		Code:    0,
		Message: "succeeded",
		Data:    data,
	}
}

// Fail sends a failed response.
func Fail(code int, data any) HttpResponse {
	return HttpResponse{
		Code:    code,
		Message: "failed",
		Data:    data,
	}
}
