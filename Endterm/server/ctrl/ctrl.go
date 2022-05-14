package ctrl

import (
	"github.com/go-pg/pg/v10"
	"google.golang.org/grpc"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
	"net"
)

type Server struct {
	db *pg.DB
	libraryService.UnsafePingServiceServer
	libraryService.UnsafeBookServiceServer
	libraryService.UnsafeUserServiceServer
	libraryService.UnsafeAdminServiceServer
	libraryService.UnsafeBookComplaintServiceServer
	libraryService.UnsafeUserBookServiceServer
}


func New(network, address string, db *pg.DB) (*Server, error) {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listening on address: %s, network: %s\n", address, network)

	s := &Server{
		db: db,
	}

	grpcServer := grpc.NewServer()

	libraryService.RegisterPingServiceServer(grpcServer, s)
	libraryService.RegisterBookServiceServer(grpcServer, s)
	libraryService.RegisterUserServiceServer(grpcServer, s)
	libraryService.RegisterAdminServiceServer(grpcServer, s)
	libraryService.RegisterBookComplaintServiceServer(grpcServer, s)
	libraryService.RegisterUserBookServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	return s, nil
}
