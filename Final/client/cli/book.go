package cli

import (
	"context"
	"google.golang.org/grpc/metadata"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
)

func (cli *Client) GetBooks(client libraryService.BookServiceClient) error {
	ctx := metadata.NewOutgoingContext(context.Background(), cli.Metadata)

	response, err := client.GetBooks(ctx, &libraryService.GetBooksRequest{})
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}


func (cli *Client) PostBooks(client libraryService.BookServiceClient, req *libraryService.AddBookRequest) error {
	ctx := metadata.NewOutgoingContext(context.Background(), cli.Metadata)

	response, err := client.AddBook(ctx, req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}


