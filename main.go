package main

import (
	"fmt"
	"grpc-envoy/actions"
	"grpc-envoy/api"
	"log"
	"net"

	"google.golang.org/grpc"
)

// main start a gRPC server and waits for connection
func main() {

	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := actions.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	api.RegisterPingServer(grpcServer, &s)

	// start the server
	fmt.Println("gRPC Server Started ...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
