package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcAsan/client/cli"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
)

const (
	login = "asan"
	password = "12345"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	client := cli.New(login, password)

	log.Println(client.Metadata)

	log.Println("client tests started")

	bookService := libraryService.NewBookServiceClient(conn)
	bookComplaintService := libraryService.NewBookComplaintServiceClient(conn)
	adminService := libraryService.NewAdminServiceClient(conn)
	userService := libraryService.NewUserServiceClient(conn)
	userBookService := libraryService.NewUserBookServiceClient(conn)

	err = client.GetBooks(bookService)
	if err != nil {
		log.Println(err)
	}

	err = client.PostBooks(bookService, &libraryService.AddBookRequest{Book: &libraryService.Book{
		BookID: 1,
		Author: "amogus",
		Title:  "sus",
		Rating: 0,
	}})
	if err != nil {
		log.Println(err)
	}

	err = client.PostBooksCompliant(bookComplaintService, &libraryService.AddBookComplaintRequest{BookComplaint: &libraryService.BookComplaint{
		ID:        1,
		BookID:    1,
		Title:     "123",
		Complaint: "1231231231bad",
	}})
	if err != nil {
		log.Println(err)
	}

	err = client.PutBooksCompliant(bookComplaintService, &libraryService.EditBookComplaintRequest{BookComplaint: &libraryService.BookComplaint{
		ID:        1,
		BookID:    1,
		Title:     "123",
		Complaint: "very bad",
	}})
	if err != nil {
		log.Println(err)
	}

	err = client.GetAdmin(adminService, &libraryService.GetAdminRequest{AdminID: 1})
	if err != nil {
		log.Println(err)
	}

	err = client.GetUser(userService, &libraryService.GetUserRequest{UserID: 1})
	if err != nil {
		log.Println(err)
	}

	err = client.GetUserBooks(userBookService)
	if err != nil {
		log.Println(err)
	}

	err = client.PostUserBook(userBookService, &libraryService.AddUserBookRequest{UserBook: &libraryService.UserBook{
		ID:     1,
		UserID: 1,
		BookID: 1,
		Title:  "adad",
		Status: "123",
		Star:   false,
	}})
	if err != nil {
		log.Println(err)
	}

	err = client.PutUserBook(userBookService, &libraryService.EditUserBookRequest{UserBook: &libraryService.UserBook{
		ID:     1,
		UserID: 1,
		BookID: 1,
		Title:  "amogus",
		Status: "123123",
		Star:   false,
	}})
	if err != nil {
		log.Println(err)
	}

	err = client.DeleteUserBook(userBookService, &libraryService.DeleteUserBookRequest{ID: 1})
	if err != nil {
		log.Println(err)
	}
}
