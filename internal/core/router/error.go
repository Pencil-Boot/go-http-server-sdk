package router

import httpservercontract "github.com/Pencil-Boot/go-http-server-sdk/pkg/httpserver/contract"

type httpError struct {
	status  int
	code    string
	message string
	causes  []string
}

func NewError(status int, //TODO: use status from core lib
	code string,
	message string,
	causes []string,
) httpservercontract.HTTPError {
	if causes == nil {
		panic("Error code cannot be nil")
	}
	return &httpError{
		status:  status,
		code:    code,
		message: message,
		causes:  causes,
	}
}

func (e *httpError) Error() string {
	return e.message
}

func (e *httpError) Status() int {
	return e.status
}

func (e *httpError) Code() string {
	return e.code
}

func (e *httpError) Message() string {
	return e.message
}

func (e *httpError) Causes() []string {
	return e.causes
}
