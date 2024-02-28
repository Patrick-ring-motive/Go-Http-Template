package httpne

import (
	"net/http"
)

type HttpResponseWriter struct {
	Value *http.ResponseWriter
}

func (responseWriter HttpResponseWriter) Write(bytes []byte) (int) {
	length, err := (*responseWriter.Value).Write(bytes)
  if(err != nil){
    return 0
  }
	return length
}
