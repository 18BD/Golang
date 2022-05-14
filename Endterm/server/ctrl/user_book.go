package ctrl

import (
	"context"
	libraryService "grpcAsan/libraryService/protoFiles"
	"grpcAsan/server/models"
)

func (s *Server) GetUserBook(ctx context.Context, request *libraryService.GetUserBooksRequest) (*libraryService.GetUserBooksResponse, error) {
	var userBooks []models.UserBook

	err := s.db.Model(&userBooks).Select()
	if err != nil {
		return nil, err
	}

	var res []*libraryService.UserBook
	for _, b := range userBooks {
		res = append(res, &libraryService.UserBook{
			ID:     int32(b.ID),
			UserID: int32(b.UserID),
			BookID: int32(b.BookID),
			Title:  b.Title,
			Status: b.Status,
			Star:   b.Star,
		})
	}

	return &libraryService.GetUserBooksResponse{UserBooks: res}, nil
}

func (s *Server) AddUserBook(ctx context.Context, request *libraryService.AddUserBookRequest) (*libraryService.AddUserBookResponse, error) {
	userBook := models.UserBook{
		ID:     int(request.UserBook.ID),
		UserID: int(request.UserBook.UserID),
		BookID: int(request.UserBook.BookID),
		Title:  request.UserBook.Title,
		Status: request.UserBook.Status,
		Star:   request.UserBook.Star,
	}

	_, err := s.db.Model(&userBook).Insert()
	if err != nil {
		return nil, err
	}

	return &libraryService.AddUserBookResponse{Response: "OK"}, nil
}

func (s *Server) EditUserBook(ctx context.Context, request *libraryService.EditUserBookRequest) (*libraryService.EditUserBookResponse, error) {
	userBook := models.UserBook{
		ID:     int(request.UserBook.ID),
		UserID: int(request.UserBook.UserID),
		BookID: int(request.UserBook.BookID),
		Title:  request.UserBook.Title,
		Status: request.UserBook.Status,
		Star:   request.UserBook.Star,
	}

	_, err := s.db.Model(&userBook).Where("id = ?0", userBook.ID).Update()
	if err != nil {
		return nil, err
	}

	return &libraryService.EditUserBookResponse{Response: "OK"}, nil
}

func (s *Server) DeleteUserBook(ctx context.Context, request *libraryService.DeleteUserBookRequest) (*libraryService.DeleteUserBookResponse, error) {
	userBook := models.UserBook{ID: int(request.ID)}

	_, err := s.db.Model(&userBook).Where("id = ?0", userBook.ID).Delete()
	if err != nil {
		return nil, err
	}

	return &libraryService.DeleteUserBookResponse{Response: "OK"}, nil
}