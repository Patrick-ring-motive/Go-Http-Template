package indexgo

import (
	"net/http"
)



func OnServerlessRequest(responseWriter http.ResponseWriter, request *http.Request) {
	HandleRequest(&responseWriter, request)
}

func HandleRequest(responseWriter *http.ResponseWriter, request *http.Request) {
	(*responseWriter).WriteHeader(200)
	defer (*responseWriter).Write([]byte("Hello World"))
}



