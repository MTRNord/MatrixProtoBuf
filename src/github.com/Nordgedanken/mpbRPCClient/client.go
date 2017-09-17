package main

import (
	pb "github.com/Nordgedanken/MatrixProtoBuf/matrixProtos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	Login(conn)
	GetHSVersions(conn)
}

func Login(conn *grpc.ClientConn) {
	lc := pb.NewLoginServiceClient(conn)

	// Contact the server and print out its response.
	//TODO REMOVE MY PASSWORD!!!!!
	lr, err := lc.Login(context.Background(), &pb.LoginRequest{Type: "m.login.password", User: "@MTRNord:matrix.org", Password: "", Medium: "", Address: "matrix.org"})
	if err != nil {
		log.Fatalf("could not get Login: %v - %v", grpc.Code(err), grpc.ErrorDesc(err))
	}
	log.Printf("Greeting: %s", lr.String())
}

func GetHSVersions(conn *grpc.ClientConn) {
	vc := pb.NewVersionsServiceClient(conn)

	// Contact the server and print out its response.
	vr, err := vc.Versions(context.Background(), &pb.VersionRequest{Address: "matrix.org"})
	if err != nil {
		log.Fatalf("could not get Version: %v - %v", grpc.Code(err), grpc.ErrorDesc(err))
	}

	log.Println("Versions:")
	for _, k := range vr.Versions {
		log.Println(k.Version)
	}
}
