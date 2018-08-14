package main


import (
	"github.com/smallnest/rpcx/server"
	"LearnRpcx/service"
	"fmt"
	"os"
	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"


)

func main() {

	//zipkin
	collector, err := zipkin.NewHTTPCollector("http://localhost:9411/api/v1/spans")
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
		os.Exit(-1)
	}
	// Create our recorder.
	recorder := zipkin.NewRecorder(collector, true, "127.0.0.1:7703", "World")

	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
	}

	opentracing.InitGlobalTracer(tracer)

	s := server.NewServer()
	s.Register(new(service.WorldService), "")
	s.Serve("tcp", "127.0.0.1:7703")
}
