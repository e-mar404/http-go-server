package emar_http

import (
	"errors"
)

type Server interface {
  Start() error
  TestConnection()  (string, error)
}

func CreateServer(protocol string, address string) (Server, error) {
  switch protocol {
    case "tcp":
        return &TCPServer{address: address}, nil
  }

  return nil, errors.New("Invalid protocol")
}

