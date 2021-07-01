package actions

import (
	"context"
	. "grpc-envoy/api"
	"log"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "pong"}, nil
}

func (ns *Server) AddUser(ctx context.Context, req *AddUserRequest) (*User, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	user := &User{
		Id:        uuid.New().String(),
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		CreatedAt: time.Now().String(),
	}
	return user, nil
}
