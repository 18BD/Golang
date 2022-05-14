package main

import (
	"context"
	"grpcAsan/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	network = "tcp"
	address = ":9000"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &chat.Message{Body: reverse(in.Body)}, nil
}

func main() {

	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listening on address: %s, network: %s\n", address, network)

	s := &Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func reverse(s string) (res string) {
	for _, v := range s {
		res = string(v) + res
	}
	return
}
