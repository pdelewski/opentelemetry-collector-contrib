module github.com/open-telemetry/opentelemetry-collector-contrib/examples/demo/server

go 1.17

require (
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.31.0
	go.opentelemetry.io/otel v1.11.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.29.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.29.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.6.3
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.6.1
	go.opentelemetry.io/otel/metric v0.29.0
	go.opentelemetry.io/otel/sdk v1.11.0
	go.opentelemetry.io/otel/sdk/metric v0.29.0
	go.opentelemetry.io/otel/trace v1.11.0
	google.golang.org/grpc v1.45.0
)

require (
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.6.3 // indirect
	go.opentelemetry.io/proto/otlp v0.15.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace go.opentelemetry.io/collector/semconv => go.opentelemetry.io/collector/semconv v0.0.0-20220422001137-87ab5de64ce4
