package handler

import (
  . "handler/api/main"
  "net/http"
  . "github.com/Patrick-ring-motive/httpne"
 )

func OnRequest(responseWriter http.ResponseWriter, request *http.Request) {
  HandleRequest(HttpResponseWriter{Value:&responseWriter} , HttpRequest{Value:request})
}