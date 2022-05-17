package cli

import (
	"context"
	"google.golang.org/grpc/metadata"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
)

func (cli *Client) GetAdmin(client libraryService.AdminServiceClient, req *libraryService.GetAdminRequest) error {
	ctx := metadata.NewOutgoingContext(context.Background(), cli.Metadata)

	response, err := client.GetAdmin(ctx, req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}
