package main

import (
	"LearnRpcx/lib"
	"github.com/smallnest/rpcx/server"
)

func main() {
	init := lib.Init("hello")
	defer init.Close()

	s := server.NewServer()
	s.Register(new(lib.HelloService), "")
	s.Serve("tcp", "127.0.0.1:7702")
}
