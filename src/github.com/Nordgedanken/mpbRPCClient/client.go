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
	lc := pb.NewLoginServiceClient(conn)

	// Contact the server and print out its response.
	lr, err := lc.Login(context.Background(), &pb.LoginRequest{"m.password", "test", "", "", ""})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", lr.String())

	vc := pb.NewVersionsServiceClient(conn)

	// Contact the server and print out its response.
	vr, err := vc.Versions(context.Background(), &pb.VersionRequest{Address: "matrix.org"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println("Versions:")
	for _, k := range vr.Versions {
		log.Println(k.Version)
	}
}
