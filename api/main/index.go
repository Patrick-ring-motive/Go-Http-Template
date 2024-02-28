package indexgo

import (
  "fmt"
	"net/http"
  . "handler/api/main/src/gone/httpne"
)



func OnServerlessRequest(responseWriter http.ResponseWriter, request *http.Request) {
	HandleRequest(HttpResponseWriter{Value: &responseWriter} , request)
}

func HandleRequest(responseWriter HttpResponseWriter, request *http.Request) {
	(*responseWriter.Value).WriteHeader(200)
	writeRes:=responseWriter.Write([]byte("Hello World"))
  responseWriter.Write([]byte(fmt.Sprint(writeRes)))
}



