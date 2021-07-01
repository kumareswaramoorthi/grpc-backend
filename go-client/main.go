package main

import (
	"fmt"
	"grpc-envoy/api"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("192.168.1.2:8081", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := api.NewPingClient(conn)
	response, err := c.AddUser(context.Background(), &api.AddUserRequest{
		Email:     "john@example.com",
		FirstName: "John",
		LastName:  "Doe",
	})
	fmt.Println("came")
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response)
}
