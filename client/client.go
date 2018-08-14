package main

import (
	"github.com/smallnest/rpcx/client"
	"fmt"
	"context"
	"os"
	"github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"
	rpcxplugins "github.com/nosixtools/rpcx-plugins/opentracing"
	"github.com/smallnest/rpcx/log"
)

func main() {
	d := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:7702", "")
	xclient := client.NewXClient("HelloService", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()


	collector, err := zipkin.NewHTTPCollector("http://localhost:9411/api/v1/spans")
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
		os.Exit(-1)
	}

	// Create our recorder.
	recorder := zipkin.NewRecorder(collector, true, "127.0.0.1:8080", "hello_world")

	tracer, err := zipkin.NewTracer(recorder, zipkin.ClientServerSameSpan(true), zipkin.TraceID128Bit(true))
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
	}

	opentracing.InitGlobalTracer(tracer)

	span, ctx, err := rpcxplugins.GenerateSpanWithContext(context.Background(), "start")
	result := ""
	xclient.Call(ctx, "Hello", "nosix", &result)
	fmt.Println(result)
	log.Info(span)
	span.Finish()
	collector.Close()
}
