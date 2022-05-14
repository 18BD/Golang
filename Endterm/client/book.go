package main

import (
	"context"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
)

func getBooks(client libraryService.BookServiceClient) error {
	response, err := client.GetBooks(context.Background(), &libraryService.GetBooksRequest{})
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}


func postBooks(client libraryService.BookServiceClient, req *libraryService.AddBookRequest) error {
	response, err := client.AddBook(context.Background(), req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}


