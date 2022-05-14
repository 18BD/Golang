package ctrl

import (
	"context"
	libraryService "grpcAsan/libraryService/protoFiles"
	"grpcAsan/server/models"
)

func (s *Server) GetUser(ctx context.Context, request *libraryService.GetUserRequest) (*libraryService.GetUserResponse, error) {
	user := models.User{}

	err := s.db.Model(&user).Where("user_id = ?0", request.UserID).Select()
	if err != nil {
		return nil, err
	}

	return &libraryService.GetUserResponse{User: &libraryService.User{
		UserID:   int32(user.UserID),
		Login:    user.Login,
		Password: user.Password,
	}}, nil
}
