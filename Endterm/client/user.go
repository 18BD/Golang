package main

import (
	"context"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
)

func getUser(client libraryService.UserServiceClient, req *libraryService.GetUserRequest) error {
	response, err := client.GetUser(context.Background(), req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}
