package ctrl

import (
	"context"
	libraryService "grpcAsan/libraryService/protoFiles"
	"grpcAsan/server/models"
)

func (s *Server) GetAdmin(ctx context.Context, request *libraryService.GetAdminRequest) (*libraryService.GetAdminResponse, error) {
	admin := models.Admin{}

	err := s.db.Model(&admin).Where("admin_id = ?0", request.AdminID).Select()
	if err != nil {
		return nil, err
	}

	return &libraryService.GetAdminResponse{Admin: &libraryService.Admin{
		AdminID:  int32(admin.AdminID),
		Login:    admin.Login,
		Password: admin.Password,
	}}, nil
}