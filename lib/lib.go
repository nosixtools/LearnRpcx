package lib

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/smallnest/rpcx/share"
	"log"
)

func Init(serverName string) reporter.Reporter {
	// set up a span reporter
	reporter := zipkinhttp.NewReporter("http://127.0.0.1:9411/api/v2/spans")
	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(serverName, "myservice.mydomain.com:80")
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
	return reporter
}

// Generate span info from context, generate a new span when context is empty or
// will generate span from parentSpan
func GenerateSpanWithContext(ctx context.Context, operationName string) (opentracing.Span, context.Context, error) {

	md := ctx.Value(share.ReqMetaDataKey)
	var span opentracing.Span
	var parentSpan opentracing.Span

	tracer := opentracing.GlobalTracer()

	if md != nil {
		carrier := opentracing.TextMapCarrier(md.(map[string]string))
		spanContext, err := tracer.Extract(opentracing.TextMap, carrier)
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			log.Printf("metadata error %s\n", err)
		} else {
			parentSpan = tracer.StartSpan(operationName, ext.RPCServerOption(spanContext))
		}
	}

	if parentSpan != nil {
		span = opentracing.GlobalTracer().StartSpan(operationName, opentracing.ChildOf(parentSpan.Context()))
	} else {
		span = opentracing.StartSpan(operationName)
	}

	metadata := opentracing.TextMapCarrier(make(map[string]string))
	err := tracer.Inject(span.Context(), opentracing.TextMap, metadata)
	if err != nil {
		return nil, nil, err
	}
	ctx = context.WithValue(context.Background(), share.ReqMetaDataKey, (map[string]string)(metadata))
	return span, ctx, nil
}
