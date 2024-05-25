package emar_http

import "testing"

func TestCreateTCPServer(t *testing.T) {
  server := &TCPServer{address: ":8080"}
  _, err := server.TestConnection()
  if err != nil {
    return
  }
}
