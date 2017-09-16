package main

import (
	"log"

	pb "github.com/Nordgedanken/MatrixProtoBuf/matrixProtos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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
	c := pb.NewLoginServiceClient(conn)

	// Contact the server and print out its response.
	r, err := c.Login(context.Background(), &pb.LoginRequest{"m.password", "test", "", "", ""})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.String())
}
