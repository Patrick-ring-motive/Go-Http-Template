package indexgo

import (
  . "github.com/Patrick-ring-motive/httpne"
  . "github.com/Patrick-ring-motive/ione"
  . "github.com/Patrick-ring-motive/String"
	"net/http"
  "strings"
  "io"
  . "github.com/Patrick-ring-motive/utils"
  "fmt"
)

var HostTarget = "go.dev"

func OnServerlessRequest(responseWriter HttpResponseWriter, request *http.Request) {
  fmt.Println("handling1")
	HandleRequest(responseWriter, HttpRequest{Value:request})
}

func HandleRequest(responseWriter HttpResponseWriter, request HttpRequest) {
  fmt.Println(request.Value.URL.RequestURI())
  hostTarget := HostTarget
  hostProxy := request.Value.Host
  if request.Value.URL.Query().Has("hostname") {
    hostTarget = request.Value.URL.Query().Get("hostname")
  }
  request.Value.Host = hostTarget
  reqURI := NewStrin(request.Value.URL.RequestURI()).
    ReplaceAll("?hostname="+request.Value.Host,"").
    ReplaceAll("&hostname="+request.Value.Host,"")
  pathname := "https://" + hostTarget + reqURI[0]
  response := ProxyFetch(pathname, request.Value)
  fmt.Println(response.StatusCode)
  ProxyResponseHeaders(responseWriter.Value, response, hostTarget, hostProxy)
  responseWriter.WriteHeader(response.StatusCode)
  body:=IoReadCloser{Value:&response.Body}
  responseWriter.Write(IoReadAll(body))
}

func CreateRequest(method string, url string, body io.Reader) HttpRequest {
  request := HttpNewRequest(method, url, body)
  return request
}

func TransferRequestHeaders(newRequest *http.Request, oldRequest *http.Request) {
  reqHeaders := oldRequest.Header
  hostOld := oldRequest.Host
  hostNew := newRequest.Host
  for key, val := range reqHeaders {
    for i := 0; i < len(val); i++ {
      newRequest.Header.Add(key, strings.Replace(val[i],
        hostOld,
        hostNew,
        -1))
    }
  }
}

func ProxyResponseHeaders(res *http.ResponseWriter, response *http.Response, hostOld string, hostNew string) {
  resHeaders := (*res).Header()
  responseHeaders := response.Header
  for key, val := range responseHeaders {
    for i := 0; i < len(val); i++ {
      resHeaders.Add(key,
        strings.Replace(val[i],
          hostOld,
          hostNew,
          -1))
    }
  }
  resHeaderMap := ForceType((*res).Header(),func(map[string][]string){})
  delete(resHeaderMap,"X-Frame-Options")
  delete(resHeaderMap,"Content-Security-Policy")
  resHeaderMap["Access-Control-Allow-Origin"]=[]string{"*"}
}


func ProxyFetch(url string, req *http.Request) *http.Response {
  request := CreateRequest(req.Method, url, req.Body)
  TransferRequestHeaders(request.Value, req)
  response := DefaultHttpClient.Do(request)
  return response.Value
}


