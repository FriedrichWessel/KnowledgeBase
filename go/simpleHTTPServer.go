package main

import "net/http"
import "fmt"

func main(){
  http.HandleFunc("/", serve)
  http.ListenAndServe(":8081", nil)
}

func serve(response http.ResponseWriter, request *http.Request){
  path := request.URL.Path[1:]
  fmt.Printf("Serving path: " + path)
  http.ServeFile(response, request, request.URL.Path[1:])
}
