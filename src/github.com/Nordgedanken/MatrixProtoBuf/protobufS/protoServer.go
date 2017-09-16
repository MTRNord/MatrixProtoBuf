package protobufS

import (
	"fmt"
	pb "github.com/Nordgedanken/MatrixProtoBuf/matrixProtos"
	"github.com/golang/protobuf/proto"
	"log"
	"net/http"
	"os"
)

func StartProtoServer() {
	log.Println("Started ProtoBuf Server")

	http.HandleFunc("/versions", func(w http.ResponseWriter, r *http.Request) {
		resp := &pb.VersionsResponse{
			Versions: []*pb.Version{
				{"r0.0.1"},
				{"r0.1.0"},
				{"r0.2.0"},
			},
		}
		out, err := proto.Marshal(resp)
		checkError(err)

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(out)
		checkError(err)
	})
	http.ListenAndServe(":3000", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
