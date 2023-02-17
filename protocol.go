package easy

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Succeed(data any) Response {
	return Response{
		Code:    0,
		Message: "succeeded",
		Data:    data,
	}
}

func Fail(code int, data any) Response {
	return Response{
		Code:    code,
		Message: "failed",
		Data:    data,
	}
}
