package main

import (
	"fmt"
	"log"
	"net"

	"chat"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Our first gRPC server")
	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create new grpc server
	grpcServer := grpc.NewServer()

	ch := chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, &ch)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
