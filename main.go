package main

import (
	. "fmt"
	. "handler/api/main"
	. "net/http"
  . "github.com/Patrick-ring-motive/httpne"
)




func main() {
  Println("starting")
	HandleFunc("/", OnRequest)
  Println("handling0")
	ListenAndServe(":0", nil)
	Print("http server up!")
}

func OnRequest(responseWriter ResponseWriter, request *Request) {
  Println("onrequest")
  HandleRequest(HttpResponseWriter{Value:&responseWriter} , HttpRequest{Value:request})
}
