package main

import (
  "net/http"
  "fmt"
  "github.com/hpcloud/tail"
  "golang.org/x/net/websocket"
)

func handlerTail(w http.ResponseWriter, r *http.Request) {
  t, _ := tail.TailFile("/var/log/nginx/access.log", tail.Config{Follow: true,ReOpen:true})
  for line := range t.Lines {
      fmt.Println(line.Text)
  }
}

func handlerFollow(ws *websocket.Conn) {
  t, _ := tail.TailFile("/var/log/nginx/access.log", tail.Config{Follow: true,ReOpen:true})
  for line := range t.Lines {
      fmt.Println(line.Text)
      ws.Write([]byte(line.Text))
  }
}

func main() {
  http.HandleFunc("/tail", handlerTail)
  http.Handle("/follow", websocket.Handler(handlerFollow))
  http.Handle("/", http.FileServer(http.Dir("./static/")))
  http.ListenAndServe(":8080", nil)
}
