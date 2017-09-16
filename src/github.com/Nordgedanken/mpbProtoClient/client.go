package main

import (
	"fmt"
	pb "github.com/Nordgedanken/MatrixProtoBuf/matrixProtos"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

func main() {
	getVersions()
}

func getVersions() {
	var versions pb.VersionsResponse

	resp, err := http.Get("http://localhost:3000/versions")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := proto.Unmarshal(body, &versions); err != nil {
		fmt.Println(err)
	}

	for _, v := range versions.Versions {
		fmt.Println(v.Version)
	}
}
