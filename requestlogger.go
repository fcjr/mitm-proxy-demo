package proxy

import (
	"fmt"
	"net/http"
)

func NewRequestLogger() *RequestLogger {
	return &RequestLogger{}
}

// RequestLogger is a RequestModifier logs all request url
type RequestLogger struct{}

// RequestLogger is a RequestModifier logs all request url
func (r RequestLogger) ModifyRequest(req *http.Request) error {
	fmt.Println(req.URL.String())
	return nil
}
