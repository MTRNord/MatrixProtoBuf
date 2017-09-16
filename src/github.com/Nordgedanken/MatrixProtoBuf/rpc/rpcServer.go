package rpc

import (
	pb "github.com/Nordgedanken/MatrixProtoBuf/matrixProtos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.LoginServiceServer.
type loginServ struct{}

// Login implements helloworld.LoginServiceServer
func (s *loginServ) Login(_ context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{"test", "localhost", r.User, "test", "", ""}, nil
}

func StartRPCServer() {
	log.Println("Started RPC Proxy")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoginServiceServer(s, &loginServ{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
