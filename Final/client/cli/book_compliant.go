package cli

import (
	"context"
	"google.golang.org/grpc/metadata"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
)

func (cli *Client) PostBooksCompliant(client libraryService.BookComplaintServiceClient, req *libraryService.AddBookComplaintRequest) error {
	ctx := metadata.NewOutgoingContext(context.Background(), cli.Metadata)

	response, err := client.AddBookComplaint(ctx, req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}

func (cli *Client) PutBooksCompliant(client libraryService.BookComplaintServiceClient, req *libraryService.EditBookComplaintRequest) error {
	ctx := metadata.NewOutgoingContext(context.Background(), cli.Metadata)

	response, err := client.EditBookComplaint(ctx, req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}