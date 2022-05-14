package main

import (
	"context"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
)

func getAdmin(client libraryService.AdminServiceClient, req *libraryService.GetAdminRequest) error {
	response, err := client.GetAdmin(context.Background(), req)
	if err != nil {
		return err
	}

	log.Printf("Response from server: %s", response)

	return nil
}
