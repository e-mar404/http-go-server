package main

import (
	"log"

  "e-mar404/http-go-server/internal/emar_http"
)

func main() {
  log.SetPrefix("server: ")
    
  server, err := emar_http.CreateServer("tcp", "localhost:8080")

  if err != nil {
    log.Fatalln(err)
  }

  err = server.Start()
  if err != nil {
    log.Fatalln(err)
  }
}
