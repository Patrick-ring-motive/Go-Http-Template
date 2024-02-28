package main

import (
	. "fmt"
	. "handler/api/main"
	. "net/http"
  . "handler/api/main/src/gone/httpne"
)




func main() {
	HandleFunc("/", OnRequest)
	ListenAndServe(":0", nil)
	Print("http server up!")
}

func OnRequest(responseWriter ResponseWriter, request *Request) {
  HandleRequest(HttpResponseWriter{Value: &responseWriter} , request)
}
