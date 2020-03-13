package main

import (
	"context"
	"fmt"
	rpcxplugins "github.com/nosixtools/rpcx-plugins/opentracing"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/smallnest/rpcx/client"
	"log"
)

func main() {
	d := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:7702", "")
	xclient := client.NewXClient("HelloService", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	{
		// set up a span reporter
		reporter := zipkinhttp.NewReporter("http://192.168.0.110:9411/api/v2/spans")
		defer reporter.Close()

		// create our local service endpoint
		endpoint, err := zipkin.NewEndpoint("client", "myservice.mydomain.com:80")
		if err != nil {
			log.Fatalf("unable to create local endpoint: %+v\n", err)
		}

		// initialize our tracer
		nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
		if err != nil {
			log.Fatalf("unable to create tracer: %+v\n", err)
		}

		// use zipkin-go-opentracing to wrap our tracer
		tracer := zipkinot.Wrap(nativeTracer)

		// optionally set as Global OpenTracing tracer instance
		opentracing.SetGlobalTracer(tracer)
	}

	span, ctx, _ := rpcxplugins.GenerateSpanWithContext(context.Background(), "start")
	// 设置标签
	span.SetTag("设置标签", "标签值")
	// 设置注释
	span.LogKV(
		"event", "soft error",
		"type", "cache timeout",
		"waited.millis", 1500)

	result := ""
	xclient.Call(ctx, "Hello", "nosix", &result)
	fmt.Println(result)
	fmt.Printf("%+v\n", span)
	span.Finish()
}
