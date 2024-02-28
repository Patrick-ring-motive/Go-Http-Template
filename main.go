package main

import (
	. "fmt"
	. "handler/api/main"
	. "net/http"
)




func main() {
	HandleFunc("/", OnRequest)
	HandleFunc("/search*", OnRequest)
	ListenAndServe(":0", nil)
	Print("http server up!")
}

func OnRequest(responseWriter ResponseWriter, request *Request) {
	HandleRequest(&responseWriter, request)
}
