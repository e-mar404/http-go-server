package emar_http

import (
	"errors"
	"fmt"
	"log"
	"net"
)

type TCPServer struct {
  address     string 
}

func (ts *TCPServer) Start() error {
  log.SetPrefix("TCPServer.Start(): ")

  server, err := net.Listen("tcp", ts.address)
  if err != nil {
    return errors.New(err.Error())
  }

  fmt.Printf("Listening @ %s\n", ts.address)

  defer server.Close()
  for {
    conn, err := server.Accept()
  
    if err != nil {
      log.Fatalln(err)
    }

    go HandleConnection(conn) 
  }
}

func HandleConnection(conn net.Conn) {
  tmp := make([]byte, 1024)
  _, err := conn.Read(tmp)

  if err != nil {
    log.Fatalln(err)
  }

  fmt.Printf("request\n%s", tmp)

  res := []byte("HTTP/1.1 200 OK\r\n\r\n")
  n, err := conn.Write(res)
  if err != nil {
    log.Fatal(err.Error())
  }

  fmt.Printf("bytes writen: %d", n)

  conn.Close()
}
