module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter

go 1.17

require (
	contrib.go.opencensus.io/exporter/stackdriver v0.13.11
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/collector v0.28.0
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.49.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus v0.49.0
	github.com/stretchr/testify v1.8.1
	go.opencensus.io v0.24.0
	go.opentelemetry.io/collector v0.49.1-0.20220422001137-87ab5de64ce4
	go.opentelemetry.io/collector/pdata v0.49.1-0.20220422001137-87ab5de64ce4
	go.opentelemetry.io/collector/semconv v0.0.0-20220422001137-87ab5de64ce4
	google.golang.org/api v0.105.0
	google.golang.org/genproto v0.0.0-20221206210731-b1a01be3a5f6
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

require (
	cloud.google.com/go/compute v1.13.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.2 // indirect
	cloud.google.com/go/monitoring v1.8.0 // indirect
	cloud.google.com/go/trace v1.4.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v1.4.0 // indirect
	github.com/aws/aws-sdk-go v1.43.41 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.0 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/knadh/koanf v1.4.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/collector/model v0.49.0 // indirect
	go.opentelemetry.io/otel v1.6.3 // indirect
	go.opentelemetry.io/otel/metric v0.29.0 // indirect
	go.opentelemetry.io/otel/sdk v1.6.3 // indirect
	go.opentelemetry.io/otel/trace v1.6.3 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/net v0.0.0-20221014081412-f15817d10f9b // indirect
	golang.org/x/oauth2 v0.0.0-20221014153046-6fdb5e3db783 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus => ../../pkg/translator/opencensus

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace go.opentelemetry.io/collector/semconv => go.opentelemetry.io/collector/semconv v0.0.0-20220422001137-87ab5de64ce4
