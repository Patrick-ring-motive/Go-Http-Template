package indexgo

import (
  "fmt"
	"net/http"
)



func OnServerlessRequest(responseWriter http.ResponseWriter, request *http.Request) {
	HandleRequest(&responseWriter, request)
}

func HandleRequest(responseWriter *http.ResponseWriter, request *http.Request) {
	(*responseWriter).WriteHeader(200)
	writeRes,err := (*responseWriter).Write([]byte("Hello World"))
  if(err!=nil){
    fmt.Println(err)
  }
  (*responseWriter).Write([]byte(fmt.Sprint(writeRes)))
}



