package easy

type HttpError struct {
	statusCode int
	code       int
	message    string
}

func NewHttpError(statusCode int, code int, message string) HttpError {
	return HttpError{
		statusCode: statusCode,
		code:       code,
		message:    message,
	}
}

func (he HttpError) Error() string {
	return he.message
}

func (he HttpError) StatusCode() int {
	return he.statusCode
}

func (he HttpError) Code() int {
	return he.code
}
