package ctrl

import (
	"context"
	libraryService "grpcAsan/libraryService/protoFiles"
	"grpcAsan/server/models"
)

func (s *Server) AddBookComplaint(ctx context.Context, request *libraryService.AddBookComplaintRequest) (*libraryService.AddBookComplaintResponse, error) {
	bookComp := models.BookComplaint{
		ID:        int(request.BookComplaint.ID),
		BookID:    int(request.BookComplaint.BookID),
		Title:     request.BookComplaint.Title,
		Complaint: request.BookComplaint.Complaint,
	}

	_, err := s.db.Model(&bookComp).Insert()
	if err != nil {
		return nil, err
	}

	return &libraryService.AddBookComplaintResponse{Response: "OK"}, nil
}

func (s *Server) EditBookComplaint(ctx context.Context, request *libraryService.EditBookComplaintRequest) (*libraryService.EditBookComplaintResponse, error) {
	bookComp := models.BookComplaint{
		ID:        int(request.BookComplaint.ID),
		BookID:    int(request.BookComplaint.BookID),
		Title:     request.BookComplaint.Title,
		Complaint: request.BookComplaint.Complaint,
	}

	_, err := s.db.Model(&bookComp).Where("id = ?0", bookComp.ID).Update()
	if err != nil {
		return nil, err
	}

	return &libraryService.EditBookComplaintResponse{Response: "OK"}, nil
}
