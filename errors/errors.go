package errors

type ServerError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (e *ServerError) Error() string {
	return e.Message
}

func NewServerError(code int, msg string) *ServerError {
	return &ServerError{
		Code:    code,
		Message: msg,
	}
}
