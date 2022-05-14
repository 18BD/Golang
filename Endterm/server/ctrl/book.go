package ctrl

import (
	"context"
	libraryService "grpcAsan/libraryService/protoFiles"
	"grpcAsan/server/models"
)

func (s *Server) AddBook(ctx context.Context, request *libraryService.AddBookRequest) (*libraryService.AddBookResponse, error) {
	book := models.Book{
		BookID: int(request.Book.BookID),
		Author: request.Book.Author,
		Title:  request.Book.Title,
		Rating: int(request.Book.Rating),
	}

	_, err := s.db.Model(&book).Insert()
	if err != nil {
		return nil, err
	}

	return &libraryService.AddBookResponse{Response: "OK"}, nil
}

func (s *Server) GetBooks(ctx context.Context, request *libraryService.GetBooksRequest) (*libraryService.GetBooksResponse, error) {
	var books []models.Book

	err := s.db.Model(&books).Select()
	if err != nil {
		return nil, err
	}

	var res []*libraryService.Book
	for _, b := range books {
		res = append(res, &libraryService.Book{
			BookID: int32(b.BookID),
			Author: b.Author,
			Title:  b.Title,
			Rating: int32(b.Rating),
		})
	}

	return &libraryService.GetBooksResponse{Books: res}, nil
}


