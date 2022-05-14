package main

import (
	"context"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
)

func postBooksCompliant(client libraryService.BookComplaintServiceClient, req *libraryService.AddBookComplaintRequest) error {
	response, err := client.AddBookComplaint(context.Background(), req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}

func putBooksCompliant(client libraryService.BookComplaintServiceClient, req *libraryService.EditBookComplaintRequest) error {
	response, err := client.EditBookComplaint(context.Background(), req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}