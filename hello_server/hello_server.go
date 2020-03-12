package main

import (
	"LearnRpcx/service"
	"github.com/smallnest/rpcx/server"
	"log"

	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func main() {

	{
		// set up a span reporter
		reporter := zipkinhttp.NewReporter("http://192.168.0.110:9411/api/v2/spans")
		defer reporter.Close()

		// create our local service endpoint
		endpoint, err := zipkin.NewEndpoint("hello", "myservice.mydomain.com:80")
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

	s := server.NewServer()
	s.Register(new(service.HelloService), "")
	s.Serve("tcp", "127.0.0.1:7702")
}
