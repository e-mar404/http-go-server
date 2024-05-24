package main

import (
	"errors"
	"fmt"
	"log"
	"net"
)

type TCPServer struct {
  port          string 
  network       string
  ln            net.Listener
}

func (ts *TCPServer) Start() (string, error) {
  
  ln, err := net.Listen(ts.network, ts.port)
  ts.ln = ln

  if err != nil {
    return "", errors.New(err.Error())
  }
  
  return "Listening on port: " + ts.port, nil
}

func main() {
  log.SetPrefix("server: ")

  server := TCPServer{port: ":8080", network: "tcp"}
  msg, err := server.Start()

  if err != nil {
    log.Fatal(err)
  }

  fmt.Print(msg)

  for {
    conn, err := server.ln.Accept()

    if err != nil {
      log.Fatal(err.Error())
    }

    go func (conn net.Conn) {
      tmp := make([]byte, 1024)

      n, err := conn.Read(tmp)
      if err != nil {
        log.Fatal(err.Error())
      }

      fmt.Printf("total length: %d, request: %s", n, tmp)

      conn.Close()
    } (conn)
  }
}
