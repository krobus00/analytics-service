package model

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
	Errors     any    `json:"errors,omitempty"`
	StatusCode int    `json:"-"`
	Err        error  `json:"error,omitempty"`
}

func (m *Response) Error() string {
	return m.Message
}

func NewResponse() *Response {
	return new(Response)
}

func NewDefaultResponse() *Response {
	return &Response{
		Message:    "ok",
		StatusCode: http.StatusOK,
	}
}

func WithBadRequestResponse(errors any) *Response {
	return &Response{
		Message:    "bad Request",
		Errors:     errors,
		StatusCode: http.StatusBadRequest,
	}
}

func (m *Response) WithMessage(message string) *Response {
	m.Message = message
	return m
}

func (m *Response) WithStatusCode(code int) *Response {
	m.StatusCode = code
	return m
}

func (m *Response) WithData(data any) *Response {
	m.Data = data
	return m
}

func (m *Response) BuildResponse(c echo.Context) error {
	return c.JSON(m.StatusCode, m)
}
