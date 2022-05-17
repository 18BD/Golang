package ctrl

import (
	"context"
	"grpcAsan/libraryService/protoFiles"
	"log"
)

func (s *Server) Ping(ctx context.Context, in *libraryService.PingMessage) (*libraryService.PingMessage, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &libraryService.PingMessage{Body: in.Body}, nil
}


