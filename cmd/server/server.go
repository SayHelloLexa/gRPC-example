package main

import (
	"log"
	"net"

	"github.com/sayhellolexa/grpc-message-server/pkg/api"
	"github.com/sayhellolexa/grpc-message-server/pkg/msgserver"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &msgserver.GRPCServer{}
	api.RegisterMessageServiceServer(s, srv)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
