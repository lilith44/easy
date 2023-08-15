package easy

// HttpError defines an error with a http status code and a custom error code.
type HttpError struct {
	statusCode int
	code       int
	message    string
}

// NewHttpError creates a new http error.
func NewHttpError(statusCode int, code int, message string) HttpError {
	return HttpError{
		statusCode: statusCode,
		code:       code,
		message:    message,
	}
}

// StatusCode gets the http status code.
func (he HttpError) StatusCode() int {
	return he.statusCode
}

// Code gets the custom error code.
func (he HttpError) Code() int {
	return he.code
}

// Error implements errors.Error.
func (he HttpError) Error() string {
	return he.message
}
