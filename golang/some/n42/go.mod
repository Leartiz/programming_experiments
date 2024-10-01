module n42

go 1.23.0

require go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.55.0

require (
	github.com/google/uuid v1.6.0 // indirect
	go.opentelemetry.io/otel/log v0.6.0
	go.opentelemetry.io/otel/sdk v1.30.0
	go.opentelemetry.io/otel/sdk/log v0.6.0
	go.opentelemetry.io/otel/sdk/metric v1.30.0
	golang.org/x/sys v0.25.0 // indirect
)

require (
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutlog v0.6.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.30.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.30.0
	go.opentelemetry.io/otel/metric v1.30.0
	go.opentelemetry.io/otel/trace v1.30.0 // indirect
)
