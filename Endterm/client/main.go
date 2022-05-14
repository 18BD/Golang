package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	log.Println("client tests started")

	bookService := libraryService.NewBookServiceClient(conn)
	bookComplaintService := libraryService.NewBookComplaintServiceClient(conn)
	adminService := libraryService.NewAdminServiceClient(conn)
	userService := libraryService.NewUserServiceClient(conn)
	userBookService := libraryService.NewUserBookServiceClient(conn)

	err = getBooks(bookService)
	if err != nil {
		log.Println(err)
	}

	err = postBooks(bookService, &libraryService.AddBookRequest{Book: &libraryService.Book{
		BookID: 2,
		Author: "amogus",
		Title:  "sus",
		Rating: 0,
	}})
	if err != nil {
		log.Println(err)
	}

	err = postBooksCompliant(bookComplaintService, &libraryService.AddBookComplaintRequest{BookComplaint: &libraryService.BookComplaint{
		ID:        1,
		BookID:    1,
		Title:     "123",
		Complaint: "1231231231bad",
	}})
	if err != nil {
		log.Println(err)
	}

	err = putBooksCompliant(bookComplaintService, &libraryService.EditBookComplaintRequest{BookComplaint: &libraryService.BookComplaint{
		ID:        1,
		BookID:    1,
		Title:     "123",
		Complaint: "very bad",
	}})
	if err != nil {
		log.Println(err)
	}

	err = getAdmin(adminService, &libraryService.GetAdminRequest{AdminID: 1})
	if err != nil {
		log.Println(err)
	}

	err = getUser(userService, &libraryService.GetUserRequest{UserID: 1})
	if err != nil {
		log.Println(err)
	}

	err = getUserBooks(userBookService)
	if err != nil {
		log.Println(err)
	}

	err = postUserBook(userBookService, &libraryService.AddUserBookRequest{UserBook: &libraryService.UserBook{
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

	err = putUserBook(userBookService, &libraryService.EditUserBookRequest{UserBook: &libraryService.UserBook{
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

	err = deleteUserBook(userBookService, &libraryService.DeleteUserBookRequest{ID: 1})
	if err != nil {
		log.Println(err)
	}
}
