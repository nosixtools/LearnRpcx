module LearnRpcx

go 1.13

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	contrib.go.opencensus.io/exporter/zipkin v0.1.1
	github.com/nosixtools/rpcx-plugins v0.0.0-20180814093430-27e5e99f6176
	github.com/opentracing/opentracing-go v1.1.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.1
	github.com/smallnest/rpcx v0.0.0-20200310110228-122cece1047a
	go.opencensus.io v0.22.3

)
