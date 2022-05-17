package cli

import (
	"context"
	"google.golang.org/grpc/metadata"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
)

func (cli *Client) GetUserBooks(client libraryService.UserBookServiceClient) error {
	ctx := metadata.NewOutgoingContext(context.Background(), cli.Metadata)

	response, err := client.GetUserBook(ctx, &libraryService.GetUserBooksRequest{})
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}

func (cli *Client) PostUserBook(client libraryService.UserBookServiceClient, req *libraryService.AddUserBookRequest) error {
	ctx := metadata.NewOutgoingContext(context.Background(), cli.Metadata)

	response, err := client.AddUserBook(ctx, req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}

func (cli *Client) PutUserBook(client libraryService.UserBookServiceClient, req *libraryService.EditUserBookRequest) error {
	ctx := metadata.NewOutgoingContext(context.Background(), cli.Metadata)

	response, err := client.EditUserBook(ctx, req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}

func (cli *Client) DeleteUserBook(client libraryService.UserBookServiceClient, req *libraryService.DeleteUserBookRequest) error {
	ctx := metadata.NewOutgoingContext(context.Background(), cli.Metadata)

	response, err := client.DeleteUserBook(ctx, req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}