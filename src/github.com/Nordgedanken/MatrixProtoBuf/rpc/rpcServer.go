package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/Nordgedanken/MatrixProtoBuf/jsonTypes"
	"github.com/Nordgedanken/MatrixProtoBuf/matrix/noAuth"
	pb "github.com/Nordgedanken/MatrixProtoBuf/matrixProtos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

const (
	port = ":50051"
)

// server is used to implement pb.LoginServiceServer.
type loginServ struct{}

// Login implements pb.VersionsServiceServer
func (s *loginServ) Login(_ context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	var hsLoginResp jsonTypes.RespLogin
	hsLoginReq := jsonTypes.ReqLogin{
		Type:     r.Type,
		User:     r.User,
		Password: r.Password,
	}

	noAuth.Login(r.Address, &hsLoginReq, &hsLoginResp)

	return &pb.LoginResponse{AccessToken: hsLoginResp.AccessToken, HomeServer: hsLoginResp.HomeServer, UserId: hsLoginResp.UserID, RefreshToken: hsLoginResp.RefreshToken}, nil
}

// server is used to implement pb.LoginServiceServer.
type versionsServ struct{}

// Login implements pb.VersionsServiceServer
func (s *versionsServ) Versions(_ context.Context, r *pb.VersionRequest) (*pb.VersionsResponse, error) {
	var hsVersions jsonTypes.HSVersions
	reqresp, err := http.Get("https://" + r.Address + "/_matrix/client/versions")
	if err != nil {
		panic(err)
	}
	defer reqresp.Body.Close()
	body, err := ioutil.ReadAll(reqresp.Body)

	if err = json.Unmarshal(body, &hsVersions); err != nil {
		fmt.Println(err)
	}

	resp := &pb.VersionsResponse{
		Versions: []*pb.Version{},
	}

	for _, k := range hsVersions.Versions {
		resp.Versions = append(resp.Versions, &pb.Version{Version: k})
	}

	return resp, nil
}

func StartRPCServer() {
	log.Println("Started RPC Proxy")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoginServiceServer(s, &loginServ{})
	pb.RegisterVersionsServiceServer(s, &versionsServ{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
