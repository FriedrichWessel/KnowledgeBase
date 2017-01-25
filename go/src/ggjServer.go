package main

import "net/http"
import "fmt"
import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{}


func main(){
  http.HandleFunc("/", wsPage)
  fmt.Printf("Server is running\n")
  http.ListenAndServe(":7777", nil)
}

func wsPage(response http.ResponseWriter, request *http.Request){
  conn, err := upgrader.Upgrade(response, request, nil)
  _, msg, err = conn.ReadMessage()
  if err != nil {
    fmt.Printf("Something went wronk! " + err)
  }
  fmt.Printf("Received Msg: " + msg)
}
