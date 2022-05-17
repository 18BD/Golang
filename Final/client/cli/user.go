package cli

import (
	"context"
	"google.golang.org/grpc/metadata"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
)

func (cli *Client) GetUser(client libraryService.UserServiceClient, req *libraryService.GetUserRequest) error {
	ctx := metadata.NewOutgoingContext(context.Background(), cli.Metadata)

	response, err := client.GetUser(ctx, req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}
