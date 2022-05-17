package ctrl

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	libraryService "grpcAsan/libraryService/protoFiles"
	"log"
	"net"

	"github.com/go-pg/pg/v10"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
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

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_auth.StreamServerInterceptor(authMiddleware),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_auth.UnaryServerInterceptor(authMiddleware),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

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
