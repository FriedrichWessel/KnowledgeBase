package main

import "net/http"
import "fmt"
import "flag"

func main(){
  portPtr := flag.String("port", "10800", "port to serve")
  flag.Parse()
  http.HandleFunc("/", serve)
  fmt.Printf("Server is running on port: " + *portPtr +"\n")
  http.ListenAndServe(":"+*portPtr, nil)
}

func serve(response http.ResponseWriter, request *http.Request){
  path := request.URL.Path[1:]
  fmt.Printf("Serving path: " + path + "\n")
  http.ServeFile(response, request, request.URL.Path[1:])
}
