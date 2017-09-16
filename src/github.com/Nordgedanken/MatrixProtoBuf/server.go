package main

import (
	"github.com/Nordgedanken/MatrixProtoBuf/protobufS"
	"github.com/Nordgedanken/MatrixProtoBuf/rpc"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go rpc.StartRPCServer()
	wg.Add(1)
	go protobufS.StartProtoServer()
	wg.Wait()
}
