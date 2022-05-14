package main

import (
	"context"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
)

func getUserBooks(client libraryService.UserBookServiceClient) error {
	response, err := client.GetUserBook(context.Background(), &libraryService.GetUserBooksRequest{})
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}

func postUserBook(client libraryService.UserBookServiceClient, req *libraryService.AddUserBookRequest) error {
	response, err := client.AddUserBook(context.Background(), req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}

func putUserBook(client libraryService.UserBookServiceClient, req *libraryService.EditUserBookRequest) error {
	response, err := client.EditUserBook(context.Background(), req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}

func deleteUserBook(client libraryService.UserBookServiceClient, req *libraryService.DeleteUserBookRequest) error {
	response, err := client.DeleteUserBook(context.Background(), req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}