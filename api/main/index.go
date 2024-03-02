package indexgo

import (
	. "handler/api/main/src/gone/httpne"
	"net/http"
)

func OnServerlessRequest(responseWriter http.ResponseWriter, request *http.Request) {
	HandleRequest(HttpResponseWriter{Value: &responseWriter}, HttpRequest{Value:request})
}

func HandleRequest(responseWriter HttpResponseWriter, request HttpRequest) {
	responseWriter.Write([]byte("Hello World"))
}
