package handler

import (
	. "handler/api/main"
	"net/http"
  . "handler/api/main/src/gone/httpne"
)

func OnRequest(responseWriter http.ResponseWriter, request *http.Request) {
  HandleRequest(HttpResponseWriter{Value: &responseWriter} , HttpRequest{Value:request})
}
