package indexgo

import (
	. "handler/api/main/src/gone/httpne"
	"net/http"
)

func OnServerlessRequest(responseWriter http.ResponseWriter, request *http.Request) {
	HandleRequest(HttpResponseWriter{Value: &responseWriter}, request)
}

func HandleRequest(responseWriter HttpResponseWriter, request *http.Request) {
	responseWriter.Write([]byte("Hello World"))
}
