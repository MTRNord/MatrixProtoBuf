package protobufS

import (
	"encoding/json"
	"fmt"
	"github.com/Nordgedanken/MatrixProtoBuf/jsonTypes"
	pb "github.com/Nordgedanken/MatrixProtoBuf/matrixProtos"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func StartProtoServer() {
	log.Println("Started ProtoBuf Proxy")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	http.HandleFunc("/versions", func(w http.ResponseWriter, r *http.Request) {
		var versionRequest pb.VersionRequest
		var hsVersions jsonTypes.HSVersions

		req, err := ioutil.ReadAll(r.Body)

		if err := proto.Unmarshal(req, &versionRequest); err != nil {
			fmt.Println(err)
		}

		reqresp, err := http.Get("https://" + versionRequest.Address + "/_matrix/client/versions")
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
