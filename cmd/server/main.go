package main

import (
	"fmt"
	"log"
	"net"

	"github.com/thanhtoan1742/task-services/api"
	"github.com/thanhtoan1742/task-services/internal/server"
	"google.golang.org/grpc"
)

const (
	port    int32  = 10443
	address string = "localhost"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		log.Fatalln("Failed to start the server.")
	}

	grpcServer := grpc.NewServer()
	api.RegisterTaskServiceServer(grpcServer, &server.Server{})
	log.Printf("Server is running on [%s:%d]\n", address, port)
	grpcServer.Serve(lis)
}
