package protobufS

import (
	"fmt"
	pb "github.com/Nordgedanken/MatrixProtoBuf/matrixProtos"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"os"
)

func StartProtoServer() {
	log.Println("Started ProtoBuf Server")
	c := make(chan *pb.Version)
	//Listen to the TCP port
	listener, err := net.Listen("tcp", "127.0.0.1:2110")
	checkError(err)
	for {
		if conn, err := listener.Accept(); err == nil {
			//If err is nil then that means that data is available for us so we take up this data and pass it to a new goroutine
			go handleProtoClient(conn, c)
		} else {
			continue
		}
	}
}

func handleProtoClient(conn net.Conn, c chan *pb.Version) {
	fmt.Println("Connection established")
	//Close the connection when the function exits
	defer conn.Close()
	//Create a data buffer of type byte slice with capacity of 4096
	data := make([]byte, 4096)
	//Read the data waiting on the connection and put it in the data buffer
	n, err := conn.Read(data)
	checkError(err)
	fmt.Println("Decoding Protobuf message")
	//Create an struct pointer of type ProtobufTest.TestMessage struct
	protodata := new(pb.Version)
	//Convert all the data retrieved into the ProtobufTest.TestMessage struct type
	err = proto.Unmarshal(data[0:n], protodata)
	checkError(err)
	//Push the protobuf message into a channel
	c <- protodata
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
