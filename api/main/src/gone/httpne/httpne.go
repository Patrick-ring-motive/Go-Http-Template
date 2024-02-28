package httpne

import (
	"net/http"
)

func ptr[T any](value T) *T {
	return &value
}

type HttpResponseWriter struct {
	Value *http.ResponseWriter
}

type HttpHeader struct {
	Value *http.Header
}

func (responseWriter HttpResponseWriter) Write(bytes []byte) int {
	length, err := (*responseWriter.Value).Write(bytes)
	if err != nil {
		return 0
	}
	return length
}

func (responseWriter HttpResponseWriter) WriteHeader(statusCode int) {
	(*responseWriter.Value).WriteHeader(statusCode)
}

func (responseWriter HttpResponseWriter) Header() HttpHeader {
	return HttpHeader{Value: ptr((*responseWriter.Value).Header())}
}
