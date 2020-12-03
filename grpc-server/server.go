package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-server/proto"
	"grpc-server/service"
	"log"
	"net"
)

const grpcPort = 6655

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatal(err)
	}

	g := grpc.NewServer()
	proto.RegisterDemoServer(g, &service.Demo{})

	if err = g.Serve(l); err != nil {
		log.Fatal(err)
	}
}

