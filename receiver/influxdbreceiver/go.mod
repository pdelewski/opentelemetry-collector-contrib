module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/influxdbreceiver

go 1.17

require (
	github.com/influxdata/influxdb-observability/common v0.2.35
	github.com/influxdata/influxdb-observability/influx2otel v0.2.35
	github.com/influxdata/line-protocol/v2 v2.2.1
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.49.0
	go.opentelemetry.io/collector v0.49.1-0.20220422001137-87ab5de64ce4
	go.uber.org/zap v1.21.0

)

require (
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/frankban/quicktest v1.14.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.1 // indirect
	github.com/knadh/koanf v1.4.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/rs/cors v1.8.2 // indirect
	go.opentelemetry.io/collector/pdata v1.0.0-rc2 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.31.0 // indirect
	go.opentelemetry.io/otel v1.6.3 // indirect
	go.opentelemetry.io/otel/metric v0.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.6.3 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/grpc v1.51.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../internal/common

replace go.opentelemetry.io/collector/semconv => go.opentelemetry.io/collector/semconv v0.0.0-20220422001137-87ab5de64ce4
