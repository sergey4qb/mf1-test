package grpc

import (
	"github.com/sergey4qb/mf1-test/config"
	"github.com/sergey4qb/mf1-test/services"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/sergey4qb/mf1-test/delivery/grpc/user"

	pb "github.com/sergey4qb/mf1-test/proto/pb"
)

type Server struct {
	Server      *grpc.Server
	netListener net.Listener
}

func New(services services.Services) (*Server, error) {
	listener, err := net.Listen(
		config.LoadConfig().GRPCProtocol,
		config.LoadConfig().GRPCAddress+":"+config.LoadConfig().GRPCPort,
	)
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer()

	srv := &Server{
		Server:      grpcServer,
		netListener: listener,
	}

	srv.registerServices(services)

	return srv, nil
}

func (s *Server) registerServices(services services.Services) {
	userServiceServer := user.NewUserServer(services.GetUser())
	pb.RegisterUserServiceServer(s.Server, userServiceServer)
}

func (s *Server) Start() error {
	log.Printf("gRPC started on %s", s.netListener.Addr().String())

	return s.Server.Serve(s.netListener)
}
